package db

import (
	"fmt"
	"myblogs/config"
	"myblogs/models"
)

func RunMigration() {
	err := config.DB.AutoMigrate(
		&models.User{},
	)

	if err != nil {
		panic(err)
	}

	fmt.Println("Migration Success")
}
