package mrmanager

//go:generate mockgen -build_flags -mod=vendor -self_package github.com/kakao/varlog/internal/admin/mrmanager -package mrmanager -destination manager_mock.go . ClusterMetadataView,MetadataRepositoryManager

import (
	"context"
	"fmt"
	"io"
	"math"
	"sync"
	"time"

	"github.com/pkg/errors"
	"golang.org/x/sync/singleflight"

	"github.com/kakao/varlog/pkg/mrc"
	"github.com/kakao/varlog/pkg/mrc/mrconnector"
	"github.com/kakao/varlog/pkg/types"
	"github.com/kakao/varlog/pkg/verrors"
	"github.com/kakao/varlog/proto/mrpb"
	"github.com/kakao/varlog/proto/varlogpb"
)

// ClusterMetadataView provides the latest metadata about the cluster.
type ClusterMetadataView interface {
	// ClusterMetadata returns the latest metadata of the cluster.
	ClusterMetadata(ctx context.Context) (*varlogpb.MetadataDescriptor, error)

	// StorageNode returns the storage node corresponded with the storageNodeID.
	StorageNode(ctx context.Context, storageNodeID types.StorageNodeID) (*varlogpb.StorageNodeDescriptor, error)

	// LogStreamReplicas returns all the latest LogStreamReplicaMetas for the given
	// logStreamID. The first element of the returned LogStreamReplicaMeta list is the primary
	// LogStreamReplica.
	// LogStreamReplicas(ctx context.Context, logStreamID types.LogStreamID) ([]*vpb.LogStreamMetadataDescriptor, error)
}

type ClusterMetadataViewGetter interface {
	ClusterMetadataView() ClusterMetadataView
}

const (
	ReloadInterval = time.Second
)

type MetadataRepositoryManager interface {
	ClusterMetadataViewGetter
	io.Closer

	RegisterStorageNode(ctx context.Context, storageNodeMeta *varlogpb.StorageNodeDescriptor) error

	UnregisterStorageNode(ctx context.Context, storageNodeID types.StorageNodeID) error

	RegisterTopic(ctx context.Context, topicID types.TopicID) error

	UnregisterTopic(ctx context.Context, topicID types.TopicID) error

	RegisterLogStream(ctx context.Context, logStreamDesc *varlogpb.LogStreamDescriptor) error

	UnregisterLogStream(ctx context.Context, logStreamID types.LogStreamID) error

	UpdateLogStream(ctx context.Context, logStreamDesc *varlogpb.LogStreamDescriptor) error

	// Seal seals logstream corresponded with the logStreamID. It marks the logstream in the
	// cluster metadata stored in MR  as sealed. It returns the last committed GLSN that is
	// confirmed by MR.
	Seal(ctx context.Context, logStreamID types.LogStreamID) (lastCommittedGLSN types.GLSN, err error)

	Unseal(ctx context.Context, logStreamID types.LogStreamID) error

	GetClusterInfo(ctx context.Context) (*mrpb.ClusterInfo, error)

	AddPeer(ctx context.Context, nodeID types.NodeID, peerURL, rpcURL string) error

	RemovePeer(ctx context.Context, nodeID types.NodeID) error

	NumberOfMR() int
}

var (
	_ MetadataRepositoryManager = (*mrManager)(nil)
	_ ClusterMetadataView       = (*mrManager)(nil)
	_ ClusterMetadataViewGetter = (*mrManager)(nil)
)

type mrManager struct {
	config

	mu        sync.RWMutex
	connector mrconnector.Connector

	sfg     singleflight.Group
	dirty   bool
	updated time.Time
	meta    *varlogpb.MetadataDescriptor
}

const (
	RPCAddrsFetchRetryInterval = 100 * time.Millisecond
)

func New(ctx context.Context, opts ...Option) (MetadataRepositoryManager, error) {
	cfg, err := newConfig(opts)
	if err != nil {
		return nil, err
	}

	mrConnOpts := []mrconnector.Option{
		mrconnector.WithClusterID(cfg.cid),
		mrconnector.WithInitRetryInterval(RPCAddrsFetchRetryInterval),
		mrconnector.WithConnectTimeout(cfg.connTimeout),
		mrconnector.WithRPCTimeout(cfg.callTimeout),
		mrconnector.WithSeed(cfg.metadataRepositoryAddresses),
		mrconnector.WithLogger(cfg.logger),
	}
	tryCnt := cfg.initialMRConnRetryCount + 1
	if tryCnt <= 0 {
		tryCnt = math.MaxInt32
	}

	var connector mrconnector.Connector
	for i := 0; i < tryCnt; i++ {
		connector, err = mrconnector.New(ctx, mrConnOpts...)
		if err != nil {
			time.Sleep(cfg.initialMRConnRetryBackoff)
			continue
		}
		return &mrManager{
			dirty:     true,
			connector: connector,
		}, nil
	}
	err = errors.WithMessagef(err, "mrmanager: tries = %d", tryCnt)
	return nil, err
}

