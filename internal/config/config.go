package config

//Config basic data structure for application configuration
type Config struct {
	BindAddr        string `toml:"bind_addr"`
	LogLevel        string `toml:"log_level"`
	DatabaseDNS     string `toml:"database_dns"`
	TestDatabaseDNS string `toml:"test_database_dns"`
	JWTSigningKey   string `toml:"jwt_signing_key"`
}

//NewConfig creates config structure and fills values by default
func NewConfig() *Config {
	return &Config{
		BindAddr: "8080",
		LogLevel: "debug",
	}
}
