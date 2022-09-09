package database

import (
	"asacoco/models"
	"asacoco/pkg/mysql"
	"log"
)

func RunMigration() {
	err := mysql.DB.AutoMigrate(&models.User{})

	if err != nil {
		log.Fatal(err)
		log.Fatal("Migration Failed")
	}

	log.Println("Migration Success")
}