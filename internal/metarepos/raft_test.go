package metarepos

import (
	"context"
	"fmt"
	"net/url"
	"os"
	"sync"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"go.etcd.io/etcd/raft/raftpb"
	"go.uber.org/multierr"
	"go.uber.org/zap"

	"github.com/kakao/varlog/pkg/types"
	"github.com/kakao/varlog/pkg/util/testutil"
	"github.com/kakao/varlog/pkg/util/testutil/ports"
	"github.com/kakao/varlog/vtesting"
)

type stopFunc func(bool)
type leaderFunc func() uint64

type cluster struct {
	peers       []string
	peerToIdx   map[uint64]int
	commitC     []<-chan *raftCommittedEntry
	proposeC    []chan []byte
	confChangeC []chan raftpb.ConfChange
	running     []bool
	stop        []stopFunc
	leader      []leaderFunc
	portLease   *ports.Lease
}

// newCluster creates a cluster of n nodes.
func newCluster(n int) *cluster {
	portLease, err := ports.ReserveWeaklyWithRetry(10000)
	if err != nil {
		panic(err)
	}
	peers := make([]string, n)
	for i := range peers {
		peers[i] = fmt.Sprintf("http://127.0.0.1:%d", portLease.Base()+i)
	}

	clus := &cluster{
		peers:       peers,
		peerToIdx:   make(map[uint64]int),
		commitC:     make([]<-chan *raftCommittedEntry, len(peers)),
		proposeC:    make([]chan []byte, len(peers)),
		confChangeC: make([]chan raftpb.ConfChange, len(peers)),
		running:     make([]bool, len(peers)),
		stop:        make([]stopFunc, len(peers)),
		leader:      make([]leaderFunc, len(peers)),
		portLease:   portLease,
	}

	for i, peer := range clus.peers {
		url, err := url.Parse(peer)
		if err != nil {
			return nil
		}
		nodeID := types.NewNodeID(url.Host)

		os.RemoveAll(fmt.Sprintf("raftdata/wal/%d", nodeID))  //nolint:errcheck,revive // TODO:: Handle an error returned.
		os.RemoveAll(fmt.Sprintf("raftdata/snap/%d", nodeID)) //nolint:errcheck,revive // TODO:: Handle an error returned.
		clus.proposeC[i] = make(chan []byte, 1)
		clus.confChangeC[i] = make(chan raftpb.ConfChange, 1)
		//logger, _ := zap.NewDevelopment()
		logger := zap.NewNop()

		options := raftConfig{
			nodeID:            nodeID,
			join:              false,
			peers:             clus.peers,
			snapCount:         DefaultSnapshotCount,
			snapCatchUpCount:  DefaultSnapshotCatchUpCount,
			maxSnapPurgeCount: 0,
			maxWalPurgeCount:  0,
			raftTick:          vtesting.TestRaftTick(),
			raftDir:           "raftdata",
		}

		rc := newRaftNode(
			options,
			nil,
			clus.proposeC[i],
			clus.confChangeC[i],
			newNopTelmetryStub(),
			logger,
		)

		clus.commitC[i] = rc.commitC
		clus.running[i] = true
		clus.stop[i] = rc.stop
		clus.leader[i] = rc.membership.getLeader
		clus.peerToIdx[uint64(nodeID)] = i

		rc.start()
	}

	return clus
}

func (clus *cluster) close(i int) (err error) {
	if !clus.running[i] {
		return nil
	}

	clus.stop[i](false)
	for range clus.commitC[i] {
		// drain pending commits
	}

	// clean intermediates
	url, _ := url.Parse(clus.peers[i])
	nodeID := types.NewNodeID(url.Host)

	os.RemoveAll(fmt.Sprintf("raftdata/wal/%d", nodeID))  //nolint:errcheck,revive // TODO:: Handle an error returned.
	os.RemoveAll(fmt.Sprintf("raftdata/snap/%d", nodeID)) //nolint:errcheck,revive // TODO:: Handle an error returned.

	clus.running[i] = false

	return
}

// Close closes all cluster nodes and returns an error if any failed.
func (clus *cluster) Close() (err error) {
	for i := range clus.peers {
		if erri := clus.close(i); erri != nil {
			err = multierr.Append(err, erri)
		}
	}

	os.RemoveAll("raftdata") //nolint:errcheck,revive // TODO:: Handle an error returned.

	return multierr.Append(err, clus.portLease.Release())
}

func (clus *cluster) closeNoErrors(t *testing.T) {
	if err := clus.Close(); err != nil {
		t.Fatal(err)
	}
}

func TestProposeOnFollower(t *testing.T) {
	clus := newCluster(3)
	defer clus.closeNoErrors(t)

	donec := make(chan struct{})
	for i := range clus.peers {
		// feedback for "n" committed entries, then update donec
		go func(pC chan<- []byte, cC <-chan *raftCommittedEntry) {
		Loop:
			for range cC {
				break Loop
			}
			donec <- struct{}{}
			for range cC {
				// acknowledge the commits from other nodes so
				// raft continues to make progress
			}
		}(clus.proposeC[i], clus.commitC[i])

		// one message feedback per node
		go func(i int) { clus.proposeC[i] <- []byte(fmt.Sprintf("%d", i)) }(i)
	}

	for range clus.peers {
		<-donec
	}
}

func TestFailoverLeaderElection(t *testing.T) {
	Convey("Given Raft Cluster", t, func(ctx C) {
		clus := newCluster(3)
		defer clus.closeNoErrors(t)

		cancels := make([]context.CancelFunc, 3)

		var wg sync.WaitGroup

		for i := range clus.peers {
			wg.Add(1)
			ctx, cancel := context.WithCancel(context.TODO())
			cancels[i] = cancel
			// feedback for "n" committed entries, then update donec
			go func(ctx context.Context, _ int, _ chan<- []byte, cC <-chan *raftCommittedEntry) {
				defer wg.Done()

			Loop:
				for {
					select {
					case <-ctx.Done():
						break Loop
					case <-cC:
					}
				}
			}(ctx, i, clus.proposeC[i], clus.commitC[i])
		}

		So(testutil.CompareWaitN(10, func() bool {
			return clus.leader[0]() != 0
		}), ShouldBeTrue)

		leader, ok := clus.peerToIdx[clus.leader[0]()]
		So(ok, ShouldBeTrue)

		Convey("When leader crash", func(ctx C) {
			cancels[leader]()
			clus.close(leader) //nolint:errcheck,revive // TODO:: Handle an error returned.

			Convey("Then raft should elect", func(ctx C) {
				So(testutil.CompareWaitN(50, func() bool {
					return clus.leader[(leader+1)%len(clus.peers)]() != 0
				}), ShouldBeTrue)

				for i := range clus.peers {
					cancels[i]()
				}

				wg.Wait()
			})
		})
	})
}
