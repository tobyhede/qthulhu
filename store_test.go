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

	fmt.Println(p)

	err := s.Put("hello", "world")
	ok(t, err)

	v := s.Get("hello")
	equals(t, v, "world")
}

func TestIteration(t *testing.T) {

}

func tmpPath() string {
	f := fmt.Sprintf("qthulhu-test-%d", rand.Int())
	return filepath.Join(os.TempDir(), f)
}
