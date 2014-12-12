package qthulhu

import (
	"encoding/json"
	"log"
	"os"
	"strings"
)

type Config struct {
	DataDir     string
	PeerStore   string
	LogStore    string
	StableStore string
	Snapshots   int
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

func (c *Config) LogStorePath() string {
	return c.pathify(c.LogStore)
}

func (c *Config) StableStorePath() string {
	return c.pathify(c.StableStore)
}

func (c *Config) PeerStorePath() string {
	return c.DataDir
}

func (c *Config) SnapshotDir() string {
	return c.DataDir
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

func (c *Config) pathify(s string) string {
	return strings.Join([]string{c.DataDir, s}, "")
}
