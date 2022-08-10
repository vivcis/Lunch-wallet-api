package ports

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/decadevs/lunch-api/internal/core/models"
	"mime/multipart"
)

type UserService interface {
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
	KitchenStaffResetPassword(id, newPassword string) (*models.KitchenStaff, error)
	AdminResetPassword(id, newPassword string) (*models.Admin, error)
	CreateFoodTimetable(food models.Food) error
	CreateAdmin(user *models.Admin) (*models.Admin, error)
	FindBrunchByDate(year int, month int, day int) ([]models.Food, error)
	FindDinnerByDate(year int, month int, day int) ([]models.Food, error)
	FoodBeneficiaryEmailVerification(id string) (*models.FoodBeneficiary, error)
	KitchenStaffEmailVerification(id string) (*models.KitchenStaff, error)
	AdminEmailVerification(id string) (*models.Admin, error)
	FindAllFoodBeneficiary(pagination *models.Pagination) ([]models.UserDetails, error)
	FindFoodBenefactorMealRecord(email, date string) (*models.MealRecords, error)
	CreateFoodBenefactorBrunchMealRecord(user *models.FoodBeneficiary) error
	CreateFoodBenefactorDinnerMealRecord(user *models.FoodBeneficiary) error
	UpdateFoodBenefactorBrunchMealRecord(email string) error
	UpdateFoodBenefactorDinnerMealRecord(email string) error
	GetFoodByID(id string) (*models.Food, error)
	UpdateStatus(food []models.Food, status string) error
	SearchFoodBeneficiary(text string, pagination *models.Pagination) ([]models.UserDetails, error)
	GetTotalUsers() (int, error)
	UpdateMeal(id string, food models.Food) error
	DeleteMeal(id string) error
	FindAllFoodByDate(year int, month int, day int) ([]models.Food, error)
	CreateNotification(notification models.Notification) error
	FindNotificationByDate(year int, month int, day int) ([]models.Notification, error)
	GetFoodBenefactorById(id string) (*models.FoodBeneficiary, error)
	AdminBlockFoodBeneficiary(userID string) error
	AdminRemoveFoodBeneficiary(userID string) error
	NumberOfBlockedBeneficiary() (int64, error)
	GetBlockedBeneficiary() ([]models.FoodBeneficiary, error)
}

// MailerService interface to implement mailing service
type MailerService interface {
	SendMail(subject, body, to, Private, Domain string) error
	GenerateNonAuthToken(UserEmail string, secret string) (*string, error)
	DecodeToken(token, secret string) (string, error)
}

// AWSService interface to implement AWS
type AWSService interface {
	UploadFileToS3(h *session.Session, file multipart.File, fileName string, size int64) (string, error)
}
