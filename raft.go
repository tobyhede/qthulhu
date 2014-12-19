package qthulhu

import (
	"log"
	"net"
	"os"
	"time"

	"github.com/hashicorp/raft"
)

type Closer interface {
	Close() (err error)
}

type Raft struct {
	*raft.Raft
	addr    net.Addr
	closers []Closer
}

func (r *Raft) Close() error {
	var err error
	for _, c := range r.closers {
		err = c.Close()
	}
	return err
}

func NewRaft(conf *Config) (*Raft, error) {
	// conf.Logger.Print(conf.Address())

	trans, err := raft.NewTCPTransport(conf.Address(), nil, 2, time.Second, nil)
	if err != nil {
		log.Fatal(err)
	}

	logStore, err := NewPartitionStore(conf.LogStorePath(), conf.Logger)
	stableStore, err := NewPartitionStore(conf.StableStorePath(), conf.Logger)

	// peerStore := conf.PeerStore
	peerStore := raft.NewJSONPeers(conf.PeerStorePath(), trans)

	snapshotStore, err := raft.NewFileSnapshotStore(conf.DataDir, conf.Snapshots, os.Stderr)
	if err != nil {
		// store.Close()
		// return err
		log.Fatal(err)
	}

	// puts(conf.PeerStorePath())
	// NewRaft(conf *Config, fsm FSM, logs LogStore, stable StableStore, snaps SnapshotStore, peerStore PeerStore, trans Transport)

	fsm := NewFSM("./data/fsm", conf.Logger)
	// fmt.Println("Transport: %v", trans)
	// fmt.Println("FSM: %v", fsm)
	// fmt.Println("LogStore: %v", logStore)
	// fmt.Println("StableStore: %v", stableStore)
	// fmt.Println("PeerStore: %v", peerStore)
	// fmt.Println("SnapshotStore: %v", snapshotStore)

	// Ensure local host is always included if we are in bootstrap mode
	// if conf.Bootstrap {
	// 	peers, err := peerStore.Peers()
	// 	if err != nil {
	// 		// store.Close()
	// 		return nil, err
	// 	}
	// 	if !raft.PeerContained(peers, trans.LocalAddr()) {
	// 		peerStore.SetPeers(raft.AddUniquePeer(peers, trans.LocalAddr()))
	// 	}
	// }

	node, err := raft.NewRaft(conf.Raft, fsm, logStore, stableStore, snapshotStore, peerStore, trans)
	if err != nil {
		logStore.Close()
		stableStore.Close()
		trans.Close()
		// log.Fatal(err)
	}
	closers := []Closer{logStore, stableStore, trans}
	return &Raft{node, trans.LocalAddr(), closers}, err
}
