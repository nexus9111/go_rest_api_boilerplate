package database

import (
	"fmt"
	"github.com/nexus9111/go-rest-api-boilerplate/users/shared"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

var Context Database

type Database struct {
	DB        *gorm.DB
	Connected bool
}

func InitDatabase() error {
	host := os.Getenv("POSTGRES_HOST")
	user := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	dbname := os.Getenv("POSTGRES_DB")
	port := os.Getenv("POSTGRES_PORT")

	dsn := fmt.Sprintf("postgres://%v:%v@%v:%v/%v?sslmode=disable",
		user,
		password,
		host,
		port,
		dbname,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		return err
	}

	Context.DB = db
	Context.Connected = true

	err = Context.DB.AutoMigrate(&shared.User{})
	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}
