package varlog

import (
	"context"
	"strconv"
	"testing"

	"github.com/golang/mock/gomock"
	. "github.com/smartystreets/goconvey/convey"
	"go.uber.org/zap"

	"github.com/kakao/varlog/pkg/logclient"
	"github.com/kakao/varlog/pkg/types"
	"github.com/kakao/varlog/pkg/util/runner"
	"github.com/kakao/varlog/pkg/verrors"
	"github.com/kakao/varlog/proto/varlogpb"
	_ "github.com/kakao/varlog/vtesting"
)

func TestTrim(t *testing.T) {
	Convey("Given varlog client", t, func() {
		const (
			numStorageNodes  = 10
			minStorageNodeID = 0
			topicID          = types.TopicID(1)
		)

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		replicasRetriever := NewMockReplicasRetriever(ctrl)
		replicasMap := make(map[types.LogStreamID][]varlogpb.LogStreamReplica, numStorageNodes)
		for i := minStorageNodeID; i < numStorageNodes; i++ {
			replicasMap[types.LogStreamID(i)] = []varlogpb.LogStreamReplica{
				{
					StorageNode: varlogpb.StorageNode{
						StorageNodeID: types.StorageNodeID(i),
						Address:       "127.0.0.1:" + strconv.Itoa(i),
					},
					TopicLogStream: varlogpb.TopicLogStream{
						LogStreamID: types.LogStreamID(i),
					},
				},
			}
		}
		replicasRetriever.EXPECT().All(topicID).Return(replicasMap).MaxTimes(1)

		createMockLogClientManager := func(expectedTrimResults []error) *logclient.MockLogClientManager {
			logCLManager := logclient.NewMockLogClientManager(ctrl)
			logCLManager.EXPECT().GetOrConnect(gomock.Any(), gomock.Any(), gomock.Any()).DoAndReturn(
				func(_ context.Context, storageNodeID types.StorageNodeID, storagNodeAddr string) (logclient.LogIOClient, error) {
					logCL := logclient.NewMockLogIOClient(ctrl)
					expectedTrimResult := expectedTrimResults[int(storageNodeID)]
					logCL.EXPECT().TrimDeprecated(gomock.Any(), gomock.Any(), gomock.Any()).Return(expectedTrimResult)
					return logCL, nil
				},
			).Times(numStorageNodes)
			return logCLManager
		}

		vlg := &logImpl{}
		vlg.logger = zap.L()
		vlg.runner = runner.New("varlog-test", zap.L())
		vlg.replicasRetriever = replicasRetriever

		Convey("When all storage nodes fail to trim", func() {
			var errs []error
			for i := 0; i < numStorageNodes; i++ {
				errs = append(errs, verrors.ErrInternal)
			}
			vlg.logCLManager = createMockLogClientManager(errs)

			Convey("Then the trim should fail", func() {
				err := vlg.Trim(context.TODO(), topicID, types.GLSN(1), TrimOption{})
				So(err, ShouldNotBeNil)
			})
		})

		Convey("When at least one of all storage node (first) succeeds to trim", func() {
			var errs []error
			errs = append(errs, nil)
			for i := 1; i < numStorageNodes; i++ {
				errs = append(errs, verrors.ErrInternal)
			}
			vlg.logCLManager = createMockLogClientManager(errs)

			Convey("Then the trim should succeed", func() {
				err := vlg.Trim(context.TODO(), topicID, types.GLSN(1), TrimOption{})
				So(err, ShouldBeNil)
			})
		})

		Convey("When at least one of all storage node (last) succeeds to trim", func() {
			var errs []error
			for i := 0; i < numStorageNodes-1; i++ {
				errs = append(errs, verrors.ErrInternal)
			}
			errs = append(errs, nil)
			vlg.logCLManager = createMockLogClientManager(errs)

			Convey("Then the trim should succeed", func() {
				err := vlg.Trim(context.TODO(), topicID, types.GLSN(1), TrimOption{})
				So(err, ShouldBeNil)
			})
		})

		Convey("When at least one of all storage node (middle) succeeds to trim", func() {
			var errs []error
			for i := 0; i < numStorageNodes; i++ {
				errs = append(errs, verrors.ErrInternal)
			}
			errs[numStorageNodes/2] = nil
			vlg.logCLManager = createMockLogClientManager(errs)

			Convey("Then the trim should succeed", func() {
				err := vlg.Trim(context.TODO(), topicID, types.GLSN(1), TrimOption{})
				So(err, ShouldBeNil)
			})
		})
	})
}
