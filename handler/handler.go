package handler

import (
	"github.com/V1niciusDG/gopportunities/config"
	"gorm.io/gorm"
)

var (
	logger *config.Logger
	db *gorm.DB
)

func InitializeHandler() {
	logger = config.GetLogger("handler")
	db = config.GetPostgres()
}

