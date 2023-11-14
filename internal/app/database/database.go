package database

import (
	"fmt"
	"log"
	"os"

	"gorm.io/gorm"
	"gorm.io/driver/mysql"
	"github.com/nukkua/ra-chi/internal/app/models"
)


func SetupDatabase () * gorm.DB {
	
	
	requiredEnvVars := []string{"DB_USER","DB_PASSWORD", "DB_HOST", "DB_NAME","DB_PORT"}

	for _, envVar := range requiredEnvVars{
		if os.Getenv(envVar) == ""{
			log.Fatalf("This env variable is required but not defined: %s", envVar);
		}
	}

	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPass, dbHost, dbPort, dbName)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalf("Failed to connect to database");
	}

	db.AutoMigrate(&models.User{})

	return db
}
