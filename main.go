package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/nexus9111/go-rest-api-boilerplate/config"
)

func main() {
	godotenv.Load(".env")
	var server config.Server
	server.Init()
	err := server.Run()
	if err != nil {
		fmt.Println(err)
		panic("Server failed to start")
	}
}
