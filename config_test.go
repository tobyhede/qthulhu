package qthulhu

import "testing"

func TestLoadConfig(t *testing.T) {
	path := "./config.json"
	config := LoadConfig(path)

	equals(t, config.DataDir, "./data/")  //default value
}

func TestLoadDefaultConfig(t *testing.T) {
	config := LoadDefaultConfig()
	equals(t, config.DataDir, "./data/")
}
