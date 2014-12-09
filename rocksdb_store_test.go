package qthulhu

import (
	"fmt"
	"math/rand"
	"os"
	"path/filepath"
	"testing"
	"time"
)

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

func TestRocksDBStore(t *testing.T) {
	s := NewRocksDBStore(dbPath())
	defer s.Close()

	k := uint64(9223372036854775807)
	err := s.Put(k, "world")
	ok(t, err)

	v, err := s.Get(k)
	ok(t, err)
	equals(t, v, "world")
}

func TestRocksDBIteration(t *testing.T) {

	s := NewRocksDBStore(dbPath())
	defer s.Close()

	generate(t, s, 0, 500)
	generate(t, s, 9999, 500)

	it := s.Iterator()
	defer it.Close()

	i := uint64(99)
	k := iToBA(i)
	for it.Seek(k); it.Valid(); it.Next() {
		k := baToI(it.Key())
		assert(t, i <= k, "I should be less than K")
		i++
		fmt.Printf("%d:%s\n", k, it.Value())
	}
}

func dbPath() string {
	f := fmt.Sprintf("qthulhu-test-%d", rand.Int())
	return filepath.Join(os.TempDir(), f)
}

func generate(t *testing.T, s *RocksDBStore, start, count int) {
	for i := start; i < (start + count); i++ {
		v := fmt.Sprintf("%v", i)
		err := s.Put(uint64(i), v)
		ok(t, err)
	}
}
