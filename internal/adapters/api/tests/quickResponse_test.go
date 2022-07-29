package tests

import (
	"encoding/json"
	"fmt"
	"github.com/decadevs/lunch-api/cmd/server"
	"github.com/decadevs/lunch-api/internal/adapters/api"
	"github.com/decadevs/lunch-api/internal/adapters/repository/mocks"
	"github.com/decadevs/lunch-api/internal/core/middleware"
	"github.com/decadevs/lunch-api/internal/core/models"
	"github.com/dgrijalva/jwt-go"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
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
		UserID:    beneficiary.ID,
		UserEmail: beneficiary.Email,
		Brunch:    true,
		Dinner:    false,
	}
	secret := os.Getenv("JWT_SECRET")
	accessClaims, _ := middleware.GenerateClaims(beneficiary.Email)
	accToken, _ := middleware.GenerateToken(jwt.SigningMethodHS256, accessClaims, &secret)
	t.Run("testing bad request", func(t *testing.T) {
		mockDb.EXPECT().TokenInBlacklist(gomock.Any()).Return(false)
		mockDb.EXPECT().FindFoodBenefactorByEmail(beneficiary.Email).Return(&beneficiary, nil)
		mockDb.EXPECT().FindFoodBenefactorMealRecord(beneficiary.Email, date).Return(mealRecord, nil)

		bytes, _ := json.Marshal(mealRecord)
		rw := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodGet, "/api/v1/benefactor/qrbrunch", strings.NewReader(string(bytes)))
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", *accToken))
		router.ServeHTTP(rw, req)
		assert.Equal(t, http.StatusBadRequest, rw.Code)
		assert.Contains(t, rw.Body.String(), "brunch already served")

	})

}
