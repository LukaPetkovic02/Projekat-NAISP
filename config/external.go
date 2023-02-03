package config

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Structure   string `yaml:"structure"`
	BloomFilter struct {
		Precision float64 `yaml:"precision"`
	} `yaml:"bloom_filter"`
	Cache struct {
		Size uint `yaml:"size"`
	} `yaml:"cache"`
	SkipList struct {
		MaxLevel uint `yaml:"max_level"`
	} `yaml:"skip_list"`
	Lsm struct {
		MaxLevel uint `yaml:"max_level"`
	}
	TokenBucket struct {
		Size uint   `yaml:"size"`
		Rate uint64 `yaml:"rate"`
	} `yaml:"token_bucket"`
}

func load() *Config {
	var c *Config = nil
	yamlFile, err := ioutil.ReadFile("config.yaml")
	if err != nil {
		fmt.Println("ERROR:", err)
	}
	err = yaml.Unmarshal(yamlFile, &c)
	if err != nil {
		fmt.Println("ERROR:", err)
	}

	return c
}
