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
	mockMail := mocks.NewMockMailerRepository(ctrl)

	r := &api.HTTPHandler{
		UserService:   mockDb,
		MailerService: mockMail,
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
	Link := "<strong>Here is your reset <a href='http://localhost:8080/api/v1/user/beneficiaryresetpassword/cad4fc7b-b819-4ec0-aff4-5cefefd7f8ee'>link</a></strong>"
	mockMail.EXPECT().SendMail("forgot Password", Link, "test@gmail.com", privateAPIKey, yourDomain).Return(nil)

}
