package qthulhu

import "testing"

func TestLoadConfig(t *testing.T) {
	path := "./config.json"
	config := LoadConfig(path)

	equals(t, config.DataDir, "./data/")  //default value
	equals(t, config.LogStore, "log.log") //overide value

}

func TestLoadDefaultConfig(t *testing.T) {
	config := LoadDefaultConfig()
	equals(t, config.DataDir, "./data/")
}
