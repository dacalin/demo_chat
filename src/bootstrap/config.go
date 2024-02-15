package bootstrap

import (
	"fmt"
	"github.com/caarlos0/env/v10"
)

type Config struct {
	WsPort                int    `env:"WS_PORT"`
	WsPingIntervalSeconds int    `env:"WS_PING_INTERVAL_SECONDS"`
	RedisHost             string `env:"REDIS_HOST"`
	RedisPort             int    `env:"REDIS_PORT"`
}

func GetConfig() Config {

	config := Config{}
	if err := env.Parse(&config); err != nil {
		fmt.Printf("%+v\n", err)
	}

	fmt.Printf("%+v\n", config)

	return config
}
