package qthulhu

import (
	"io"
	"log"

	"github.com/hashicorp/raft"
)

type FSM struct {
	store  *RocksDBStore
	logger *log.Logger
}

type Message struct {
	Offset []byte
	Data   []byte
}

func NewMessage(offset uint64, data []byte) *Message {
	return &Message{uint64ToBytes(offset), data}
}

func NewFSM(s *RocksDBStore, l *log.Logger) *FSM {
	return &FSM{store: s, logger: l}
}

func (f *FSM) Apply(log *raft.Log) interface{} {
	var m Message

	if err := decode(log.Data, &m); err != nil {
		f.logger.Fatalf("Failed to decode data: %v", err)
		return err
	}

	if err := f.store.Put(m.Offset, m.Data); err != nil {
		f.logger.Fatalf("Failed to store data: %v", err)
		return err
	}
	return nil
}

func (fsm *FSM) Restore(io.ReadCloser) error {
	return nil
}

func (fsm *FSM) Snapshot() (raft.FSMSnapshot, error) {
	return nil, nil
}
