package qthulhu

import (
	"fmt"
	"io"
	"log"
	"os"
	"time"

	"github.com/hashicorp/raft"
)

type FSM struct {
	// logOutput io.Writer
	// logger    *log.Logger
	// path      string
	// state     *StateStore
}

func NewFSM() *FSM {
	fsm := &FSM{}
	return fsm
}

func (fsm *FSM) Apply(log *raft.Log) interface{} {
	puts("Apply")
	inspect(log)
	return 0
}

func (fsm *FSM) Restore(io.ReadCloser) error {
	return nil
}
func (fsm *FSM) Snapshot() (raft.FSMSnapshot, error) {
	return nil, nil
}

func Raft(conf *Config) {

	conf.Logger.Print(conf.Address())
	trans, err := raft.NewTCPTransport(conf.Address(), nil, 2, time.Second, nil)
	if err != nil {
		log.Fatal(err)
	}
	logStore, err := NewPartitionStore(conf.LogStorePath())
	stableStore, err := NewPartitionStore(conf.StableStorePath())
	peerStore := raft.NewJSONPeers(conf.PeerStorePath(), trans)

	defer func() {
		logStore.Close()
		stableStore.Close()
	}()

	snapshotStore, err := raft.NewFileSnapshotStore(conf.DataDir, conf.Snapshots, os.Stderr)
	if err != nil {
		// store.Close()
		// return err
		log.Fatal(err)
	}

	// puts(conf.PeerStorePath())
	// NewRaft(conf *Config, fsm FSM, logs LogStore, stable StableStore, snaps SnapshotStore, peerStore PeerStore, trans Transport)

	fsm := NewFSM()
	fmt.Println("Transport: %v", trans)
	fmt.Println("FSM: %v", fsm)
	fmt.Println("LogStore: %v", logStore)
	fmt.Println("StableStore: %v", stableStore)
	fmt.Println("PeerStore: %v", peerStore)
	fmt.Println("SnapshotStore: %v", snapshotStore)

	node, err := raft.NewRaft(conf.Raft, fsm, logStore, stableStore, snapshotStore, peerStore, trans)

	if err != nil {
		log.Fatal(err)
	}
	puts(node)
}
