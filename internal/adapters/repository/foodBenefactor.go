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

// FindUserById finds a benefactor by location
func (p *Postgres) FindUserById(id string) (*models.FoodBeneficiary, error) {
	user := &models.FoodBeneficiary{}
	if err := p.DB.Where("id =?", id).First(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

// UserResetPassword resets a benefactor's password
func (p *Postgres) UserResetPassword(id, newPassword string) (*models.FoodBeneficiary, error) {
	user := &models.FoodBeneficiary{}
	if err := p.DB.Model(user).Where("id =?", id).Update("password_hash", newPassword).Error; err != nil {
		return nil, err
	}
	return user, nil
}

// CreateFoodBenefactor creates a benefactor in the database
func (p *Postgres) CreateFoodBenefactor(user *models.FoodBeneficiary) (*models.FoodBeneficiary, error) {
	var err error
	user.CreatedAt = time.Now()
	user.IsActive = false
	err = p.DB.Create(user).Error
	return user, err
}

//FoodBeneficiaryEmailVerification verifies the beneficiary email address
func (p *Postgres) FoodBeneficiaryEmailVerification(id string) (*models.FoodBeneficiary, error) {
	user := &models.FoodBeneficiary{}
	if err := p.DB.Model(user).Where("id =?", id).Update("is_active", true).Error; err != nil {
		return nil, err
	}
	return user, nil
}

// FindFoodBenefactorMealRecord finds a benefactor meal record
func (p *Postgres) FindFoodBenefactorMealRecord(email, date string) (*models.MealRecords, error) {
	user := &models.MealRecords{}
	err := p.DB.Where("user_email =? AND meal_date = ?", email, date).Last(user).Error
	if user.UserEmail == "" {
		return nil, err
	}
	return user, nil
}

// CreateFoodBenefactorBrunchMealRecord creates a benefactor meal record in the database
func (p *Postgres) CreateFoodBenefactorBrunchMealRecord(user *models.FoodBeneficiary) error {
	var err error
	record := &models.MealRecords{
		Model:     models.Model{},
		MealDate:  time.Now().Format("2006-01-02"),
		UserID:    user.ID,
		UserEmail: user.Email,
		Brunch:    true,
		Dinner:    false,
	}

	err = p.DB.Create(record).Error
	return err
}

//UpdateFoodBenefactorBrunchMealRecord updates the beneficiary meal record
func (p *Postgres) UpdateFoodBenefactorBrunchMealRecord(email string) error {
	user := &models.MealRecords{}
	if err := p.DB.Model(user).Where("user_email =?", email).Update("dinner", true).Error; err != nil {
		return err
	}
	return nil
}

// CreateFoodBenefactorDinnerMealRecord creates a benefactor meal record in the database
func (p *Postgres) CreateFoodBenefactorDinnerMealRecord(user *models.FoodBeneficiary) error {
	var err error
	record := &models.MealRecords{
		Model:     models.Model{},
		MealDate:  time.Now().Format("2006-01-02"),
		UserID:    user.ID,
		UserEmail: user.Email,
		Brunch:    false,
		Dinner:    true,
	}

	err = p.DB.Create(record).Error
	return err
}
