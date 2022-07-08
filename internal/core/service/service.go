package service

import "github.com/decadevs/lunch-api/internal/core/models"

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
