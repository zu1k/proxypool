package config

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/ghodss/yaml"
	"github.com/zu1k/proxypool/tool"
)

type Source struct {
	Type    string       `json:"type" yaml:"type"`
	Options tool.Options `json:"options" yaml:"options"`
}

type Config struct {
	Sources []Source `json:"sources" yaml:"sources"`
}

var SourceConfig = Config{}

func Parse(path string) (*Config, error) {
	fileData, err := readFile(path)
	if err != nil {
		return nil, err
	}

	err = yaml.Unmarshal(fileData, &SourceConfig)
	if err != nil {
		return nil, err
	}
	return &SourceConfig, nil
}

func readFile(path string) ([]byte, error) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return nil, err
	}
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	if len(data) == 0 {
		return nil, fmt.Errorf("Configuration file %s is empty", path)
	}

	return data, err
}
