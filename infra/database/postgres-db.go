package utils

import (
	"log"
	"os"

	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"

	"blog-golang/domain"
)

func ConnectPg() *gorm.DB {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error while loading environment variables" + err.Error())
	}

	dsn := os.Getenv("dsn")

	db, err := gorm.Open("postgres", dsn)
	if err != nil {
		log.Fatal("Error durring opening database connection" + err.Error())
		panic(err)
	}

	// defer db.Close()

	db.AutoMigrate(&domain.Author{}) // create database migrations from a struct
	return db

}
