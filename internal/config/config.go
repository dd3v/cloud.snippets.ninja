package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

//Config basic data structure for application configuration
type Config struct {
	BindAddr        string `yaml:"bind_addr"`
	LogLevel        string `yaml:"log_level"`
	DatabaseDNS     string `yaml:"database_dns"`
	TestDatabaseDNS string `yaml:"test_database_dns"`
	JWTSigningKey   string `yaml:"jwt_signing_key"`
}

func Load(path string) (*Config, error) {
	config := &Config{}
	file, err := ioutil.ReadFile(path)
	if err != nil {
		return config, err
	}
	err = yaml.Unmarshal(file, &config)
	return config, err
}
