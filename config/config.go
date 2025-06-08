package config

import (
	"os"

	"go.uber.org/zap"
	"gopkg.in/yaml.v3"
)

type Config struct {
	Redis RedisConfig `yaml:"redis"`
}

type RedisConfig struct {
	Host string `yaml:"host"`
}

func InitConfig(logger *zap.Logger) Config {
	 yamlFile, err := os.ReadFile("config.yaml")
	 if err != nil {
		logger.Fatal(err.Error())
	 }
	 var config Config
	 err = yaml.Unmarshal(yamlFile, &config)
	 if err != nil {
		logger.Fatal(err.Error())
	 }
	 return config
}