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

func Raft(config *Config) {
	conf := raft.DefaultConfig()
	conf.ShutdownOnRemove = false

	trans, err := raft.NewTCPTransport("127.0.0.1:0", nil, 2, time.Second, nil)
	if err != nil {
		log.Fatal(err)
	}

	logStore, err := NewPartitionStore(config.LogStorePath())
	stableStore, err := NewPartitionStore(config.StableStorePath())
	peerStore := raft.NewJSONPeers(config.PeerStorePath(), trans)

	defer func() {
		logStore.Close()
		stableStore.Close()
	}()

	snapshotStore, err := raft.NewFileSnapshotStore(config.DataDir, config.Snapshots, os.Stderr)
	if err != nil {
		// store.Close()
		// return err
		log.Fatal(err)
	}

	// puts(config.PeerStorePath())
	// NewRaft(conf *Config, fsm FSM, logs LogStore, stable StableStore, snaps SnapshotStore, peerStore PeerStore, trans Transport)

	fsm := NewFSM()
	fmt.Println("Transport: %v", trans)
	fmt.Println("FSM: %v", fsm)
	fmt.Println("LogStore: %v", logStore)
	fmt.Println("StableStore: %v", stableStore)
	fmt.Println("PeerStore: %v", peerStore)
	fmt.Println("SnapshotStore: %v", snapshotStore)

	node, err := raft.NewRaft(conf, fsm, logStore, stableStore, snapshotStore, peerStore, trans)

	if err != nil {
		log.Fatal(err)
	}
	puts(node)
}
