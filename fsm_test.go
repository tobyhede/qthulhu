package qthulhu

import "testing"

type TestStore struct {
	RaftStore
}

func NewTestRaftStore() *TestStore {
	return &TestStore{}
}

func newTestFSM() *FSM {
	l := NewTestLogger()
	s := NewRocksDBStore(dbPath())
	return NewFSM(s, l)
}

func TestFSMSanity(t *testing.T) {
	fsm := newTestFSM()
	assert(t, fsm != nil, "FSM should be created")
}

func TestFSMApply(t *testing.T) {

	fsm := newTestFSM()

	offset := uint64(0)
	data := []byte("blah")

	// for i := 1; i <= 5; i++ {
	// 	atomic.AddUint64(&offset, 1)
	// }

	// puts(offset)
	m := NewMessage(offset, data)

	b, err := encode(m)
	ok(t, err)

	log := NewTestRaftLog(b)
	res := fsm.Apply(log)
	equals(t, res, nil)

	v, err := fsm.store.Get(uint64ToBytes(offset))
	ok(t, err)
	equals(t, v, data)
}
