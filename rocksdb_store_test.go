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

	err := s.Put("hello", "world")
	ok(t, err)

	v, err := s.Get("hello")
	ok(t, err)
	equals(t, v, "world")
}

func TestRocksDBIteration(t *testing.T) {
	s := NewRocksDBStore(dbPath())
	defer s.Close()

	count := 100

	for i := 0; i < count; i++ {
		k := fmt.Sprintf("%019d", i)
		v := fmt.Sprintf("%v", i)
		err := s.Put(k, v)
		ok(t, err)
	}

	k := []byte("0000000000000000080")
	it := s.Iterator()
	defer it.Close()
	for it.Seek(k); it.Valid(); it.Next() {
		fmt.Printf("%s:%s\n", it.Key(), it.Value())
	}
}

func dbPath() string {
	f := fmt.Sprintf("qthulhu-test-%d", rand.Int())
	return filepath.Join(os.TempDir(), f)
}
