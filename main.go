package main

import (
	"fmt"
	"github.com/nexus9111/go-rest-api-boilerplate/config"
)

func main() {
	var server config.Server
	server.Init()
	err := server.Run()
	if err != nil {
		fmt.Println(err)
		panic("Server failed to start")
	}
}
