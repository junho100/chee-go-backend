package common

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Database struct {
	*gorm.DB
}

var DB *gorm.DB

func GetDB() *gorm.DB {
	return DB
}

func Init() *gorm.DB {
	db_username := os.Getenv("DB_USERNAME")
	db_password := os.Getenv("DB_PASSWORD")

	dsn := fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/test_db?charset=utf8mb4&parseTime=True&loc=Local", db_username, db_password)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalf("Error initialize database: %s", err)
	}

	DB = db

	return DB
}
