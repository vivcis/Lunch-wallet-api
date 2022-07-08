package repository

import "gorm.io/gorm"

type Postgres struct {
	DB *gorm.DB
}

func NewUser(DB *gorm.DB) ports.UserRepository {
	return &Postgres{DB}
}

func (p *Postgres) GetByID(id string) (*models.User, error) {
	var user models.User
	if err := p.DB.Where("id = ?", id).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
