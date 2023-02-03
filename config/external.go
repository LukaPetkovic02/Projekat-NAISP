package config

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Structure   string      `yaml:"structure"`
	BloomFilter BloomFilter `yaml:"bloomFilter"`
	Cache       Cache       `yaml:"cache"`
	SkipList    SkipList    `yaml:"skipList"`
	Lsm         Lsm         `yaml:"lsm"`
	TokenBucket TokenBucket `yaml:"tokenBucket"`
}
type BloomFilter struct {
	Precision float64 `yaml:"precision"`
}
type Cache struct {
	Size uint `yaml:"size"`
}
type SkipList struct {
	MaxLevel uint `yaml:"maxLevel"`
}
type Lsm struct {
	MaxLevel uint `yaml:"maxLevel"`
}
type TokenBucket struct {
	Size uint   `yaml:"size"`
	Rate uint64 `yaml:"rate"`
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
