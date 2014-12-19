package qthulhu

import (
	"io"
	"log"

	"github.com/hashicorp/raft"
)

type FSM struct {
	store  *Store
	logger *log.Logger
}

func NewFSM(s Store, l *log.Logger) *FSM {
	return &FSM{store: &s, logger: l}
}

func (fsm *FSM) Apply(log *raft.Log) interface{} {
	puts("Apply")
	inspect(log)
	return 0
}

func (fsm *FSM) Restore(io.ReadCloser) error {
	return nil
}

func (fsm *FSM) Snapshot() (raft.FSMSnapshot, error) {
	return nil, nil
}
