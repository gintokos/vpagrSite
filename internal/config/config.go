package config

import (
	"encoding/json"
	"os"
)

type Config struct {
	Tconfig      TelegramBotConfig `json:"telegramBotConfig"`
	Sconfig      ServerConfig      `json:"serverConfig"`
	GrpcdbConfig GrpcdbConfig      `json:"grpcdbConfig"`
}

type TelegramBotConfig struct {
	Token         string `json:"token"`
	Link          string `json:"link"`
	Ttlusertoken  int    `json:"ttlusertoken"`
	Usertokensize int    `json:"usertokensize"`
}

type ServerConfig struct {
	Port string `json:"port"`
}

type GrpcdbConfig struct {
	Domen string `json:"domen"`
	Port  string `json:"port"`
}

func MustInitConfig(configpath string) Config {
	var cfg Config

	file, err := os.OpenFile(configpath, os.O_RDONLY, 0644)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	err = json.NewDecoder(file).Decode(&cfg)
	if err != nil {
		panic(err)
	}

	return cfg
}
