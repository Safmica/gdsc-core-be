package database

import (
	"fmt"
	"gdsc-core-be/models"
	"log"
)

func DBMigration() {
	err := DB.AutoMigrate(
		&models.Role{},
		&models.User{},
		&models.Batch{},
		&models.Division{},
		&models.Member{},
		&models.Participant{},
		&models.Configure{},
		&models.Activity{},
		&models.FinalProject{},
		&models.Token{},
	)
	if err != nil {
		log.Println(err)
	}
	fmt.Println("Database migrated successfully")
}