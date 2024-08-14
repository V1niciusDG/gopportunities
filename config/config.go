package config

import (
	"fmt"

	"gorm.io/gorm"
)

var (
	db *gorm.DB
	logger *Logger
)

func Init() error {
	var err error

	// Initialize Postgres
	db, err = InitializePostgres()
	logger.Info("Running database")
	if err != nil{
		return fmt.Errorf("Error initializing postgres: %v", err)
	}

	return nil
}

func GetPostgres() *gorm.DB {
	return db
}

func GetLogger(p string) *Logger {
	// Initialize Logger
	logger = NewLogger(p)
	return logger
}

func GetExampleEnv() {
	err := createExampleEnv()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Println("example.env created successfully.")
	}
}