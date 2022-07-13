package service

import (
	"github.com/decadevs/lunch-api/internal/core/models"
	"github.com/decadevs/lunch-api/internal/ports"
)

type userService struct {
	userRepository ports.UserRepository
}

func NewUserService(userRepository ports.UserRepository) ports.UserService {
	return &userService{
		userRepository: userRepository,
	}
}

func (u *userService) FindFoodBenefactorByFullName(fullname string) (*models.FoodBeneficiary, error) {
	return u.userRepository.FindFoodBenefactorByFullName(fullname)
}

func (u *userService) FindFoodBenefactorByEmail(email string) (*models.FoodBeneficiary, error) {
	return u.userRepository.FindFoodBenefactorByEmail(email)
}

func (u *userService) FindFoodBenefactorByLocation(location string) (*models.FoodBeneficiary, error) {
	return u.userRepository.FindFoodBenefactorByLocation(location)
}

func (u *userService) CreateFoodBenefactor(user *models.FoodBeneficiary) (*models.FoodBeneficiary, error) {
	return u.userRepository.CreateFoodBenefactor(user)
}

func (u *userService) FindKitchenStaffByFullName(fullname string) (*models.KitchenStaff, error) {
	return u.userRepository.FindKitchenStaffByFullName(fullname)
}

func (u *userService) FindKitchenStaffByEmail(email string) (*models.KitchenStaff, error) {
	return u.userRepository.FindKitchenStaffByEmail(email)
}

func (u *userService) FindKitchenStaffByLocation(location string) (*models.KitchenStaff, error) {
	return u.userRepository.FindKitchenStaffByLocation(location)
}

func (u *userService) CreateKitchenStaff(user *models.KitchenStaff) (*models.KitchenStaff, error) {
	return u.userRepository.CreateKitchenStaff(user)
}

func (u *userService) FindAdminByEmail(email string) (*models.Admin, error) {
	return u.userRepository.FindAdminByEmail(email)
}
