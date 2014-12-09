package qthulhu

import (
	"fmt"
	"math/rand"
	"os"
	"path/filepath"
	"testing"
)

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

	count := 1000

	for i := 0; i < count; i++ {
		v := fmt.Sprintf("%v", i)
		fmt.Println(uint64(i))
		err := s.Put(uint64(i), v)
		ok(t, err)
	}

	k := iToBA(80)
	it := s.Iterator()
	defer it.Close()
	for it.Seek(k); it.Valid(); it.Next() {
		fmt.Printf("%d:%s\n", it.Key(), it.Value())
	}
}

func dbPath() string {
	f := fmt.Sprintf("qthulhu-test-%d", rand.Int())
	return filepath.Join(os.TempDir(), f)
}
