package qthulhu

import "testing"

func TestRaft(t *testing.T) {

	conf := LoadDefaultConfig()
	for i := 0; i < 3; i++ {

		c := *conf

		Raft(&c)
	}

}
