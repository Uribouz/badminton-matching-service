package config

import (
	"log"
	"strings"

	"github.com/spf13/viper"
)

type Config struct {
	App   AppConfig   `mapstructure:"app"`
	Redis RedisConfig `mapstructure:"redis"`
	Mongo MongoConfig `mapstructure:"mongo"`
}

type AppConfig struct {
	AllowOrigins []string `mapstructure:"allow_origins"`
}
type RedisConfig struct {
	Host string `mapstructure:"host"`
}

type MongoConfig struct {
	URI      string `mapstructure:"uri"`
	Database string `mapstructure:"database"`
}

func InitConfig() Config {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./configs")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		log.Printf("Error reading config file: %v", err)
	}

	viper.SetEnvPrefix("BADMINTON")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		log.Fatalf("Unable to decode config into struct: %v", err)
	}

	log.Printf("Final config URI: %v", config.Mongo.URI)
	return config
}
