package repository

import (
	"errors"
	"time"

	"github.com/decadevs/lunch-api/internal/core/models"
	"github.com/decadevs/lunch-api/internal/ports"
	"gorm.io/gorm"
)

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

// FindUserByFullName finds a user by the fullname
func (p *Postgres) FindUserByFullName(fullname string) (*models.FoodBeneficiary, error) {
	user := &models.FoodBeneficiary{}

	if err := p.DB.Where("fullname = ?", fullname).First(user).Error; err != nil {
		return nil, err
	}
	if !user.IsActive {
		return nil, errors.New("user inactive")
	}
	return user, nil
}

func (p *Postgres) FindUserByEmail(email string) (*models.FoodBeneficiary, error) {
	//user := &models.FoodBeneficiary{}
	var err error
	var user *models.FoodBeneficiary
	if err = p.DB.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, errors.New(email + " does not exist" + " user not found")
	}
	return user, nil
}

func (p *Postgres) FindUserByLocation(location string) (*models.FoodBeneficiary, error) {
	user := &models.FoodBeneficiary{}
	if err := p.DB.Where("location =?", location).First(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (p *Postgres) CreateUser(user *models.FoodBeneficiary) (*models.FoodBeneficiary, error) {
	var err error
	user.CreatedAt = time.Now()
	user.IsActive = true
	err = p.DB.Create(user).Error
	return user, err
}

func (p *Postgres) FindStaffByFullName(fullname string) (*models.KitchenStaff, error) {
	user := &models.KitchenStaff{}

	if err := p.DB.Where("fullname = ?", fullname).First(user).Error; err != nil {
		return nil, err
	}
	if !user.IsActive {
		return nil, errors.New("user inactive")
	}
	return user, nil
}

func (p *Postgres) FindStaffByEmail(email string) (*models.KitchenStaff, error) {
	var err error
	var user *models.KitchenStaff
	if err = p.DB.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, errors.New(email + " does not exist" + " user not found")
	}
	return user, nil
}

func (p *Postgres) FindStaffByLocation(location string) (*models.KitchenStaff, error) {
	user := &models.KitchenStaff{}
	if err := p.DB.Where("location =?", location).First(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (p *Postgres) CreateStaff(user *models.KitchenStaff) (*models.KitchenStaff, error) {
	var err error
	user.CreatedAt = time.Now()
	user.IsActive = true
	err = p.DB.Create(user).Error
	return user, err
}
