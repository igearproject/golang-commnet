package database

import (
	"fmt"
	"golang-comment/models"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	dsn := os.Getenv("DB_URI")
	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed connect to database")
	}
	fmt.Println("Successfully connected to database")
	database.AutoMigrate(&models.User{}, &models.Comment{})
	DB = database
}
