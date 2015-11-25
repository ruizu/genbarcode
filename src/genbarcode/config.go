package main

import (
	"github.com/ruizu/gcfg"
)

type Config struct {
	Server ServerConfig
}

type ServerConfig struct {
	Host        string
	Environment string
}

func ReadConfig(c *Config, filePath string) bool {
	if err := gcfg.ReadFileInto(c, filePath); err != nil {
		return false
	}
	return true
}
