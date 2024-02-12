package db

import "github.com/franzinBr/feedks-api/data/models"

func Migrate() {
	database := GetDB()

	database.AutoMigrate(
		&models.Role{},
		&models.User{},
		&models.Feedback{},
	)

}
