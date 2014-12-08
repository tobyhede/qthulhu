package qthulhu

import "testing"

func TestLoadConfig(t *testing.T) {
	path := "./config.json"
	config := LoadConfig(path)

	equals(t, config.DataDir, "./data/")
}

func TestDefaultConfig(t *testing.T) {
	config := DefaultConfig()

	equals(t, config.DataDir, "./data/")
}
