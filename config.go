package main

import (
	"io/ioutil"
	"path/filepath"

	"github.com/go-yaml/yaml"
)

type Config struct {
	Paths []string `yaml:"paths"`
	Mask string `yaml:"mask"`
	Commands []string `yaml:"commands"`
	Log string `yaml:"log"`
}

func NewConfig(path string) (*Config, error) {
	c := &Config{}

	filename, err := filepath.Abs(path)

	if err != nil {
		return nil, err
	}

	raw, err := ioutil.ReadFile(filename)

	if err != nil {
		return nil, err
	}

	err = yaml.Unmarshal(raw, &c)

	if err != nil {
		return nil, err
	}

	return c, nil
}