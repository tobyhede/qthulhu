package qthulhu

import (
	"encoding/json"
	"log"
	"os"
)

type Config struct {
	DataDir     string
	PeerStore   string
	LogStore    string
	StableStore string
}

func LoadDefaultConfig() *Config {
	conf := &Config{}
	conf.fromJSON("./config.default.json")
	return conf
}

func LoadConfig(path string) *Config {
	conf := LoadDefaultConfig()
	conf.fromJSON(path)
	return conf
}

func (c *Config) fromJSON(path string) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatalf("Error loading config in %v\n%v", path, err)
	}

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&c)
	if err != nil {
		log.Fatalf("Error loading config in %v\n%v", path, err)
	}
}
