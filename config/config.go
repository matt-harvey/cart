package config

import (
	"github.com/caarlos0/env"
	_ "github.com/joho/godotenv/autoload"
)

var cfg *Config

type Config struct {
	ApiUrl string `env:"API_URL" envDefault:":3000"`
}

func init() {
	cfg = &Config{}
	err := env.Parse(cfg)
	if err != nil {
		panic(err)
	}
}

func ApiUrl() string {
	return cfg.ApiUrl
}
