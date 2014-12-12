package qthulhu

import (
	"fmt"
	"testing"
	"time"
)

func TestRaft(t *testing.T) {

	conf := LoadDefaultConfig()
	for i := 0; i < 1; i++ {

		c := *conf

		if i == 0 {
			c.Raft.EnableSingleNode = true
		}
		c.DataDir = fmt.Sprintf("./data/%v/", i)

		Raft(&c)
	}

	time.Sleep(100 * time.Second)
}
