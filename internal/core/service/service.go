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

func (u *userService) GetByID(id string) (*models.User, error) {
	return u.userRepository.GetByID(id)
}

func (u *userService) FindUserByFullName(fullname string) (*models.FoodBeneficiary, error) {
	return u.userRepository.FindUserByFullName(fullname)
}

func (u *userService) FindUserByEmail(email string) (*models.FoodBeneficiary, error) {
	return u.userRepository.FindUserByEmail(email)
}

func (u *userService) FindUserByLocation(location string) (*models.FoodBeneficiary, error) {
	return u.userRepository.FindUserByLocation(location)
}

func (u *userService) CreateUser(user *models.FoodBeneficiary) (*models.FoodBeneficiary, error) {
	return u.userRepository.CreateUser(user)
}

func (u *userService) FindStaffByFullName(fullname string) (*models.KitchenStaff, error) {
	return u.userRepository.FindStaffByFullName(fullname)
}

func (u *userService) FindStaffByEmail(email string) (*models.KitchenStaff, error) {
	return u.userRepository.FindStaffByEmail(email)
}

func (u *userService) FindStaffByLocation(location string) (*models.KitchenStaff, error) {
	return u.userRepository.FindStaffByLocation(location)
}

func (u *userService) CreateStaff(user *models.KitchenStaff) (*models.KitchenStaff, error) {
	return u.userRepository.CreateStaff(user)
}
