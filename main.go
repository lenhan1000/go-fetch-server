package main

import (
	"go-fetch-server/api"
	"go-fetch-server/config"
)

func main() {
	config := &config.Config{}
	config.Read()

	server := &api.Server{}
	server.Init(config.Server, config.Database)
	server.Run(":3000")
}
