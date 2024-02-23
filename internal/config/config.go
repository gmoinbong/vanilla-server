package config

import (
	"log"
	"os"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Env        string `yaml:"env" env-default:"local"`
	HTTPServer `yaml:"http_server"`
	Storage    `yaml:"storage"`
}

type HTTPServer struct {
	Adress      string        `yaml:"address" env-default:"localhost:8080"`
	User        string        `yaml:"user"`
	Password    string        `yaml:"password"`
	Timeout     time.Duration `yaml:"4s"`
	IdleTimeout time.Duration `yaml:"30s"`
}

type Storage struct {
	DBName     int    `yaml:"DBName"`
	DBPassword string `yaml:"DBPassword"`
	DBHost     string `yaml:"DBHost"`
	DBPort     string `yaml:"DBPort"`
}

func MustLoadConfig() *Config {
	configPath := os.Getenv("CONFIG_PATH")

	if configPath == "" {
		log.Fatal("CONFIG_PATH isnt set")
	}

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("config file does not exist: %s", configPath)
	}

	var cfg Config

	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		log.Fatalf("cannot read config: %s", err)
	}
	return &cfg
}
