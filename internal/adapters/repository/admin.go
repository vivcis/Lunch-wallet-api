package repository

import (
	"errors"
	"github.com/decadevs/lunch-api/internal/core/models"
	"time"
)

// CreateAdmin creates an Admin in the database
func (p *Postgres) CreateAdmin(user *models.Admin) (*models.Admin, error) {
	var err error
	user.CreatedAt = time.Now()
	user.IsActive = false
	err = p.DB.Create(user).Error
	return user, err
}

// FindAdminByEmail finds a user by email
func (p *Postgres) FindAdminByEmail(email string) (*models.Admin, error) {
	admin := &models.Admin{}
	if err := p.DB.Where("email = ?", email).First(admin).Error; err != nil {
		return nil, errors.New(email + "not an Admin")
	}

	return admin, nil
}

// AdminResetPassword resets a benefactor's password
func (p *Postgres) AdminResetPassword(id, newPassword string) (*models.Admin, error) {
	user := &models.Admin{}
	if err := p.DB.Model(user).Where("id =?", id).Update("password_hash", newPassword).Error; err != nil {
		return nil, err
	}
	return user, nil
}

//AdminEmailVerification verifies the admin email address
func (p *Postgres) AdminEmailVerification(id string) (*models.Admin, error) {
	user := &models.Admin{}
	if err := p.DB.Model(user).Where("id =?", id).Update("is_active", true).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (p *Postgres) AdminBlockFoodBeneficiary(userID string) error {
	var user *models.FoodBeneficiary
	err := p.DB.Model(user).Where("id = ?", userID).Update("is_block", true).Error

	if err != nil {
		return err
	}
	return nil
}

func (p *Postgres) AdminRemoveFoodBeneficiary(userID string) error {
	user := models.FoodBeneficiary{}
	err := p.DB.Model(&user).Where("id = ?", userID).Delete(&user).Error
	if err != nil {
		return err
	}
	return nil
}
