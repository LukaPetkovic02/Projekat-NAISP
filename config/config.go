package config

func Get() *Config {
	var config = load()
	if config != nil {
		return config
	} else {
		return &Config{}
	}

}
