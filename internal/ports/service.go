package ports

import "github.com/decadevs/lunch-api/internal/core/models"

type UserService interface {
	GetByID(id string) (*models.User, error)
	FindUserByFullName(fullname string) (*models.FoodBeneficiary, error)
	FindUserByEmail(email string) (*models.FoodBeneficiary, error)
	FindUserByLocation(location string) (*models.FoodBeneficiary, error)
	CreateUser(user *models.FoodBeneficiary) (*models.FoodBeneficiary, error)
	FindStaffByFullName(fullname string) (*models.KitchenStaff, error)
	FindStaffByEmail(email string) (*models.KitchenStaff, error)
	FindStaffByLocation(location string) (*models.KitchenStaff, error)
	CreateStaff(user *models.KitchenStaff) (*models.KitchenStaff, error)
}
