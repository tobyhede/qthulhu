package qthulhu

import (
	"testing"

	"github.com/hashicorp/raft"
)

var k = uint64ToBytes(uint64(0))
var v = []byte("helloworld")

func TestPartitionStore(t *testing.T) {

}

func TestSetGet(t *testing.T) {
	s, err := NewPartitionStore(dbPath())
	ok(t, err)

	err = s.Set(k, v)
	ok(t, err)

	d, err := s.Get(k)
	ok(t, err)
	equals(t, d, v)
}

func TestPartitionStoreFirstIndex(t *testing.T) {
	s, err := NewPartitionStore(dbPath())
	ok(t, err)
	defer s.Close()

	ok(t, err)

	for i := 0; i < 5; i++ {
		k = uint64ToBytes(uint64(i))
		err = s.Set(k, v)
		ok(t, err)

		i, err := s.FirstIndex()
		ok(t, err)
		equals(t, i, uint64(0))
	}
}

func TestPartitionStoreLastIndex(t *testing.T) {
	s, err := NewPartitionStore(dbPath())
	ok(t, err)
	defer s.Close()

	for i := 0; i < 5; i++ {
		k = uint64ToBytes(uint64(i))
		err = s.Set(k, v)
		ok(t, err)

		idx, err := s.LastIndex()
		ok(t, err)
		equals(t, uint64(i), idx)
	}

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
