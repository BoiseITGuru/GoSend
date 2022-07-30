package main

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	Version   string         `mapstructure:"version"`
	WebServer WebServer      `mapstructure:"server"`
	DB        DatabaseConfig `mapstructure:"database"`
	Aws       AwsConfig      `mapstructure:"aws"`
}

type DatabaseConfig struct {
	GormEngine     string `mapstructure:"gorm_engine"`
	GormConnection string `mapstructure:"gorm_connection"`
}

type WebServer struct {
	Port string `mapstructure:"port"`
}

type AwsConfig struct {
	Region string       `mapstructure:"region"`
	Kms    AwsKmsConfig `mapstructure:"kms"`
}

type AwsKmsConfig struct {
	EncryptionKey string `mapstructure:"encryption_key"`
	JwtKey        string `mapstructure:"jwt_key"`
}

var AppConfig *Config

func LoadAppConfig() {
	log.Println("Loading Server Configurations...")
	// Read file path
	viper.AddConfigPath(".")
	// set config file and path
	viper.SetConfigName("config")
	viper.SetConfigType("json")
	// reading the config file
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}

	err = viper.Unmarshal(&AppConfig)
	if err != nil {
		log.Fatal(err)
	}
}
