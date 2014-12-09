package qthulhu

import (
	"testing"

	"github.com/hashicorp/raft"
)

func TestPartitionStore(t *testing.T) {

}

func TestPartitionStoreFirstIndex(t *testing.T) {
	s, err := NewPartitionStore(dbPath())
	ok(t, err)
	defer s.Close()
	i, err := s.FirstIndex()
	ok(t, err)
	equals(t, i, uint64(0))
}

func TestPartitionStoreLastIndex(t *testing.T) {
	s, err := NewPartitionStore(dbPath())
	ok(t, err)
	defer s.Close()
	i, err := s.LastIndex()
	ok(t, err)
	equals(t, i, uint64(0))
}

func TestPartitionStoreStoreLog(t *testing.T) {
	s, err := NewPartitionStore(dbPath())
	ok(t, err)
	defer s.Close()
	log := raft.Log{
		Index: 1,
		Term:  1,
		Type:  raft.LogCommand,
		Data:  []byte("first"),
	}

	err = s.StoreLog(&log)
	ok(t, err)

	i, err := s.FirstIndex()
	ok(t, err)
	equals(t, i, uint64(1))

	i, err = s.LastIndex()
	ok(t, err)
	equals(t, i, uint64(1))
}
