package repository

import (
	"errors"
	"github.com/decadevs/lunch-api/internal/core/models"
	"time"
)

// FindKitchenStaffByFullName finds a kitchen staff by full name
func (p *Postgres) FindKitchenStaffByFullName(fullname string) (*models.KitchenStaff, error) {
	user := &models.KitchenStaff{}

	if err := p.DB.Where("fullname = ?", fullname).First(user).Error; err != nil {
		return nil, err
	}
	if !user.IsActive {
		return nil, errors.New("user inactive")
	}
	return user, nil
}

// FindKitchenStaffByEmail finds a kitchen staff by email
func (p *Postgres) FindKitchenStaffByEmail(email string) (*models.KitchenStaff, error) {
	var err error
	var user *models.KitchenStaff
	if err = p.DB.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, errors.New(email + " does not exist" + " user not found")
	}
	return user, nil
}

// FindKitchenStaffByLocation finds a kitchen staff by location
func (p *Postgres) FindKitchenStaffByLocation(location string) (*models.KitchenStaff, error) {
	user := &models.KitchenStaff{}
	if err := p.DB.Where("location =?", location).First(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

// CreateKitchenStaff creates a kitchen staff in the database
func (p *Postgres) CreateKitchenStaff(user *models.KitchenStaff) (*models.KitchenStaff, error) {
	var err error
	user.CreatedAt = time.Now()
	user.IsActive = true
	err = p.DB.Create(user).Error
	return user, err
}
