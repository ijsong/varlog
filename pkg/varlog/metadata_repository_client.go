package varlog

import (
	"context"

	pb "github.com/kakao/varlog/proto/metadata_repository"
	varlogpb "github.com/kakao/varlog/proto/varlog"
)

type MetadataRepositoryClient interface {
	Propose(context.Context, uint64, *varlogpb.ProjectionDescriptor) error
	Get(context.Context, uint64) (*varlogpb.ProjectionDescriptor, error)
	Close() error
}

type metadataRepositoryClient struct {
	rpcConn *rpcConn
	client  pb.MetadataRepositoryServiceClient
}

func NewMetadataRepositoryClient(address string) (MetadataRepositoryClient, error) {
	rpcConn, err := newRpcConn(address)
	if err != nil {
		return nil, err
	}
	return NewMetadataRepositoryClientFromRpcConn(rpcConn)
}

func NewMetadataRepositoryClientFromRpcConn(rpcConn *rpcConn) (MetadataRepositoryClient, error) {
	client := &metadataRepositoryClient{
		rpcConn: rpcConn,
		client:  pb.NewMetadataRepositoryServiceClient(rpcConn.conn),
	}
	return client, nil
}

func (c *metadataRepositoryClient) Close() error {
	return c.rpcConn.close()
}

func (c *metadataRepositoryClient) Propose(ctx context.Context, epoch uint64, projection *varlogpb.ProjectionDescriptor) error {
	_, err := c.client.Propose(ctx, &pb.ProposeRequest{Epoch: epoch, Projection: projection})
	if err != nil {
		return err
	}
	return nil
}

func (c *metadataRepositoryClient) Get(ctx context.Context, epoch uint64) (*varlogpb.ProjectionDescriptor, error) {
	rsp, err := c.client.Get(ctx, &pb.GetRequest{Epoch: epoch})
	if err != nil {
		return nil, err
	}
	//projection := NewProjectionFromProto(rsp.GetProjection())
	return rsp.GetProjection(), nil
}
