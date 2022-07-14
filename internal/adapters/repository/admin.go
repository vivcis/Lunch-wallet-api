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
