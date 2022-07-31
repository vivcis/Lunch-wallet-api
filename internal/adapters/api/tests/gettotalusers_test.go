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
	"strconv"
	"strings"
	"testing"
)

func TestTotalNumberOfFoodBeneficiaries(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockDb := mocks.NewMockUserRepository(ctrl)

	r := &api.HTTPHandler{
		UserService: mockDb,
	}

	router := server.SetupRouter(r, mockDb)

	user1 := models.User{
		Model:    models.Model{},
		FullName: "Orji Cecilia",
		Email:    "cecilia.orji@decagon.dev",
		Location: "ETP",
		Password: "cece",
		IsActive: false,
		Status:   "active",
		Avatar:   "image.png",
	}

	foodBeneficiary := models.FoodBeneficiary{
		User:  user1,
		Stack: "GO",
	}
	count, _ := strconv.Atoi(foodBeneficiary.FullName)

	kitchenStaff := models.KitchenStaff{
		User: user1,
	}
	bytes, _ := json.Marshal(user1)

	secret := os.Getenv("JWT_SECRET")
	accessClaims, _ := middleware.GenerateClaims(kitchenStaff.Email)
	accToken, _ := middleware.GenerateToken(jwt.SigningMethodHS256, accessClaims, &secret)

	t.Run("testing bad request", func(t *testing.T) {
		mockDb.EXPECT().TokenInBlacklist(gomock.Any()).Return(false)
		mockDb.EXPECT().FindKitchenStaffByEmail(kitchenStaff.Email).Return(&kitchenStaff, nil)
		mockDb.EXPECT().GetTotalUsers().Return(count, errors.New("internal server error"))
		rw := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodGet, "/api/v1/staff/gettotalusers", strings.NewReader(string(bytes)))
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", *accToken))
		router.ServeHTTP(rw, req)
		assert.Equal(t, http.StatusInternalServerError, rw.Code)
		assert.Contains(t, rw.Body.String(), "internal server error")
	})

	t.Run("testing success", func(t *testing.T) {
		mockDb.EXPECT().TokenInBlacklist(gomock.Any()).Return(false)
		mockDb.EXPECT().FindKitchenStaffByEmail(kitchenStaff.Email).Return(&kitchenStaff, nil)
		mockDb.EXPECT().GetTotalUsers().Return(count, nil)
		rw := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodGet, "/api/v1/staff/gettotalusers", strings.NewReader(string(bytes)))
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", *accToken))
		router.ServeHTTP(rw, req)
		assert.Equal(t, http.StatusOK, rw.Code)
		assert.Contains(t, rw.Body.String(), "Total number")
	})
}
