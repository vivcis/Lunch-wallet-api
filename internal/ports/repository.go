package ports

import (
	"github.com/decadevs/lunch-api/internal/core/models"
)

type UserRepository interface {
	FindFoodBenefactorByFullName(fullname string) (*models.FoodBeneficiary, error)
	FindFoodBenefactorByEmail(email string) (*models.FoodBeneficiary, error)
	FindFoodBenefactorByLocation(location string) (*models.FoodBeneficiary, error)
	CreateFoodBenefactor(user *models.FoodBeneficiary) (*models.FoodBeneficiary, error)
	FindKitchenStaffByFullName(fullname string) (*models.KitchenStaff, error)
	FindKitchenStaffByEmail(email string) (*models.KitchenStaff, error)
	FindKitchenStaffByLocation(location string) (*models.KitchenStaff, error)
	CreateKitchenStaff(user *models.KitchenStaff) (*models.KitchenStaff, error)
	FindAdminByEmail(email string) (*models.Admin, error)
	TokenInBlacklist(token *string) bool
	AddTokenToBlacklist(email string, token string) error
	FindUserById(id string) (*models.FoodBeneficiary, error)
	UserResetPassword(id, newPassword string) (*models.FoodBeneficiary, error)
}

// MailerRepository interface to implement mailing service
type MailerRepository interface {
	SendMail(subject, body, to, Private, Domain string) error
}
