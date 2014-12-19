package qthulhu

import (
	"log"
	"os"
	"testing"
)

type TestStore struct {
	Store
}

func NewTestPartitionStore() *TestStore {
	return &TestStore{}
}

func NewTestLogger() *log.Logger {
	return log.New(os.Stdout, "", 1)
}

func TestFSMSanity(t *testing.T) {
	s := NewTestPartitionStore()
	l := NewTestLogger()
	fsm := NewFSM(s, l)

	assert(t, fsm != nil, "FSM should be created")
}
