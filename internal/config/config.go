package config

//Config basic data structure for application configuration
type Config struct {
	BindAddr     string `toml:"bind_addr"`
	LogLevel     string `toml:"log_level"`
	DatabaseURL  string `toml:"database_url"`
	DatabaseName string `toml:"database_name"`
}

//NewConfig creates config structure and fills values by default
func NewConfig() *Config {
	return &Config{
		BindAddr: "8080",
		LogLevel: "debug",
	}
}
