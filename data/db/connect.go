package db

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var dbClient *gorm.DB

func Init() error {
	var err error

	dsn := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=disable",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_DATABASE"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
	)

	dbClient, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}

	conn, _ := dbClient.DB()

	if err := conn.Ping(); err != nil {
		return err
	}

	return nil

}

func GetDB() *gorm.DB {
	return dbClient
}

func CloseDB() {
	if dbClient != nil {
		conn, _ := dbClient.DB()
		conn.Close()
	}
}
