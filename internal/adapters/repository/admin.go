package repository

import (
	"errors"
	"github.com/decadevs/lunch-api/internal/core/models"
)

// FindAdminByEmail finds a user by email
func (p *Postgres) FindAdminByEmail(email string) (*models.Admin, error) {
	admin := &models.Admin{}
	if err := p.DB.Where("email = ?", email).First(admin).Error; err != nil {
		return nil, errors.New(email + "not an Admin")
	}

	return admin, nil
}

// AdminResetPassword resets a benefactor's password
func (p *Postgres) AdminResetPassword(id, newPassword string) (*models.KitchenStaff, error) {
	user := &models.KitchenStaff{}
	if err := p.DB.Model(user).Where("id =?", id).Update("password_hash", newPassword).Error; err != nil {
		return nil, err
	}
	return user, nil
}
