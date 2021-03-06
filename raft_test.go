package qthulhu

import (
	"fmt"
	"log"
	"testing"
	"time"
)

const (
	applyTimeout = 30 * time.Second
)

func TestRaft(t *testing.T) {
	var nodes []*Raft

	conf := LoadDefaultConfig()
	// conf.PeerStore := raft.NewJSONPeers(conf.PeerStorePath(), trans)
	nodeCount := 3
	for i := 0; i < nodeCount; i++ {

		c := *conf

		if i == 0 {
			c.Bootstrap = true
			c.Raft.EnableSingleNode = true
		}

		inspect(c)
		c.Port = fmt.Sprintf("800%v", i)
		c.DataDir = fmt.Sprintf("./data/%v/", i)

		n, err := NewRaft(&c)
		if err != nil {
			log.Fatal(err)
		}
		nodes = append(nodes, n)
	}

	leader := nodes[0]

	for _, n := range nodes {
		// inspect(n.addr)
		// puts(n.addr)
		leader.AddPeer(n.addr)
	}

	// time.Sleep(5 * time.Second)

	WaitForLeader(leader)

	future := leader.Apply([]byte("test"), applyTimeout)
	if err := future.Error(); err != nil {
		t.Fatalf("err: %v", err)
	}

	time.Sleep(100 * time.Second)
}
