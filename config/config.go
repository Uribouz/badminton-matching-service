package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Redis RedisConfig `yaml:"redis"`
	Mongo MongoConfig `yaml:"mongo"`
}

type RedisConfig struct {
	Host string `yaml:"host"`
}

type MongoConfig struct {
	URI      string `yaml:"uri"`
	Database string `yaml:"database"`
}

func InitConfig() Config {
	yamlFile, err := os.ReadFile("configs/config.yml")
	if err != nil {
		panic(err.Error())
	}
	var config Config
	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		panic(err.Error())
	}
	return config
}
