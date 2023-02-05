package config

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v3"
)

type Config struct {
	WalSegment  int         `yaml:"walSegment"`
	Structure   string      `yaml:"structure"`
	BloomFilter BloomFilter `yaml:"bloomFilter"`
	Cache       Cache       `yaml:"cache"`
	SkipList    SkipList    `yaml:"skipList"`
	Lsm         Lsm         `yaml:"lsm"`
	TokenBucket TokenBucket `yaml:"tokenBucket"`
	Memtable    Memtable    `yaml:"memtable"`
	Summary     Summary     `yaml:"summary"`
	Btree       Btree       `yaml:"btree"`
}

type Btree struct {
	MaxNode int `yaml:"maxNode"`
}
type BloomFilter struct {
	Precision float64 `yaml:"precision"`
}
type Cache struct {
	Size uint `yaml:"size"`
}
type SkipList struct {
	MaxLevel uint `yaml:"maxLevel"`
	Height   int  `yaml:"height"`
}
type Lsm struct {
	MaxLevel uint `yaml:"maxLevel"`
}
type TokenBucket struct {
	Size uint   `yaml:"size"`
	Rate uint64 `yaml:"rate"`
}

type Memtable struct {
	Size      uint   `yaml:"size"`
	Threshold uint64 `yaml:"threshold"`
	Use       string `yaml:"use"`
}

type Summary struct {
	BlockSize int `yaml:"blockSize"`
}

func loadExternal() *Config {
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
