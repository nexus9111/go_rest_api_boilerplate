package main

import (
	"GO_API_BOILERPLATE/config"
	"fmt"
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
