package repository

import (
	"errors"
	"github.com/decadevs/lunch-api/internal/core/models"
	"time"
)

// FindFoodBenefactorByFullName finds a benefactor by full name
func (p *Postgres) FindFoodBenefactorByFullName(fullname string) (*models.FoodBeneficiary, error) {
	user := &models.FoodBeneficiary{}

	if err := p.DB.Where("fullname = ?", fullname).First(user).Error; err != nil {
		return nil, err
	}
	if !user.IsActive {
		return nil, errors.New("user inactive")
	}
	return user, nil
}

// FindFoodBenefactorByEmail finds a benefactor by email
func (p *Postgres) FindFoodBenefactorByEmail(email string) (*models.FoodBeneficiary, error) {
	//user := &models.FoodBeneficiary{}
	var err error
	var user *models.FoodBeneficiary
	if err = p.DB.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, errors.New(email + " does not exist" + " user not found")
	}
	return user, nil
}

// FindFoodBenefactorByLocation finds a benefactor by location
func (p *Postgres) FindFoodBenefactorByLocation(location string) (*models.FoodBeneficiary, error) {
	user := &models.FoodBeneficiary{}
	if err := p.DB.Where("location =?", location).First(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

// CreateFoodBenefactor creates a benefactor in the database
func (p *Postgres) CreateFoodBenefactor(user *models.FoodBeneficiary) (*models.FoodBeneficiary, error) {
	var err error
	user.CreatedAt = time.Now()
	user.IsActive = true
	err = p.DB.Create(user).Error
	return user, err
}
