package qthulhu

import (
	"fmt"
	"math/rand"
	"os"

	"testing"
	"time"
)

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

func partition() string {
	return fmt.Sprintf("qthulhu-test-%d", rand.Int())
}

func config() *Config {
	return &Config{DataDir: os.TempDir()}
}

func TestRocksDBStore(t *testing.T) {

	s, err := NewRocksDBStore(partition(), config())
	defer s.Close()
	defer s.Destroy()
	ok(t, err)

	k := uint64(9223372036854775807)
	err = s.Set(k, []byte("world"))
	ok(t, err)

	v, err := s.Get(k)
	ok(t, err)
	equals(t, string(v), "world")

	v, err = s.Get(uint64(1))
	equals(t, len(v), 0)
	equals(t, "not found", err.Error())
}

func TestRocksDBNextKey(t *testing.T) {
	s, err := NewRocksDBStore(partition(), config())
	defer s.Close()
	defer s.Destroy()
	ok(t, err)
	generate(t, s, 0, 99)

	k := s.NextKey()
	equals(t, 100, int(k))
}

func TestRocksDBAppend(t *testing.T) {

	s, err := NewRocksDBStore(partition(), config())
	defer s.Close()
	defer s.Destroy()
	ok(t, err)

	err = s.Append([]byte("hello!"))
	ok(t, err)

	err = s.Append([]byte("world!"))
	ok(t, err)

	v, err := s.Get(uint64(2))
	ok(t, err)
	equals(t, string(v), "world!")

	k := s.NextKey()
	equals(t, 3, int(k))
}

// func BenchmarkAppend(b *testing.B) {
// 	s, _ := NewRocksDBStore(partition(), config())
// 	defer s.Close()
// 	defer s.Destroy()
//
// 	b.ResetTimer()
// 	for i := 0; i < b.N; i++ {
// 		v := fmt.Sprintf("%v", i)
// 		s.Append([]byte(v))
// 	}
// }

func TestRocksDBIteration(t *testing.T) {

	s, err := NewRocksDBStore(partition(), config())
	defer s.Close()
	defer s.Destroy()
	ok(t, err)

	k := uint64(1)
	err = s.Set(k, []byte("world"))
	ok(t, err)

	k = uint64(9999)
	err = s.Set(k, []byte("world"))
	ok(t, err)

	k = uint64(999)
	err = s.Set(k, []byte("world"))
	ok(t, err)

	generate(t, s, 0, 99)

	it := s.Iterator()
	defer it.Close()

	i := uint64(0)
	for it.Seek(uint64ToBytes(i)); it.Valid(); it.Next() {
		k := bytesToUint64(it.Key())
		if i > k {
			t.FailNow()
		}
		// assert(t, i <= k, "I should be less than K")
		// fmt.Printf("%d:%s\n", k, it.Value())
		i++
	}

	equals(t, 100, int(s.NextKey()))

	err = s.ResetKey()
	ok(t, err)

	k = s.LastKey()
	equals(t, 9999, int(k))
}

func TestRocksDBFileSize(t *testing.T) {

	s, err := NewRocksDBStore("999999999999", LoadDefaultConfig())
	defer s.Close()
	defer s.Destroy()
	ok(t, err)

	_, err = s.Get(uint64(50396800))
	// ok(t, err)
	// equals(t, err, nil)
	if err == nil {
		t.FailNow()
	}
	// equals(t, string(v), "world")
	// 50,386,801
	// 999,999,999,999
	// generate(t, s, 1, 9, 999, 999)
}

func generate(t *testing.T, s Store, start, count int) {
	for i := start; i < (start + count); i++ {
		v := fmt.Sprintf("{id: %v, data: 'hello world %v' }", i, i)
		err := s.Append([]byte(v))
		ok(t, err)
	}
}
