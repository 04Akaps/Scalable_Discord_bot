package config

import (
	"github.com/naoina/toml"
	"go.uber.org/zap"
	"os"
)

type Config struct {
	ServiceInfo struct {
		Env string
	}

	MySQL map[string]struct {
		Database string
		Host     string
		User     string
		Password string

		Collections map[string]string
	}

	Logger *zap.Logger
}

func NewConfig(file string) *Config {
	c := new(Config)

	if f, err := os.Open(file); err != nil {
		panic(err)
	} else {
		if err = toml.NewDecoder(f).Decode(c); err != nil {
			panic(err)
		} else {
			return c
		}
	}
}
