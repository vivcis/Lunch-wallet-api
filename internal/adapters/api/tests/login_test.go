package tests

import (
	"encoding/json"
	"fmt"
	"github.com/decadevs/lunch-api/cmd/server"
	"github.com/decadevs/lunch-api/internal/adapters/api"
	"github.com/decadevs/lunch-api/internal/adapters/repository/mocks"
	"github.com/decadevs/lunch-api/internal/core/models"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

// TestLoginKitchenStaffHandler test kitchen staff login handle
func TestLoginKitchenStaffHandler(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockDb := mocks.NewMockUserRepository(ctrl)

	r := &api.HTTPHandler{
		UserService: mockDb,
	}

	router := server.SetupRouter(r, mockDb)

	t.Run("testing bad request", func(t *testing.T) {
		KitchenStaffLoginRequest := &struct {
			Email    string `json:"email" binding:"required"`
			Password string `json:"password" binding:"required"`
		}{
			Email:    "",
			Password: "12345566666",
		}
		bytes, _ := json.Marshal(KitchenStaffLoginRequest)
		rw := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodPost, "/api/v1/user/kitchenstafflogin", strings.NewReader(string(bytes)))

		router.ServeHTTP(rw, req)
		assert.Equal(t, http.StatusBadRequest, rw.Code)
		assert.Contains(t, rw.Body.String(), "bad request")
	})

	t.Run("find kitchen staff by email", func(t *testing.T) {

		kitchenStaff := &models.KitchenStaff{
			User: models.User{
				Email:    "mike123@decagon.dev",
				Password: "12345566666",
				IsActive: true,
			},
		}
		_ = kitchenStaff.HashPassword()

		KitchenStaffLoginRequest := &struct {
			Email    string `json:"email" binding:"required"`
			Password string `json:"password" binding:"required"`
		}{
			Email:    "mike123@decagon.dev",
			Password: "12345566666",
		}
		mockDb.EXPECT().FindKitchenStaffByEmail(kitchenStaff.Email).Return(kitchenStaff, nil)
		bytes, _ := json.Marshal(KitchenStaffLoginRequest)
		rw := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodPost, "/api/v1/user/kitchenstafflogin", strings.NewReader(string(bytes)))

		router.ServeHTTP(rw, req)
		assert.Equal(t, http.StatusOK, rw.Code)
		fmt.Println(rw.Body.String())
		assert.Contains(t, rw.Body.String(), KitchenStaffLoginRequest.Email)
	})
}

// TestLoginFoodBenefactorHandler tests benefactor login handler
func TestLoginFoodBenefactorHandler(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockDb := mocks.NewMockUserRepository(ctrl)

	r := &api.HTTPHandler{
		UserService: mockDb,
	}

	router := server.SetupRouter(r, mockDb)

	t.Run("testing bad request", func(t *testing.T) {

		benefactorLoginRequest := &struct {
			Email    string `json:"email" binding:"required"`
			Password string `json:"password" binding:"required"`
		}{
			Email:    "",
			Password: "12345566666",
		}
		bytes, _ := json.Marshal(benefactorLoginRequest)
		rw := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodPost, "/api/v1/user/benefactorlogin", strings.NewReader(string(bytes)))

		router.ServeHTTP(rw, req)
		assert.Equal(t, http.StatusBadRequest, rw.Code)
		assert.Contains(t, rw.Body.String(), "bad request")
	})

	t.Run("find food benefactor by email", func(t *testing.T) {

		benefactor := &models.FoodBeneficiary{
			User: models.User{
				Email:    "mike123@decagon.dev",
				Password: "12345566666",
				IsActive: true,
			},
		}
		_ = benefactor.HashPassword()

		benefactorLoginRequest := &struct {
			Email    string `json:"email" binding:"required"`
			Password string `json:"password" binding:"required"`
		}{
			Email:    "mike123@decagon.dev",
			Password: "12345566666",
		}
		mockDb.EXPECT().FindFoodBenefactorByEmail(benefactor.Email).Return(benefactor, nil)
		bytes, _ := json.Marshal(benefactorLoginRequest)
		rw := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodPost, "/api/v1/user/benefactorlogin", strings.NewReader(string(bytes)))

		router.ServeHTTP(rw, req)
		assert.Equal(t, http.StatusOK, rw.Code)
		fmt.Println(rw.Body.String())
		assert.Contains(t, rw.Body.String(), benefactorLoginRequest.Email)
	})
}

// TestLoginAdminHandler tests Admin Login handler
func TestLoginAdminHandler(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockDb := mocks.NewMockUserRepository(ctrl)

	r := &api.HTTPHandler{
		UserService: mockDb,
	}

	router := server.SetupRouter(r, mockDb)

	t.Run("testing bad request", func(t *testing.T) {

		adminLoginRequest := &struct {
			Email    string `json:"email" binding:"required"`
			Password string `json:"password" binding:"required"`
		}{
			Email:    "",
			Password: "12345566666",
		}
		bytes, _ := json.Marshal(adminLoginRequest)
		rw := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodPost, "/api/v1/user/adminlogin", strings.NewReader(string(bytes)))

		router.ServeHTTP(rw, req)
		assert.Equal(t, http.StatusBadRequest, rw.Code)
		assert.Contains(t, rw.Body.String(), "bad request")
	})

	t.Run("find admin by email", func(t *testing.T) {

		admin := &models.Admin{
			User: models.User{
				Email:    "mike123@decagon.dev",
				Password: "12345566666",
				IsActive: true,
			},
		}
		_ = admin.HashPassword()

		adminLoginRequest := &struct {
			Email    string `json:"email" binding:"required"`
			Password string `json:"password" binding:"required"`
		}{
			Email:    "mike123@decagon.dev",
			Password: "12345566666",
		}
		mockDb.EXPECT().FindAdminByEmail(admin.Email).Return(admin, nil)
		bytes, _ := json.Marshal(adminLoginRequest)
		rw := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodPost, "/api/v1/user/adminlogin", strings.NewReader(string(bytes)))

		router.ServeHTTP(rw, req)
		assert.Equal(t, http.StatusOK, rw.Code)
		fmt.Println(rw.Body.String())
		assert.Contains(t, rw.Body.String(), adminLoginRequest.Email)
	})
}
