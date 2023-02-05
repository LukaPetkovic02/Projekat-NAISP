package config

func load() *Config {
	var config = loadExternal()
	if config != nil {
		return config
	} else {
		return &Config{
			WalSegment: 3,
			Structure:  "multiple-files",
			BloomFilter: BloomFilter{
				Precision: 0.01,
			},
			Cache: Cache{
				Size: 10,
			},
			SkipList: SkipList{
				MaxLevel: 10,
				Height:   5,
			},
			Lsm: Lsm{
				MaxLevel: 5,
			},
			TokenBucket: TokenBucket{
				Size: 10,
				Rate: 1000,
			},
			Memtable: Memtable{
				Size:      120,
				Threshold: 10,
				Use:       "skip-list",
			},
			Summary: Summary{
				BlockSize: 5,
			},
			Btree: Btree{
				MaxNode: 2,
			},
		}
	}
}

var Values *Config = load()
