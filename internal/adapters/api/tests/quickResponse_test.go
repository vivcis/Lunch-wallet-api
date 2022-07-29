package tests

import (
	"github.com/decadevs/lunch-api/cmd/server"
	"github.com/decadevs/lunch-api/internal/adapters/api"
	"github.com/decadevs/lunch-api/internal/adapters/repository/mocks"
	"github.com/decadevs/lunch-api/internal/core/models"
	"github.com/golang/mock/gomock"
	"gorm.io/gorm"
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
		UserID:    "",
		UserEmail: "",
		Brunch:    false,
		Dinner:    false,
	}
}
