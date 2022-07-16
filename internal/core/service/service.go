package service

import (
	"github.com/decadevs/lunch-api/internal/core/models"
	"github.com/decadevs/lunch-api/internal/ports"
)

type userService struct {
	userRepository ports.UserRepository
}

type mailerService struct {
	mailerRepository ports.MailerRepository
}

func NewUserService(userRepository ports.UserRepository) ports.UserService {
	return &userService{
		userRepository: userRepository,
	}
}
func NewMailerService(mailerRepository ports.MailerRepository) ports.MailerService {
	return &mailerService{
		mailerRepository: mailerRepository,
	}
}

func (m *mailerService) SendMail(subject, body, to, Private, Domain string) error {
	return m.mailerRepository.SendMail(subject, body, to, Private, Domain)
}

//func (m *mailerService) GenerateNonAuthToken(UserEmail string, secret string) (*string, error) {
//	return m.mailerRepository.GenerateNonAuthToken(UserEmail, secret)
//}
//
//func (m *mailerService) DecodeToken(token, secret string) (string, error) {
//	return m.mailerRepository.DecodeToken(token, secret)
//}

func (u *userService) FindUserById(id string) (*models.FoodBeneficiary, error) {
	return u.userRepository.FindUserById(id)
}
func (u *userService) UserResetPassword(id, newPassword string) (*models.FoodBeneficiary, error) {
	return u.userRepository.UserResetPassword(id, newPassword)
}

func (u *userService) AdminResetPassword(id, newPassword string) (*models.Admin, error) {
	return u.userRepository.AdminResetPassword(id, newPassword)
}

func (u *userService) KitchenStaffResetPassword(id, newPassword string) (*models.KitchenStaff, error) {
	return u.userRepository.KitchenStaffResetPassword(id, newPassword)
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

func (u *userService) TokenInBlacklist(token *string) bool {
	return u.userRepository.TokenInBlacklist(token)
}

func (u *userService) AddTokenToBlacklist(email string, token string) error {
	return u.userRepository.AddTokenToBlacklist(email, token)
}
