package repository

import (
	"errors"
	"fmt"
	"github.com/decadevs/lunch-api/internal/core/models"
	"time"
)

// CreateFoodTimetable creates food in timetable
func (p *Postgres) CreateFoodTimetable(food models.Food) error {
	return p.DB.Create(&food).Error
}

// FindBrunchByDate finds brunch by date
func (p *Postgres) FindBrunchByDate(year int, month time.Month, day int) ([]models.Food, error) {
	var err error
	var food []models.Food
	if err = p.DB.Where("year = ?", year).Where("month = ?", month).Where("day = ?", day).
		Where("type = ?", "BRUNCH").Find(&food).Error; err != nil {
		return nil, errors.New(" food not found")
	}
	return food, nil
}

// FindDinnerByDate finds dinner by date
func (p *Postgres) FindDinnerByDate(year int, month time.Month, day int) ([]models.Food, error) {
	var err error
	var food []models.Food
	if err = p.DB.Where("year = ?", year).Where("month = ?", month).Where("day = ?", day).
		Where("type = ?", "DINNER").Find(&food).Error; err != nil {
		return nil, errors.New(" food not found")
	}
	return food, nil
}

func (p *Postgres) GetFoodByID(id string) (*models.Food, error) {
	food := &models.Food{}
	if err := p.DB.Where("ID = ?", id).First(food).Error; err != nil {
		return nil, err
	}
	return food, nil
}

func (p *Postgres) UpdateStatus(food []models.Food, status string) error {
	for i := 0; i < len(food); i++ {
		err := p.DB.Model(&models.Food{}).Where("id = ?", food[i].ID).Update("status", status).Error
		if err != nil {
			fmt.Println("error updating brunch status in database")
			return err
		}
	}
	return nil
}

func (p *Postgres) DeleteMeal(id string) error {
	var food models.Food
	err := p.DB.Where("id =?", id).Delete(&food).Error
	if err != nil {
		fmt.Println("error deleting food")
		return err
	}
	return nil
}

func (p *Postgres) UpdateMeal(id string, food models.Food) error {
	err := p.DB.Model(models.Food{}).Where("id = ?", id).Updates(&food).Error
	if err != nil {
		fmt.Println("error updating food")
		return err
	}

	return nil
}
