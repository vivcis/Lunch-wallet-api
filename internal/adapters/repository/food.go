package repository

import (
	"github.com/decadevs/lunch-api/internal/core/models"
)

// CreateFoodTimetable creates food in timetable
func (p *Postgres) CreateFoodTimetable(food models.Food) error {
	return p.DB.Create(food).Error

}
