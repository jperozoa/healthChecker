package config

import "github.com/fulldump/goconfig"

type LoggerConfig struct {
	Level string `usage:"Levels: DEBUG,INFO,WARN,ERROR"`
}

// Config defines benchmark configuration.
type Config struct {
	URL      string `usage:"url""`
	Insecure bool   `usage:"Avoid validating server certificate""`
}

// New Configuration Constructor
func New() *Config {
	cfg := &Config{
		Insecure: false,
	}
	goconfig.Read(cfg)
	return cfg
}
