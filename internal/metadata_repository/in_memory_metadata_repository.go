package metadata_repository

import (
	"sync"

	"github.com/kakao/varlog/pkg/varlog"
	varlogpb "github.com/kakao/varlog/proto/varlog"
)

type InMemoryMetadataRepository struct {
	varlogpb.ProjectionDescriptor
	mu sync.RWMutex
}

func NewInMemoryMetadataRepository() *InMemoryMetadataRepository {
	r := &InMemoryMetadataRepository{}
	r.Epoch = 0
	return r
}

func (r *InMemoryMetadataRepository) Propose(epoch uint64, projection *varlogpb.ProjectionDescriptor) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	if epoch < r.Epoch {
		return varlog.ErrSealedEpoch
	}
	r.ProjectionDescriptor = *projection
	return nil
}

func (r *InMemoryMetadataRepository) Get() (*varlogpb.ProjectionDescriptor, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	return &r.ProjectionDescriptor, nil
}
