package config

import (
	"fmt"
	"test-k-link-indonesia/models"
	"test-k-link-indonesia/packages/connection"
)

func Migration() {
	if connection.DB == nil {
		fmt.Println("Database connection is nil")
		return
	}

	err := connection.DB.AutoMigrate(
		&models.User{},
		&models.Gender{},
		&models.Level{},
		&models.Member{},
	)

	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	fmt.Println("Migration success")

}
