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
)

func TestSearchBeneficiaries(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockDb := mocks.NewMockUserRepository(ctrl)

	r := &api.HTTPHandler{
		UserService: mockDb,
	}

	router := server.SetupRouter(r, mockDb)

	user := models.User{
		Model:    models.Model{},
		FullName: "Orji Cecilia",
		Email:    "cecilia.orji@decagon.dev",
		Location: "ETP",
		Password: "cece",
		IsActive: false,
		Status:   "active",
		Avatar:   "image.png",
	}

	kitchenStaff := models.KitchenStaff{
		User: user,
	}
	beneficiary := models.UserDetails{
		Stack: "Golang",
	}
	foodBeneficiary := []models.UserDetails{
		beneficiary,
	}

	page := models.Pagination{
		Page:  1,
		Limit: 10,
		Sort:  "created_at asc",
	}

	bytes, _ := json.Marshal(user)

	secret := os.Getenv("JWT_SECRET")
	accessClaims, _ := middleware.GenerateClaims(kitchenStaff.Email)
	accToken, _ := middleware.GenerateToken(jwt.SigningMethodHS256, accessClaims, &secret)

	t.Run("test bad request", func(t *testing.T) {
		mockDb.EXPECT().TokenInBlacklist(gomock.Any()).Return(false)
		mockDb.EXPECT().FindKitchenStaffByEmail(kitchenStaff.Email).Return(&kitchenStaff, nil)
		mockDb.EXPECT().SearchFoodBeneficiary(gomock.Any(), gomock.Any()).Return(nil, errors.New("record not found"))
		rw := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodGet, "/api/v1/staff/searchbeneficiary/python", strings.NewReader(string(bytes)))
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", *accToken))
		router.ServeHTTP(rw, req)
		assert.Equal(t, http.StatusInternalServerError, rw.Code)
		assert.Contains(t, rw.Body.String(), "internal server error")
	})

	t.Run("test successful search", func(t *testing.T) {
		mockDb.EXPECT().TokenInBlacklist(gomock.Any()).Return(false)
		mockDb.EXPECT().FindKitchenStaffByEmail(kitchenStaff.Email).Return(&kitchenStaff, nil)
		mockDb.EXPECT().SearchFoodBeneficiary(beneficiary.Stack, &page).Return(foodBeneficiary, nil)
		rw := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodGet, "/api/v1/staff/searchbeneficiary/Golang", strings.NewReader(string(bytes)))
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", *accToken))
		router.ServeHTTP(rw, req)
		assert.Equal(t, http.StatusOK, rw.Code)
		assert.Contains(t, rw.Body.String(), "information gotten")
	})
}
