package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"path/filepath"
)

type Config struct {
	Fields map[string]string `yaml:"config"`
}

func (c Config) Get(fieldName string) string {
	return c.Fields[fieldName]
}

func Load() Config {
	filename, _ := filepath.Abs("./config.yml")
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	var cfg Config
	err = yaml.Unmarshal(file, &cfg)
	if err != nil {
		panic(err)
	}

	return cfg
}
