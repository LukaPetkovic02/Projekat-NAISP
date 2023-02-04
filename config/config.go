package config

func load() *Config {
	var config = loadExternal()
	if config != nil {
		return config
	} else {
		return &Config{
			Structure: "multiple-files",
			BloomFilter: BloomFilter{
				Precision: 0.01,
			},
			Cache: Cache{
				Size: 10,
			},
			SkipList: SkipList{
				MaxLevel: 7,
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
		}
	}
}

var Values *Config = load()
