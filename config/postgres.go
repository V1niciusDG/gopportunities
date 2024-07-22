package config

import (
	"fmt"
	"os"

	"github.com/V1niciusDG/gopportunities/schemas"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitializePostgres() (*gorm.DB, error) {
	logger := GetLogger("postgres")
	err := godotenv.Load()
	if err != nil {
		logger.debug.Fatalf("Error to load file .env %v", err)
	}

	dbHost := os.Getenv("DB_HOST")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbPort := os.Getenv("DB_PORT")


	dbPath := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", dbHost, dbUser, dbPass, dbName, dbPort)

	// Create DB and connect
	db, err := gorm.Open(postgres.Open(dbPath), &gorm.Config{})
	if err != nil{
		logger.Errorf("postgres opening error: %v", err)
		return nil, err
	}

	var exists bool
	result := db.Raw("SELECT EXISTS(SELECT 1 FROM pg_database WHERE datname = 'postgo')").Scan(&exists)
	if result.Error != nil {
		logger.Errorf("error checking database existence: %v", result.Error)
	}

	if !exists {
		logger.Info("database not found")
	}

	// Migrate the Schema
	err = db.AutoMigrate(&schemas.Opening{})
	if err != nil{
		logger.Errorf("postgres automigration error: %v", err)
		return nil, err
	}
	
	return db, nil
}