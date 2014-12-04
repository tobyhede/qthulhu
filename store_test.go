package qthulhu

import (
	"fmt"
	"math/rand"
	"os"
	"path/filepath"
	"testing"
)

func TestStore(t *testing.T) {
	p := tmpPath()
	s := NewStore(p)
	defer s.Close()

	err := s.Put("hello", "world")
	ok(t, err)

	v, err := s.Get("hello")
	ok(t, err)
	equals(t, v, "world")
}

func TestIteration(t *testing.T) {

}

func tmpPath() string {
	f := fmt.Sprintf("qthulhu-test-%d", rand.Int())
	return filepath.Join(os.TempDir(), f)
}
