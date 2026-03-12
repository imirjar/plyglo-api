package config

import (
	"log"

	"github.com/caarlos0/env/v11"
)

type Config struct {
	Port   string `env:"PORT"`
	DBConn string `env:"DB_CONN"`
}

func New() *Config {
	cfg := Config{}

	// Всегда читаем переменные окружения (перезаписывают файл)
	err := cfg.readEnv()
	if err != nil {
		log.Print(err)
		return nil
	}

	return &cfg
}

func (cfg *Config) readEnv() error {
	err := env.Parse(cfg)
	if err != nil {
		log.Printf("%+v\n", err)
	}
	return err

}
