package qthulhu

import "testing"

func TestNewStore(t *testing.T) {
	_, err := NewStore("test", config())
	ok(t, err)
}
