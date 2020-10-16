package vms

import (
	"context"
	"errors"
	"io"
	"sync"
	"time"

	"github.com/kakao/varlog/pkg/varlog"
	"github.com/kakao/varlog/pkg/varlog/types"
	mrpb "github.com/kakao/varlog/proto/metadata_repository"
	vpb "github.com/kakao/varlog/proto/varlog"
	"go.uber.org/zap"
)

// ClusterMetadataView provides the latest metadata about the cluster.
// TODO: It should have a way to guarantee that ClusterMetadata is the latest.
// TODO: See https://github.com/kakao/varlog/pull/198#discussion_r215542
type ClusterMetadataView interface {
	// ClusterMetadata returns the latest metadata of the cluster.
	ClusterMetadata(ctx context.Context) (*vpb.MetadataDescriptor, error)

	// StorageNode returns the storage node corresponded with the storageNodeID.
	StorageNode(ctx context.Context, storageNodeID types.StorageNodeID) (*vpb.StorageNodeDescriptor, error)

	// LogStreamReplicas returns all of the latest LogStreamReplicaMetas for the given
	// logStreamID. The first element of the returned LogStreamReplicaMeta list is the primary
	// LogStreamReplica.
	// LogStreamReplicas(ctx context.Context, logStreamID types.LogStreamID) ([]*vpb.LogStreamMetadataDescriptor, error)
}

type ClusterMetadataViewGetter interface {
	ClusterMetadataView() ClusterMetadataView
}

var (
	errCMVNoStorageNode = errors.New("cmview: no such storage node")
)

type MetadataRepositoryManager interface {
	ClusterMetadataViewGetter
	//MetadataGetter
	io.Closer

	RegisterStorageNode(ctx context.Context, storageNodeMeta *vpb.StorageNodeDescriptor) error

	UnregisterStorageNode(ctx context.Context, storageNodeID types.StorageNodeID) error

	RegisterLogStream(ctx context.Context, logStreamDesc *vpb.LogStreamDescriptor) error

	UnregisterLogStream(ctx context.Context, logStreamID types.LogStreamID) error

	UpdateLogStream(ctx context.Context, logStreamDesc *vpb.LogStreamDescriptor) error

	// Seal seals logstream corresponded with the logStreamID. It marks the logstream in the
	// cluster metadata stored in MR  as sealed. It returns the last committed GLSN that is
	// confirmed by MR.
	Seal(ctx context.Context, logStreamID types.LogStreamID) (lastCommittedGLSN types.GLSN, err error)

	Unseal(ctx context.Context, logStreamID types.LogStreamID) error

	GetClusterInfo(ctx context.Context) (*mrpb.ClusterInfo, error)

	AddPeer(ctx context.Context, nodeID types.NodeID, peerURL, rpcURL string) error

	RemovePeer(ctx context.Context, nodeID types.NodeID) error
}

var (
	_ MetadataRepositoryManager = (*mrManager)(nil)
	_ ClusterMetadataView       = (*mrManager)(nil)
	_ ClusterMetadataViewGetter = (*mrManager)(nil)
)

type mrManager struct {
	clusterID types.ClusterID

	addrs map[types.NodeID]string
	mu    sync.RWMutex

	connectedNodeID types.NodeID
	cli             varlog.MetadataRepositoryClient
	mcli            varlog.MetadataRepositoryManagementClient

	dirty bool
	meta  *vpb.MetadataDescriptor

	logger *zap.Logger
}

const (
	MRMANAGER_INIT_TIMEOUT time.Duration = 5 * time.Second
)

func NewMRManager(clusterID types.ClusterID, mrAddrs []string, logger *zap.Logger) (MetadataRepositoryManager, error) {
	if logger == nil {
		logger = zap.NewNop()
	}
	logger = logger.Named("mrmanager")

	if len(mrAddrs) == 0 {
		return nil, varlog.ErrInvalid
	}

	mrm := &mrManager{
		clusterID: clusterID,
		dirty:     true,
		logger:    logger,
	}

	ctx, cancel := context.WithTimeout(context.Background(), MRMANAGER_INIT_TIMEOUT)
	defer cancel()

	var err error
Loop:
	for _, addr := range mrAddrs {
		cli, e := varlog.NewMetadataRepositoryManagementClient(addr)
		if e != nil {
			err = e
			continue Loop
		}
		defer cli.Close()

	GET_ADDR:
		for {
			rsp, e := cli.GetClusterInfo(ctx, clusterID)
			if e != nil {
				err = e
				continue Loop
			}

			mrm.updateMemberFromClusterInfo(rsp.GetClusterInfo())
			if len(mrm.addrs) == 0 {
				time.Sleep(100 * time.Millisecond)
				continue GET_ADDR
			}

			break
		}
		return mrm, nil
	}

	return nil, err
}

