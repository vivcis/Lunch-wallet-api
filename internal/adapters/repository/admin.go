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
	user.IsActive = true
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
