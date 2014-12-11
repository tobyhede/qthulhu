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

	v, err := s.Get([]byte("vtha"))
	equals(t, len(v), 0)
	equals(t, "Key Not Found", err.Error())
}

func TestSetGetUint64(t *testing.T) {
	s, err := NewPartitionStore(dbPath())
	ok(t, err)

	v := uint64(9237409173409)
	err = s.SetUint64(k, v)
	ok(t, err)

	d, err := s.GetUint64(k)
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

func TestPartitionStoreStoreGetLog(t *testing.T) {
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

	var getted raft.Log
	err = s.GetLog(uint64(1), &getted)
	ok(t, err)

	equals(t, getted.Index, uint64(1))
}

func TestPartitionStoreDeleteRange(t *testing.T) {
	s, err := NewPartitionStore(dbPath())
	ok(t, err)
	defer s.Close()

	min := uint64(0)
	max := uint64(5)

	for i := min; i <= max; i++ {
		k = uint64ToBytes(uint64(i))
		err = s.Set(k, v)
		ok(t, err)
	}
	last, err := s.LastIndex()
	equals(t, last, max)
	s.DeleteRange(min, max)

	last, err = s.LastIndex()
	equals(t, last, min)

	d, err := s.Get(uint64ToBytes(max))
	ok(t, err)
	equals(t, d, 1)

}
