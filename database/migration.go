package database

import (
	"fmt"
	"waysbucks/models"
	"waysbucks/pkg/connection"
)

func RunMigration() {
	err := connection.DB.AutoMigrate(
		&models.User{},
		&models.Product{},
		&models.Toping{},
		&models.Cart{},
		&models.Transaction{},
	)

	if err != nil {
		fmt.Println(err)
		panic("Failed! Create Table to Database")
	}
	fmt.Println("Create Table Success")
}
