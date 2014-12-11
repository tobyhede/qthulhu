package qthulhu

import (
	"fmt"
	"math/rand"

	"testing"
	"time"
)

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

func TestRocksDBStore(t *testing.T) {
	s := NewRocksDBStore(dbPath())
	defer s.Close()

	k := uint64ToBytes(uint64(9223372036854775807))
	err := s.Put(k, []byte("world"))
	ok(t, err)

	v, err := s.Get(k)
	ok(t, err)
	equals(t, string(v), "world")
}

func TestRocksDBIteration(t *testing.T) {

	s := NewRocksDBStore(dbPath())
	defer s.Close()

	generate(t, s, 0, 500)
	generate(t, s, 9999, 500)

	it := s.Iterator()
	defer it.Close()

	i := uint64(99)
	for it.Seek(uint64ToBytes(i)); it.Valid(); it.Next() {
		k := bytesToUint64(it.Key())
		assert(t, i <= k, "I should be less than K")
		// fmt.Printf("%d:%s\n", k, it.Value())
		i++
	}

	k, err := s.LastKey()
	ok(t, err)
	equals(t, int(k), 9999+500-1)
}

func generate(t *testing.T, s *RocksDBStore, start, count int) {
	for i := start; i < (start + count); i++ {
		v := fmt.Sprintf("%v", i)
		k := uint64ToBytes(uint64(i))
		err := s.Put(k, []byte(v))
		ok(t, err)
	}
}
