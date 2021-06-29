package executor

import (
	"context"
	"sync"
	"time"

	"github.com/kakao/varlog/pkg/types"
)

var commitTaskPool = sync.Pool{
	New: func() interface{} {
		return &commitTask{}
	},
}

type commitTask struct {
	highWatermark      types.GLSN
	prevHighWatermark  types.GLSN
	committedGLSNBegin types.GLSN
	committedGLSNEnd   types.GLSN
	committedLLSNBegin types.LLSN

	createdTime    time.Time
	poppedTime     time.Time
	processingTime time.Time
}

func newCommitTask() *commitTask {
	ct := commitTaskPool.Get().(*commitTask)
	ct.createdTime = time.Now()
	return ct
}

func (t *commitTask) release() {
	t.highWatermark = types.InvalidGLSN
	t.prevHighWatermark = types.InvalidGLSN
	t.committedGLSNBegin = types.InvalidGLSN
	t.committedGLSNEnd = types.InvalidGLSN
	t.committedLLSNBegin = types.InvalidLLSN
	t.createdTime = time.Time{}
	t.poppedTime = time.Time{}
	t.processingTime = time.Time{}
	commitTaskPool.Put(t)
}

func (t *commitTask) stale(globalHWM types.GLSN) bool {
	return t.highWatermark <= globalHWM
}

func (t *commitTask) annotate(ctx context.Context, m MeasurableExecutor, discarded bool) {
	if t.createdTime.IsZero() || t.poppedTime.IsZero() || !t.poppedTime.After(t.createdTime) {
		return
	}

	// queue latency
	ms := float64(t.poppedTime.Sub(t.createdTime).Microseconds()) / 1000.0
	m.Stub().Metrics().ExecutorCommitQueueTime.Record(ctx, ms)

	if t.processingTime.IsZero() || !t.processingTime.After(t.poppedTime) {
		return
	}
	// processing time
	ms = float64(t.processingTime.Sub(t.poppedTime).Microseconds()) / 1000.0
	m.Stub().Metrics().ExecutorCommitTime.Record(ctx, ms)
}