func (mrm *mrManager) c() (mrc.MetadataRepositoryClient, error) {
	// FIXME: use context
	return mrm.connector.Client(context.TODO())
}

func (mrm *mrManager) mc() (mrc.MetadataRepositoryManagementClient, error) {
	// FIXME: use context
	return mrm.connector.ManagementClient(context.TODO())
}

func (mrm *mrManager) Close() error {
	mrm.mu.Lock()
	defer mrm.mu.Unlock()

	return errors.Wrap(mrm.connector.Close(), "mrmanager")
}

func (mrm *mrManager) ClusterMetadataView() ClusterMetadataView {
	return mrm
}

func (mrm *mrManager) clusterMetadata(ctx context.Context) (*varlogpb.MetadataDescriptor, error) {
	cli, err := mrm.c()
	if err != nil {
		return nil, errors.WithMessage(err, "mrmanager: not accessible")
	}

	meta, err := cli.GetMetadata(ctx)
	if err != nil {
		_ = cli.Close()
		return nil, err
	}

	return meta, err
}

func (mrm *mrManager) RegisterStorageNode(ctx context.Context, storageNodeMeta *varlogpb.StorageNodeDescriptor) (err error) {
	mrm.mu.Lock()
	defer func() {
		if err == nil {
			mrm.dirty = true
		}
		mrm.mu.Unlock()
	}()

	cli, err := mrm.c()
	if err != nil {
		return errors.WithMessage(err, "mrmanager: not accessible")
	}

	err = cli.RegisterStorageNode(ctx, storageNodeMeta)
	if err != nil {
		_ = cli.Close()
		return err
	}

	return err
}

func (mrm *mrManager) UnregisterStorageNode(ctx context.Context, storageNodeID types.StorageNodeID) (err error) {
	mrm.mu.Lock()
	defer func() {
		if err == nil {
			mrm.dirty = true
		}
		mrm.mu.Unlock()
	}()

	cli, err := mrm.c()
	if err != nil {
		return errors.WithMessage(err, "mrmanager: not accessible")
	}

	err = cli.UnregisterStorageNode(ctx, storageNodeID)
	if err != nil {
		_ = cli.Close()
		return err
	}

	return err
}

func (mrm *mrManager) RegisterTopic(ctx context.Context, topicID types.TopicID) (err error) {
	mrm.mu.Lock()
	defer func() {
		if err == nil {
			mrm.dirty = true
		}
		mrm.mu.Unlock()
	}()

	cli, err := mrm.c()
	if err != nil {
		return errors.WithMessage(err, "mrmanager: not accessible")
	}

	err = cli.RegisterTopic(ctx, topicID)
	if err != nil {
		_ = cli.Close()
		return err
	}

	return err
}

func (mrm *mrManager) UnregisterTopic(ctx context.Context, topicID types.TopicID) (err error) {
	mrm.mu.Lock()
	defer func() {
		if err == nil {
			mrm.dirty = true
		}
		mrm.mu.Unlock()
	}()

	cli, err := mrm.c()
	if err != nil {
		return errors.WithMessage(err, "mrmanager: not accessible")
	}

	err = cli.UnregisterTopic(ctx, topicID)
	if err != nil {
		_ = cli.Close()
		return err
	}

	return err
}

func (mrm *mrManager) RegisterLogStream(ctx context.Context, logStreamDesc *varlogpb.LogStreamDescriptor) (err error) {
	mrm.mu.Lock()
	defer func() {
		if err == nil {
			mrm.dirty = true
		}
		mrm.mu.Unlock()
	}()

	cli, err := mrm.c()
	if err != nil {
		return errors.WithMessage(err, "mrmanager: not accessible")
	}

	err = cli.RegisterLogStream(ctx, logStreamDesc)
	if err != nil {
		_ = cli.Close()
		return err
	}
	return nil
}

func (mrm *mrManager) UnregisterLogStream(ctx context.Context, logStreamID types.LogStreamID) (err error) {
	mrm.mu.Lock()
	defer func() {
		if err == nil {
			mrm.dirty = true
		}
		mrm.mu.Unlock()
	}()

	cli, err := mrm.c()
	if err != nil {
		return errors.WithMessage(err, "mrmanager: not accessible")
	}

	err = cli.UnregisterLogStream(ctx, logStreamID)
	if err != nil {
		_ = cli.Close()
		return err
	}
	return err
}

