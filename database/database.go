package database

import (
	"aida/models"
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func ConnectDB() *gorm.DB {
	DB, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		log.Default().Print("Failed to cnnecct to database ")
		return nil
	}
	DB.AutoMigrate(&models.User{},
		&models.Challenges{},
		&models.Course{},
		&models.Student{},
		&models.Teacher{},
		&models.Preference{},
		&models.Product{},
		&models.Product{},
	)
	return DB
}
