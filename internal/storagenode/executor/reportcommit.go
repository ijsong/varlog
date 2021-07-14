package executor

import (
	"context"

	"github.com/pkg/errors"

	"github.com/kakao/varlog/pkg/types"
	"github.com/kakao/varlog/proto/snpb"
)

func (e *executor) GetReport(_ context.Context) (snpb.LogStreamUncommitReport, error) {
	if err := e.guard(); err != nil {
		return snpb.InvalidLogStreamUncommitReport, err
	}
	defer e.unguard()

	globalHighWatermark, uncommittedLLSNBegin := e.lsc.reportCommitBase()
	uncommittedLLSNEnd := e.lsc.uncommittedLLSNEnd.Load()
	return snpb.LogStreamUncommitReport{
		LogStreamID:           e.logStreamID,
		HighWatermark:         globalHighWatermark,
		UncommittedLLSNOffset: uncommittedLLSNBegin,
		UncommittedLLSNLength: uint64(uncommittedLLSNEnd - uncommittedLLSNBegin),
	}, nil
}

func (e *executor) Commit(ctx context.Context, commitResult *snpb.LogStreamCommitResult) error {
	if err := e.guard(); err != nil {
		return err
	}
	defer e.unguard()

	// TODO: check validate logic again
	globalHighWatermark, _ := e.lsc.reportCommitBase()
	if commitResult.HighWatermark <= globalHighWatermark {
		// too old
		return errors.New("too old commit result")
	}

	ct := newCommitTask()
	ct.highWatermark = commitResult.HighWatermark
	ct.prevHighWatermark = commitResult.PrevHighWatermark
	ct.committedGLSNBegin = commitResult.CommittedGLSNOffset
	ct.committedGLSNEnd = commitResult.CommittedGLSNOffset + types.GLSN(commitResult.CommittedGLSNLength)
	ct.committedLLSNBegin = commitResult.CommittedLLSNOffset
	if err := e.committer.sendCommitTask(ctx, ct); err != nil {
		ct.release()
		return err
	}
	return nil
}
