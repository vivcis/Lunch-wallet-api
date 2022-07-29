package tests

import (
	"github.com/decadevs/lunch-api/cmd/server"
	"github.com/decadevs/lunch-api/internal/adapters/api"
	"github.com/decadevs/lunch-api/internal/adapters/repository/mocks"
	"github.com/decadevs/lunch-api/internal/core/middleware"
	"github.com/decadevs/lunch-api/internal/core/models"
	"github.com/dgrijalva/jwt-go"
	"github.com/golang/mock/gomock"
	"gorm.io/gorm"
	"os"
	"testing"
	"time"
)

func TestBeneficiaryQRBrunch(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockDb := mocks.NewMockUserRepository(ctrl)

	r := &api.HTTPHandler{
		UserService: mockDb,
	}
	router := server.SetupRouter(r, mockDb)
	userModel := models.Model{
		ID:        "userID",
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: gorm.DeletedAt{},
	}
	user := models.User{
		Model:    userModel,
		FullName: "Michael Gbenle",
		Email:    "michael.gbenle@decagon.dev",
		Location: "Edo Tech Park",
		IsActive: true,
	}
	beneficiary := models.FoodBeneficiary{
		User:  user,
		Stack: "Golang",
	}
	model := models.Model{
		ID:        "hello",
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: gorm.DeletedAt{},
	}
	date := time.Now().Format("2006-01-02")
	mealRecord := &models.MealRecords{
		Model:     model,
		MealDate:  date,
		UserID:    "userid",
		UserEmail: "michael.gbenle@decagon.dev",
		Brunch:    true,
		Dinner:    false,
	}
	secret := os.Getenv("JWT_SECRET")
	accessClaims, _ := middleware.GenerateClaims(admin.Email)
	accToken, _ := middleware.GenerateToken(jwt.SigningMethodHS256, accessClaims, &secret)

}
