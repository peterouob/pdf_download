package config

import (
	"github.com/joho/godotenv"
	"os"
)

type config struct {
	RedisAddr string
	Addr      string
}

var Cfg *config

func LoadConfig() {
	if err := godotenv.Load(); err != nil {
		panic("Error to load the env file")
	}
	Cfg = &config{
		RedisAddr: os.Getenv("redis"),
		Addr:      os.Getenv("ADDR"),
	}
}
