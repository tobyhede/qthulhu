package qthulhu

import "testing"

func TestRaft(t *testing.T) {
	conf := LoadDefaultConfig()
	Raft(conf)
}
