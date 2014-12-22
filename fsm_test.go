package qthulhu

import (
	"log"
	"os"
	"testing"

	"github.com/hashicorp/raft"
)

type TestStore struct {
	RaftStore
}

func NewTestRaftStore() *TestStore {
	return &TestStore{}
}

func NewTestLogger() *log.Logger {
	return log.New(os.Stdout, "", 1)
}

func NewRaftLog(data []byte) *raft.Log {
	return &raft.Log{
		Index: 1,
		Term:  1,
		Type:  raft.LogCommand,
		Data:  data,
	}
}

// func newTestFSM() *FSM {
// 	s := NewTestRaftStore()
// 	l := NewTestLogger()
// 	return NewFSM(s, l)
// }

// func TestFSMSanity(t *testing.T) {
// 	fsm := newTestFSM()
// 	assert(t, fsm != nil, "FSM should be created")
// }

func TestFSMApply(t *testing.T) {
	l := NewTestLogger()
	s := NewRocksDBStore(dbPath())

	offset := uint64(1)
	fsm := NewFSM(s, l)

	m := NewMessage(offset, []byte("blah"))

	b, err := encode(m)
	ok(t, err)

	log := NewRaftLog(b)
	fsm.Apply(log)

	v, err := s.Get(uint64ToBytes(offset))
	ok(t, err)
	inspect(v)
	inspect(string(v))
}
