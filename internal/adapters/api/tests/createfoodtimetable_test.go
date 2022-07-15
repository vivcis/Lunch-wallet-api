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
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
)

func TestCreateFoodTimetableHandle(t *testing.T) {
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

	admin := models.Admin{
		User: user,
	}

	secret := os.Getenv("JWT_SECRET")
	accessClaims, _ := middleware.GenerateClaims(admin.Email)
	accToken, _ := middleware.GenerateToken(jwt.SigningMethodHS256, accessClaims, &secret)

	mockDb.EXPECT().TokenInBlacklist(gomock.Any()).Return(false).Times(2)
	mockDb.EXPECT().FindAdminByEmail(admin.Email).Return(&admin, nil).Times(2)
	t.Run("testing bad request", func(t *testing.T) {
		foodTimetable := &struct {
			Name    string `json:"name" binding:"required"`
			Type    string `json:"type" binding:"required"`
			Date    int    `json:"date" binding:"required"`
			Month   int    `json:"month" binding:"required"`
			Year    int    `json:"year" binding:"required"`
			Weekday string `json:"weekday" binding:"required"`
		}{
			Name: "Joseph A",
			Type: "brunch",
		}
		bytes, _ := json.Marshal(foodTimetable)
		rw := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodPost, "/api/v1/admin/createtimetable", strings.NewReader(string(bytes)))
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", *accToken))
		router.ServeHTTP(rw, req)
		assert.Equal(t, http.StatusBadRequest, rw.Code)
		assert.Contains(t, rw.Body.String(), "bad request")
	})

	t.Run("Successfully Created", func(t *testing.T) {
		foodTimetable := &struct {
			Name    string `json:"name" binding:"required"`
			Type    string `json:"type" binding:"required"`
			Date    int    `json:"date" binding:"required"`
			Month   int    `json:"month" binding:"required"`
			Year    int    `json:"year" binding:"required"`
			Weekday string `json:"weekday" binding:"required"`
		}{
			Name:    "Joseph A",
			Type:    "brunch",
			Date:    14,
			Month:   7,
			Year:    2022,
			Weekday: "Thursday",
		}

		bytes, _ := json.Marshal(foodTimetable)

		mockDb.EXPECT().CreateFoodTimetable(gomock.Any()).Return(nil)
		rw := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodPost, "/api/v1/admin/createtimetable", strings.NewReader(string(bytes)))
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", *accToken))
		router.ServeHTTP(rw, req)
		assert.Equal(t, http.StatusCreated, rw.Code)
		assert.Contains(t, rw.Body.String(), "Successfully Created")
	})
}
