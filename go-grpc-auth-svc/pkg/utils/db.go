package utils

import (
	"log"

	"github.com/eatrisno/go-grpc-auth-svc/pkg/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DBHandler struct {
	DB *gorm.DB
}

func DBInit(url string) DBHandler {
	db, err := gorm.Open(postgres.Open(url), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	db.AutoMigrate(&models.User{})

	return DBHandler{db}
}
