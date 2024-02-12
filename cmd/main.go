package main

import (
	"log"

	"github.com/franzinBr/feedks-api/api"
	"github.com/franzinBr/feedks-api/data/db"
	"github.com/joho/godotenv"
)

func main() {

	if err := godotenv.Load(); err != nil {
		log.Fatal("Error in load .env file")
	}

	err := db.Init()
	defer db.CloseDB()

	if err != nil {
		log.Fatal(err)
	}

	db.Migrate()

	api.InitServer()

}