func (mrm *mrManager) UpdateLogStream(ctx context.Context, logStreamDesc *varlogpb.LogStreamDescriptor) (err error) {
	mrm.mu.Lock()
	defer func() {
		if err == nil {
			mrm.dirty = true
		}
		mrm.mu.Unlock()
	}()

	cli, err := mrm.c()
	if err != nil {
		return errors.WithMessage(err, "mrmanager: not accessible")
	}

	err = cli.UpdateLogStream(ctx, logStreamDesc)
	if err != nil {
		_ = cli.Close()
		return err
	}
	return err
}

// It implements MetadataRepositoryManager.Seal method.
func (mrm *mrManager) Seal(ctx context.Context, logStreamID types.LogStreamID) (lastCommittedGLSN types.GLSN, err error) {
	mrm.mu.Lock()
	defer func() {
		if err == nil {
			mrm.dirty = true
		}
		mrm.mu.Unlock()
	}()

	cli, err := mrm.c()
	if err != nil {
		return types.InvalidGLSN, errors.WithMessage(err, "mrmanager: not accessible")
	}

	if lastCommittedGLSN, err = cli.Seal(ctx, logStreamID); err != nil {
		_ = cli.Close()
		return types.InvalidGLSN, err
	}
	return lastCommittedGLSN, err
}

func (mrm *mrManager) Unseal(ctx context.Context, logStreamID types.LogStreamID) (err error) {
	mrm.mu.Lock()
	defer func() {
		if err == nil {
			mrm.dirty = true
		}
		mrm.mu.Unlock()
	}()

	cli, err := mrm.c()
	if err != nil {
		return errors.WithMessage(err, "mrmanager: not accessible")
	}

	err = cli.Unseal(ctx, logStreamID)
	if err != nil {
		_ = cli.Close()
		return err
	}
	return err
}

func (mrm *mrManager) GetClusterInfo(ctx context.Context) (*mrpb.ClusterInfo, error) {
	mrm.mu.RLock()
	defer mrm.mu.RUnlock()

	cli, err := mrm.mc()
	if err != nil {
		return nil, errors.WithMessage(err, "mrmanager: not accessible")
	}

	rsp, err := cli.GetClusterInfo(ctx, mrm.cid)
	if err != nil {
		_ = cli.Close()
		return nil, err
	}
	return rsp.GetClusterInfo(), err
}

func (mrm *mrManager) AddPeer(ctx context.Context, nodeID types.NodeID, peerURL, rpcURL string) error {
	mrm.mu.Lock()
	defer mrm.mu.Unlock()

	cli, err := mrm.mc()
	if err != nil {
		return errors.WithMessage(err, "mrmanager: not accessible")
	}

	if err := cli.AddPeer(ctx, mrm.cid, nodeID, peerURL); err != nil {
		if !errors.Is(err, verrors.ErrAlreadyExists) {
			_ = cli.Close()
			return err
		}
		return err
	}

	mrm.connector.AddRPCAddr(nodeID, rpcURL)
	return nil
}

func (mrm *mrManager) RemovePeer(ctx context.Context, nodeID types.NodeID) error {
	mrm.mu.Lock()
	defer mrm.mu.Unlock()

	cli, err := mrm.mc()
	if err != nil {
		return errors.WithMessage(err, "mrmanager: not accessible")
	}

	if err := cli.RemovePeer(ctx, mrm.cid, nodeID); err != nil {
		_ = cli.Close()
		return err
	}
	mrm.connector.DelRPCAddr(nodeID)

	return nil
}

func (mrm *mrManager) ClusterMetadata(ctx context.Context) (*varlogpb.MetadataDescriptor, error) {
	// fail-fast
	if err := ctx.Err(); err != nil {
		return nil, ctx.Err()
	}

	// fast path
	mrm.mu.RLock()
	if !mrm.dirty && time.Since(mrm.updated) <= ReloadInterval {
		meta := mrm.meta
		mrm.mu.RUnlock()
		return meta, nil
	}
	mrm.mu.RUnlock()

	// slow path
	md, err, _ := mrm.sfg.Do("cluster_metadata", func() (interface{}, error) {
		meta, err := mrm.clusterMetadata(ctx)
		if err != nil {
			return nil, fmt.Errorf("cluster metadata: %w", err)
		}
		mrm.mu.Lock()
		mrm.meta = meta
		mrm.dirty = false
		mrm.updated = time.Now()
		mrm.mu.Unlock()
		return meta, nil
	})
	if err != nil {
		return nil, err
	}
	return md.(*varlogpb.MetadataDescriptor), nil
}

func (mrm *mrManager) StorageNode(ctx context.Context, storageNodeID types.StorageNodeID) (*varlogpb.StorageNodeDescriptor, error) {
	meta, err := mrm.ClusterMetadata(ctx)
	if err != nil {
		return nil, err
	}
	return meta.MustHaveStorageNode(storageNodeID)
}

func (mrm *mrManager) NumberOfMR() int {
	return mrm.connector.NumberOfMR()
}
