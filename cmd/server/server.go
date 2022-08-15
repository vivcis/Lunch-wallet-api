package server

import (
	"log"

	"github.com/decadevs/lunch-api/internal/adapters/repository"
	"github.com/decadevs/lunch-api/internal/core/helpers"
	"github.com/decadevs/lunch-api/internal/core/models"
	"gorm.io/gorm"
)

func Run() (*gorm.DB, error) {
	err := helpers.Load()
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	db, err := repository.ConnectDb(&helpers.Config{
		DBUser:     helpers.Instance.DBUser,
		DBPass:     helpers.Instance.DBPass,
		DBHost:     helpers.Instance.DBHost,
		DBName:     helpers.Instance.DBName,
		DBPort:     helpers.Instance.DBPort,
		DBTimeZone: helpers.Instance.DBTimeZone,
		DBMode:     helpers.Instance.DBMode,
	})
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	err = repository.MigrateAll(db)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	user := &models.Admin{
		User: models.User{
			FullName: "Chinenye Ikpa",
			Email:    "chinenyei@decagonhq.com",
			Location: "ETP",
			Password: "Admin@123",
			IsActive: true,
			Token:    "",
		},
	}
	if err = user.HashPassword(); err != nil {
		return nil, err
	}
	db.Create(&user)
	return db, nil
}
