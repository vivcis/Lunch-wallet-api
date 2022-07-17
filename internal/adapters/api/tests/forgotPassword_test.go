package tests

import (
	"github.com/decadevs/lunch-api/cmd/server"
	"github.com/decadevs/lunch-api/internal/adapters/api"
	"github.com/decadevs/lunch-api/internal/adapters/repository/mocks"
	"github.com/decadevs/lunch-api/internal/core/models"
	"github.com/golang/mock/gomock"
	"os"
	"testing"
)

func TestBuyerSendForgotPasswordEMailHandler(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockDb := mocks.NewMockUserRepository(ctrl)

	r := &api.HTTPHandler{
		UserService: mockDb,
	}

	router := server.SetupRouter(r, mockDb)

	resetPassword := struct {
		Email string `json:"email"`
	}{
		Email: "test@testmail.com",
	}
	beneficiary := models.FoodBeneficiary{
		User: models.User{
			Email:        resetPassword.Email,
			PasswordHash: "passwordHash",
		},
		Stack: "golang",
	}
	privateAPIKey := os.Getenv("MAILGUN_API_KEY")
	yourDomain := os.Getenv("DOMAIN_STRING")

	mockDb.EXPECT().FindFoodBenefactorByEmail(resetPassword.Email).Return(beneficiary, nil)

}
