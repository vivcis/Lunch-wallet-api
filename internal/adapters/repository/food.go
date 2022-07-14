package repository

import (
	"errors"
	"github.com/decadevs/lunch-api/internal/core/models"
	"time"
)

// CreateFoodTimetable creates food in timetable
func (p *Postgres) CreateFoodTimetable(food models.Food) error {
	return p.DB.Create(&food).Error
}

// FindBrunchByDate finds brunch by date
func (p *Postgres) FindBrunchByDate(year int, month time.Month, day int) (*models.Food, error) {
	var err error
	var food *models.Food
	if err = p.DB.Where("year = ?", year).Where("month = ?", month).Where("day = ?", day).
		Where("type = ?", "brunch").First(&food).Error; err != nil {
		return nil, errors.New(" food not found")
	}
	return food, nil
}
