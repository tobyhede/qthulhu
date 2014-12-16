package qthulhu

import (
	"fmt"
	"testing"
	"time"
)

func TestRaft(t *testing.T) {

	conf := LoadDefaultConfig()
	// conf.PeerStore := raft.NewJSONPeers(conf.PeerStorePath(), trans)

	for i := 0; i < 3; i++ {

		c := *conf

		if i == 0 {
			c.Raft.EnableSingleNode = true
		}
		c.DataDir = fmt.Sprintf("./data/%v/", i)

		NewRaft(&c)
	}

	time.Sleep(100 * time.Second)
}
