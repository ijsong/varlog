package metadata_repository

import (
	varlogpb "github.com/kakao/varlog/proto/varlog"
)

type MetadataRepository interface {
	Propose(epoch uint64, projection *varlogpb.ProjectionDescriptor) error
	Get() (*varlogpb.ProjectionDescriptor, error)
}
