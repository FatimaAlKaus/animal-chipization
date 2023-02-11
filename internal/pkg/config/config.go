package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Postgres Postgres `yaml:"postgres"`
	Server   Server   `yaml:"server"`
}

func Load(file string) (Config, error) {
	rFile, err := os.Open(file)
	if err != nil {
		return Config{}, err
	}
	defer rFile.Close()

	var cfg Config
	err = yaml.NewDecoder(rFile).Decode(&cfg)

	return cfg, err
}

type Postgres struct {
	Port     int    `yaml:"port"`
	Host     string `yaml:"host"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Database string `yaml:"database"`
	SSLMode  string `yaml:"sslmode"`
}

type Server struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
}
