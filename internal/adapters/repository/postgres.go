package repository

import (
	"fmt"
	"github.com/decadevs/lunch-api/internal/core/helpers"
	"github.com/decadevs/lunch-api/internal/core/models"
	"github.com/decadevs/lunch-api/internal/ports"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

type Postgres struct {
	DB *gorm.DB
}

func NewUser(DB *gorm.DB) ports.UserRepository {
	return &Postgres{DB}
}

func ConnectDb(config *helpers.Config) (*gorm.DB, error) {
	var dsn string
	databaseURL := os.Getenv("DATABASE_URL")
	if databaseURL == "" {
		dsn = fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v sslmode=%v TimeZone=%v", config.DBHost, config.DBUser, config.DBPass, config.DBName, config.DBPort, config.DBMode, config.DBTimeZone)
	} else {
		dsn = databaseURL
	}

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	fmt.Println("Established database connection")
	return db, nil
}

func MigrateAll(db *gorm.DB) error {
	return db.AutoMigrate(&models.FoodBeneficiary{}, &models.KitchenStaff{}, &models.Admin{}, &models.Food{}, &models.Blacklist{}, &models.MealRecords{})
}
