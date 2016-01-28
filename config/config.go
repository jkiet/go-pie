package config

import "gopkg.in/yaml.v2"

type Config struct {
	Section uint64   `yaml:"section"`
	Layout  []uint64 `yaml:"layout"`
}

func NewConfig() *Config {
	return &Config{}
}

func (c *Config) Read(s []byte) error {
	return yaml.Unmarshal(s, c)
}
