package tests

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/decadevs/lunch-api/cmd/server"
	"github.com/decadevs/lunch-api/internal/adapters/api"
	"github.com/decadevs/lunch-api/internal/adapters/repository/mocks"
	"github.com/decadevs/lunch-api/internal/core/middleware"
	"github.com/decadevs/lunch-api/internal/core/models"
	"github.com/dgrijalva/jwt-go"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
	"time"
)

func TestGetBrunchHandle(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockDb := mocks.NewMockUserRepository(ctrl)

	r := &api.HTTPHandler{
		UserService: mockDb,
	}
	router := server.SetupRouter(r, mockDb)

	user := models.User{
		FullName: "Joseph A",
		Email:    "joseph@decagon.dev",
		IsActive: true,
	}

	benefactor := models.FoodBeneficiary{
		User:  user,
		Stack: "Golang",
	}
	year, month, day := time.Now().Date()
	food := &models.Food{
		Name:      "Egusi and Swallow",
		Type:      "BRUNCH",
		AdminName: "Joseph Asuquo",
		Year:      year,
		Month:     month,
		Day:       day,
		Weekday:   "Friday",
		Status:    "Not serve",
	}

	bytes, _ := json.Marshal(food)

	secret := os.Getenv("JWT_SECRET")
	accessClaims, _ := middleware.GenerateClaims(benefactor.Email)
	accToken, _ := middleware.GenerateToken(jwt.SigningMethodHS256, accessClaims, &secret)

	mockDb.EXPECT().TokenInBlacklist(gomock.Any()).Return(false).Times(2)
	mockDb.EXPECT().FindFoodBenefactorByEmail(benefactor.Email).Return(&benefactor, nil).Times(2)
	t.Run("testing bad request", func(t *testing.T) {

		mockDb.EXPECT().FindBrunchByDate(year, month, day).Return(nil, errors.New("internal server error"))
		rw := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodGet, "/api/v1/benefactor/brunch", strings.NewReader(string(bytes)))
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", *accToken))
		router.ServeHTTP(rw, req)
		assert.Equal(t, http.StatusInternalServerError, rw.Code)
		assert.Contains(t, rw.Body.String(), "internal server error")
	})

	t.Run("testing Successful request", func(t *testing.T) {

		mockDb.EXPECT().FindBrunchByDate(year, month, day).Return(food, nil)
		rw := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodGet, "/api/v1/benefactor/brunch", strings.NewReader(string(bytes)))
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", *accToken))
		router.ServeHTTP(rw, req)
		assert.Equal(t, http.StatusOK, rw.Code)
		assert.Contains(t, rw.Body.String(), "Brunch found")
	})
}
