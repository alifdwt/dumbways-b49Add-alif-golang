package database

import (
	"fmt"
	"myapp/models"
	"myapp/pkg/postgres"
)

func RunMigration() {
	err := postgres.DB.AutoMigrate(&models.Paslon{}, &models.Voter{}, &models.Party{}, &models.User{})

	if err != nil {
		fmt.Println(err)
		panic("Migration failed")
	}

	fmt.Println("Migration Success")
}