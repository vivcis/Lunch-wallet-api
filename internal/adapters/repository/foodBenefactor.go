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

// FindUserById finds a benefactor by ID
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

func (p *Postgres) GetFoodBenefactorById(id string) (*models.FoodBeneficiary, error) {
	var user *models.FoodBeneficiary
	if err := p.DB.Model(&user).Where("id = ?", id).First(&user).Error; err != nil {
		return nil, err

	}
	return user, nil
}

// FoodBeneficiaryEmailVerification verifies the beneficiary email address
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

// FindFoodBenefactorMealRecord finds a benefactor meal record
func (p *Postgres) FindActiveUsersByMonth(date string) (*models.MealRecords, error) {
	user := &models.MealRecords{}

	err := p.DB.Where("user_email =? AND meal_date = ?", date).Last(user).Error
	if user.UserEmail == "" {
		return nil, err
	}
	return user, nil
}

// FindFoodBenefactorQRCodeMealRecord finds a benefactor QR meal record
func (p *Postgres) FindFoodBenefactorQRCodeMealRecord(mealId, userId string) (*models.QRCodeMealRecords, error) {
	record := &models.QRCodeMealRecords{}
	err := p.DB.Where("meal_id = ? AND user_id = ?", mealId, userId).First(record).Error
	if record.MealId == "" {
		return nil, err
	}
	return record, nil
}

// CreateFoodBenefactorQRMealRecord creates a benefactor meal record in the database
func (p *Postgres) CreateFoodBenefactorQRMealRecord(mealRecord *models.QRCodeMealRecords) error {
	var err error
	record := &models.QRCodeMealRecords{
		Model:  models.Model{},
		MealId: mealRecord.MealId,
		UserId: mealRecord.UserId,
	}
	err = p.DB.Create(record).Error
	return err
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

// UpdateFoodBenefactorBrunchMealRecord updates the beneficiary meal record
func (p *Postgres) UpdateFoodBenefactorBrunchMealRecord(email string) error {
	user := &models.MealRecords{}
	if err := p.DB.Model(user).Where("user_email =?", email).Update("brunch", true).Error; err != nil {
		return err
	}
	return nil
}

// UpdateFoodBenefactorDinnerMealRecord updates the beneficiary meal record
func (p *Postgres) UpdateFoodBenefactorDinnerMealRecord(email string) error {
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

// FindAllFoodBeneficiary finds and list all food beneficiaries
func (p *Postgres) FindAllFoodBeneficiary(pagination *models.Pagination) ([]models.UserDetails, error) {
	var foodBeneficiary []models.UserDetails
	offset := (pagination.Page - 1) * pagination.Limit
	queryBuider := p.DB.Limit(pagination.Limit).Offset(offset).Order(pagination.Sort)
	result := queryBuider.Model(&models.FoodBeneficiary{}).Find(&foodBeneficiary).Error
	if result != nil {
		return nil, result
	}
	return foodBeneficiary, nil
}

// SearchFoodBeneficiary searches for food beneficiary
func (p *Postgres) SearchFoodBeneficiary(text string, pagination *models.Pagination) ([]models.UserDetails, error) {
	var users []models.FoodBeneficiary
	offset := (pagination.Page - 1) * pagination.Limit
	queryBuider := p.DB.Limit(pagination.Limit).Offset(offset).Order(pagination.Sort)
	err := queryBuider.Where("full_name = ?", text).Or("location = ?", text).Or("stack = ?", text).Find(&users).Error
	var result []models.UserDetails
	for i, _ := range users {
		userDetails := models.UserDetails{
			FullName: users[i].FullName,
			Stack:    users[i].Stack,
			Location: users[i].Location,
		}
		result = append(result, userDetails)
	}
	return result, err
}

func (p *Postgres) NumberOfBlockedBeneficiary() (int64, error) {
	var user []models.FoodBeneficiary
	var number int64
	if err := p.DB.Where("is_block =?", true).Find(&user).Count(&number).Error; err != nil {
		return 0, err
	}
	return number, nil
}

func (p *Postgres) GetBlockedBeneficiary() ([]models.FoodBeneficiary, error) {
	var user []models.FoodBeneficiary
	if err := p.DB.Where("is_block =?", true).Find(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

// GetAllFoodBeneficiaries  returns all the sellers in the updated database
func (p *Postgres) GetAllFoodBeneficiaries() ([]models.FoodBeneficiary, error) {
	var foodBeneficiary []models.FoodBeneficiary
	err := p.DB.Model(&models.FoodBeneficiary{}).Find(&foodBeneficiary).Error
	if err != nil {
		return nil, err
	}
	return foodBeneficiary, nil
}
