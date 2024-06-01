package main

import (
	"mnc-finance/config"
	"mnc-finance/internal"
)

func main() {
	cfg := config.InitConfig()
	server := internal.InitServer(cfg)
	server.Run()
}
