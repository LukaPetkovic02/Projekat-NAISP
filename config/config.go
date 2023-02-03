package config

func Get() *Config {
	var config = load()
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
		}

	}

}