func (mrm *mrManager) conn() {
	if mrm.cli != nil {
		return
	}

	var rpcConn *varlog.RpcConn
	for nodeID, addr := range mrm.addrs {
		if addr == "" {
			continue
		}

		conn, e := varlog.NewRpcConn(addr)
		if e != nil {
			continue
		}

		mrm.connectedNodeID = nodeID
		rpcConn = conn
		break
	}

	if rpcConn != nil {
		mrm.cli, _ = varlog.NewMetadataRepositoryClientFromRpcConn(rpcConn)
		mrm.mcli, _ = varlog.NewMetadataRepositoryManagementClientFromRpcConn(rpcConn)
	}
}

func (mrm *mrManager) c() varlog.MetadataRepositoryClient {
	if mrm.cli != nil {
		return mrm.cli
	}

	mrm.conn()
	return mrm.cli
}

func (mrm *mrManager) mc() varlog.MetadataRepositoryManagementClient {
	if mrm.cli != nil {
		return mrm.mcli
	}

	mrm.conn()
	return mrm.mcli
}

func (mrm *mrManager) closeClient() error {
	var err error
	if mrm.cli != nil {
		err = mrm.cli.Close()
	}

	mrm.cli = nil
	mrm.mcli = nil
	mrm.connectedNodeID = types.InvalidNodeID

	return err
}

func (mrm *mrManager) Close() error {
	mrm.mu.Lock()
	defer mrm.mu.Unlock()

	return mrm.closeClient()
}

func (mrm *mrManager) ClusterMetadataView() ClusterMetadataView {
	return mrm
}

func (mrm *mrManager) clusterMetadata(ctx context.Context) (*vpb.MetadataDescriptor, error) {
	cli := mrm.c()
	if cli == nil {
		return nil, varlog.ErrNotAccessible
	}

	meta, err := cli.GetMetadata(ctx)
	if err != nil {
		mrm.closeClient()
	}

	return meta, err
}

func (mrm *mrManager) RegisterStorageNode(ctx context.Context, storageNodeMeta *vpb.StorageNodeDescriptor) error {
	mrm.mu.Lock()
	defer mrm.mu.Unlock()

	cli := mrm.c()
	if cli == nil {
		return varlog.ErrNotAccessible
	}

	err := cli.RegisterStorageNode(ctx, storageNodeMeta)
	if err != nil {
		mrm.closeClient()
	}
	mrm.dirty = true
	return err
}

func (mrm *mrManager) UnregisterStorageNode(ctx context.Context, storageNodeID types.StorageNodeID) error {
	mrm.mu.Lock()
	defer mrm.mu.Unlock()

	cli := mrm.c()
	if cli == nil {
		return varlog.ErrNotAccessible
	}

	err := cli.UnregisterStorageNode(ctx, storageNodeID)
	if err != nil {
		mrm.closeClient()
	}
	mrm.dirty = true
	return err
}

func (mrm *mrManager) RegisterLogStream(ctx context.Context, logStreamDesc *vpb.LogStreamDescriptor) error {
	mrm.mu.Lock()
	defer mrm.mu.Unlock()

	cli := mrm.c()
	if cli == nil {
		return varlog.ErrNotAccessible
	}

	err := cli.RegisterLogStream(ctx, logStreamDesc)
	if err != nil {
		mrm.closeClient()
	}
	mrm.dirty = true
	return err
}

func (mrm *mrManager) UnregisterLogStream(ctx context.Context, logStreamID types.LogStreamID) error {
	mrm.mu.Lock()
	defer mrm.mu.Unlock()

	cli := mrm.c()
	if cli == nil {
		return varlog.ErrNotAccessible
	}

	err := cli.UnregisterLogStream(ctx, logStreamID)
	if err != nil {
		mrm.closeClient()
	}
	mrm.dirty = true
	return err
}

