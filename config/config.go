package config

import (
	"github.com/apex/log"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
	"os"
	"strconv"
)

type Config struct {
	Host       string
	Port       int
	ServerHost string
	DB         gorm.DB
}

// var configs *Config

func InitConfig() *Config {
	err := godotenv.Load(".env")
	if err != nil {
		panic(".env is not loaded properly")
	}
	c := new(Config)
	c.Port, _ = strconv.Atoi(os.Getenv(`PORT`))
	c.ServerHost = os.Getenv(`HOST`)

	log.Info("All config & secrets set")
	return c
}
