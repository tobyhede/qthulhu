package qthulhu

import (
	"encoding/json"
	"fmt"
	"os"
)

type Config struct {
	DataDir string
}

func DefaultConfig() *Config {
	return configFromJSON("./config.default.json")
}

func LoadConfig(path string) *Config {
	conf := configFromJSON(path)
	return setDefaultConfig(conf)
}

func setDefaultConfig(conf *Config) *Config {
	def := DefaultConfig()
	conf.DataDir = def.DataDir
	return conf
}

func configFromJSON(path string) *Config {
	file, err := os.Open(path)
	if err != nil {
		fmt.Println("error:", err)
	}

	decoder := json.NewDecoder(file)
	conf := &Config{}
	err = decoder.Decode(&conf)
	if err != nil {
		fmt.Println("error:", err)
	}
	return conf
}
