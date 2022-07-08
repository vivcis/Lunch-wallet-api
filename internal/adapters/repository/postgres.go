package repository

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDb(config *helpers.Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v sslmode=%v TimeZone=%v", config.DBHost, config.DBUser, config.DBPass, config.DBName, config.DBPort, config.DBMode, config.DBTimeZone)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	fmt.Println("Established database connection")
	return db, nil
}

func MigrateAll(db *gorm.DB) error {
	return db.AutoMigrate(&models.FoodBeneficiary{}, &models.KitchenStaff{}, &models.Admin{})
}
