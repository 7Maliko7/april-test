package config

import (
	"context"

	"github.com/joho/godotenv"
	"github.com/sethvargo/go-envconfig"
)

type Config struct {
	ListenAddress  string `env:"LISTEN_ADDRESS"`
	RWDB           DB     `env:", prefix=RWDB_"`
	CarInfoAddress string `env:"CAR_INFO_ADDRESS"`
}

type DB struct {
	ConnectionString string `env:"CONNECTION_STRING"`
}

func New(path string) (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, err
	}

	c := Config{}
	if err := envconfig.Process(context.Background(), &c); err != nil {
		return nil, err
	}

	return &c, nil
}