func (mrm *mrManager) UpdateLogStream(ctx context.Context, logStreamDesc *vpb.LogStreamDescriptor) error {
	mrm.mu.Lock()
	defer mrm.mu.Unlock()

	cli := mrm.c()
	if cli == nil {
		return varlog.ErrNotAccessible
	}

	err := cli.UpdateLogStream(ctx, logStreamDesc)
	if err != nil {
		mrm.closeClient()
	}
	mrm.dirty = true
	return err
}

// It implements MetadataRepositoryManager.Seal method.
func (mrm *mrManager) Seal(ctx context.Context, logStreamID types.LogStreamID) (lastCommittedGLSN types.GLSN, err error) {
	mrm.mu.Lock()
	defer mrm.mu.Unlock()

	cli := mrm.c()
	if cli == nil {
		return types.InvalidGLSN, varlog.ErrNotAccessible
	}

	glsn, err := cli.Seal(ctx, logStreamID)
	if err != nil {
		mrm.closeClient()
	}
	mrm.dirty = true
	return glsn, err
}

func (mrm *mrManager) Unseal(ctx context.Context, logStreamID types.LogStreamID) error {
	mrm.mu.Lock()
	defer mrm.mu.Unlock()

	cli := mrm.c()
	if cli == nil {
		return varlog.ErrNotAccessible
	}

	err := cli.Unseal(ctx, logStreamID)
	if err != nil {
		mrm.closeClient()
	}
	mrm.dirty = true
	return err
}

func (mrm *mrManager) updateMemberFromClusterInfo(cinfo *mrpb.ClusterInfo) {
	addrs := make(map[types.NodeID]string)
	for nodeID, member := range cinfo.GetMembers() {
		if member.GetEndpoint() != "" {
			addrs[nodeID] = member.GetEndpoint()
		}
	}

	for nodeID := range mrm.addrs {
		if _, ok := addrs[nodeID]; !ok {
			mrm.closeClient()
			break
		}
	}

	mrm.addrs = addrs
}

func (mrm *mrManager) GetClusterInfo(ctx context.Context) (*mrpb.ClusterInfo, error) {
	mrm.mu.Lock()
	defer mrm.mu.Unlock()

	cli := mrm.mc()
	if cli == nil {
		return nil, varlog.ErrNotAccessible
	}

	rsp, err := cli.GetClusterInfo(ctx, mrm.clusterID)
	if err != nil {
		mrm.closeClient()
		return nil, err
	}

	mrm.updateMemberFromClusterInfo(rsp.GetClusterInfo())

	return rsp.GetClusterInfo(), err
}

func (mrm *mrManager) AddPeer(ctx context.Context, nodeID types.NodeID, peerURL, rpcURL string) error {
	mrm.mu.Lock()
	defer mrm.mu.Unlock()

	cli := mrm.mc()
	if cli == nil {
		return varlog.ErrNotAccessible
	}

	err := cli.AddPeer(ctx, mrm.clusterID, nodeID, peerURL)
	if err != nil {
		mrm.closeClient()
		return err
	}

	mrm.addrs[nodeID] = rpcURL

	return nil
}

func (mrm *mrManager) RemovePeer(ctx context.Context, nodeID types.NodeID) error {
	mrm.mu.Lock()
	defer mrm.mu.Unlock()

	cli := mrm.mc()
	if cli == nil {
		return varlog.ErrNotAccessible
	}

	err := cli.RemovePeer(ctx, mrm.clusterID, nodeID)
	if err != nil {
		mrm.closeClient()
		return err
	}

	delete(mrm.addrs, nodeID)
	if mrm.connectedNodeID == nodeID {
		mrm.closeClient()
	}

	return nil
}

func (mrm *mrManager) ClusterMetadata(ctx context.Context) (*vpb.MetadataDescriptor, error) {
	mrm.mu.Lock()
	defer mrm.mu.Unlock()

	if mrm.dirty {
		meta, err := mrm.clusterMetadata(ctx)
		if err != nil {
			return nil, err
		}
		mrm.meta = meta
		mrm.dirty = false
	}
	return mrm.meta, nil
}

func (mrm *mrManager) StorageNode(ctx context.Context, storageNodeID types.StorageNodeID) (*vpb.StorageNodeDescriptor, error) {
	meta, err := mrm.ClusterMetadata(ctx)
	if err != nil {
		return nil, err
	}
	if sndesc := meta.GetStorageNode(storageNodeID); sndesc != nil {
		return sndesc, nil
	}
	return nil, errCMVNoStorageNode
}
