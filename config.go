package main

import (
	"io/ioutil"

	"github.com/BurntSushi/toml"
)

// Config represents main configuration of application
type Config struct {
	Port     string         `toml:"port"`
	Database configDatabase `toml:"database"`
}

type configDatabase struct {
	Type     string `toml:"type"`
	Host     string `toml:"host"`
	Port     string `toml:"port"`
	User     string `toml:"user"`
	Password string `toml:"password"`
	Name     string `toml:"name"`
	SslMode  string `toml:"sslmode"`
}

func initConfig(configPath string, conf *Config) error {
	data, err := ioutil.ReadFile(configPath)
	if err != nil {
		return err
	}

	if _, err := toml.Decode(string(data), conf); err != nil {
		return err
	}

	return nil
}
