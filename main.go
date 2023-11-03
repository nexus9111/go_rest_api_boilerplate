package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/nexus9111/go-rest-api-boilerplate/config"
	"github.com/nexus9111/go-rest-api-boilerplate/config/database"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println(err)
		panic("Error loading .env file")
	}

	err = database.InitDatabase()
	if err != nil {
		fmt.Println(err)
		panic("Database failed to connect")
	}

	var server config.Server
	server.Init()

	err = server.Run()
	if err != nil {
		fmt.Println(err)
		panic("Server failed to start")
	}
}
