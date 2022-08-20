package server

import (
	"github.com/decadevs/lunch-api/internal/adapters/repository"
	"github.com/decadevs/lunch-api/internal/core/helpers"
	"github.com/decadevs/lunch-api/internal/core/models"
	"gorm.io/gorm"
	"log"
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

	image1 := models.Image{
		Url: "https://img-global.cpcdn.com/recipes/74a859bd76b085cb/751x532cq70/rolled-eba-and-egusi-soup-recipe-main-photo.jpg",
	}

	imageBrunch := []models.Image{
		image1,
	}

	image2 := models.Image{
		Url: "https://theblowfishgroup.com/purple/wp-content/uploads/2019/01/pur1001-1.jpg",
	}

	imageDinner := []models.Image{
		image2,
	}

	food1Brunch := models.Food{
		Model:     models.Model{},
		Name:      "Egusi and Swallow",
		Type:      "BRUNCH",
		AdminName: "Joseph Asuquo",
		Year:      2022,
		Month:     8,
		Day:       16,
		Weekday:   "Tuesday",
		Status:    "Not serving",
		Images:    imageBrunch,
		Kitchen:   "Edo-Tech Park",
	}

	food1Dinner := models.Food{
		Name:      "Fried Rice and Peppered Beef",
		Type:      "DINNER",
		AdminName: "Joseph Asuquo",
		Year:      2022,
		Month:     8,
		Day:       16,
		Weekday:   "Tuesday",
		Status:    "Not serving",
		Images:    imageDinner,
		Kitchen:   "Edo-Tech Park",
	}

	food2Brunch := models.Food{
		Name:      "Egusi and Swallow",
		Type:      "BRUNCH",
		AdminName: "Joseph Asuquo",
		Year:      2022,
		Month:     8,
		Day:       17,
		Weekday:   "Wednesday",
		Status:    "Not serving",
		Images:    imageBrunch,
		Kitchen:   "Edo-Tech Park",
	}

	food2Dinner := models.Food{
		Name:      "Fried Rice and Peppered Beef",
		Type:      "DINNER",
		AdminName: "Joseph Asuquo",
		Year:      2022,
		Month:     8,
		Day:       17,
		Weekday:   "Wednesday",
		Status:    "Not serving",
		Images:    imageDinner,
		Kitchen:   "Edo-Tech Park",
	}

	result := db.Where("type = ?", "BRUNCH").Find(&models.Food{})
	if result.RowsAffected < 1 {
		db.Create(&food1Brunch)
		db.Create(&food1Dinner)
		db.Create(&food2Brunch)
		db.Create(&food2Dinner)
	}

	return db, nil
}
