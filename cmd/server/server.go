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
	user := &models.FoodBeneficiary{
		User: models.User{
			FullName:     "jdoe",
			Email:        "a@decagon.dev",
			Location:     "ETP",
			PasswordHash: "$223456788878878989",
			IsActive:     true,
			Token:        "",
		},
		Stack: "GO",
	}
	user1 := &models.FoodBeneficiary{
		User: models.User{
			FullName:     "Chuks",
			Email:        "chuks@decagon.dev",
			Location:     "ETP",
			PasswordHash: "$223456788878878989",
			IsActive:     true,
			Token:        "",
		},
		Stack: "GO",
	}
	db.Create(&user)
	db.Create(&user1)
	return db, nil
}
