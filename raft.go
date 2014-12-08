package qthulhu

import (
	"fmt"
	"log"
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
	fmt.Println("Apply")
	inspect(log)
	// m.logs = append(m.logs, log.Data)
	// return len(m.logs)
	return 0
}

func Raft() {
	conf := raft.DefaultConfig()
	conf.ShutdownOnRemove = false

	// NewRaft(conf *Config, fsm FSM, logs LogStore, stable StableStore, snaps SnapshotStore, peerStore PeerStore, trans Transport)
	trans, err := raft.NewTCPTransport("127.0.0.1:0", nil, 2, time.Second, nil)
	if err != nil {
		// t.Fatalf("err: %v", err)
		log.Fatal(err)
	}

	fsm := NewFSM()
	fmt.Println("Transport: %v", trans)
	fmt.Println("FSM: %v", fsm)
	// node := NewRaft(conf, fsm, logs LogStore, stable StableStore, snaps SnapshotStore, peerStore PeerStore, trans)
}
