package qthulhu

import (
	"testing"

	"github.com/hashicorp/raft"
)

func TestEncodeDecode(t *testing.T) {
	log := raft.Log{
		Index: 9,
		Term:  99,
		Type:  raft.LogCommand,
		Data:  []byte("first"),
	}

	b, err := encode(&log)
	ok(t, err)

	var decoded raft.Log
	err = decode(b, &decoded)
	ok(t, err)
	equals(t, decoded.Index, uint64(9))
	equals(t, decoded.Type, raft.LogCommand)
}
