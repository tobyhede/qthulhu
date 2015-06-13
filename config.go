package qthulhu

import (
	"encoding/json"
	"log"
	"os"
	"strings"
)

type Config struct {
	DataDir string
	Logger  *log.Logger
}

func LoadDefaultConfig() *Config {
	conf := &Config{}
	conf.fromJSON("./config.default.json")
	conf.Logger = log.New(os.Stdout, "", log.Ldate|log.Ltime|log.Lshortfile)

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

func (c *Config) pathify(s string) string {
	return strings.Join([]string{c.DataDir, s}, "")
}
