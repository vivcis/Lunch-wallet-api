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
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
)

// TestLogoutFoodBeneficiary tests beneficiary logout handler
func TestLogoutFoodBeneficiary(t *testing.T) {
	err := godotenv.Load("../../.env")
	if err != nil {
		log.Println(err.Error())
	}
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockDb := mocks.NewMockUserRepository(ctrl)

	r := &api.HTTPHandler{
		UserService: mockDb,
	}
	router := server.SetupRouter(r, mockDb)

	accClaim, _ := middleware.GenerateClaims("cece@decagon.dev")
	secret := os.Getenv("JWT_SECRET")
	acc, err := middleware.GenerateToken(jwt.SigningMethodHS256, accClaim, &secret)
	if err != nil {
		t.Fail()
	}
	user := models.User{
		Email: "cece@decagon.dev",
		Token: *acc,
	}
	beneficiary := &models.FoodBeneficiary{
		User: user,
	}

	mockDb.EXPECT().TokenInBlacklist(acc).Return(false).Times(1)
	mockDb.EXPECT().FindFoodBenefactorByEmail(beneficiary.Email).Return(beneficiary, nil).Times(1)
	mockDb.EXPECT().AddTokenToBlacklist(beneficiary.Email, beneficiary.Token).Return(nil).Times(1)

	rw := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/api/v1/benefactor/beneficiarylogout", nil)
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", *acc))
	router.ServeHTTP(rw, req)
	assert.Equal(t, http.StatusOK, rw.Code)
	assert.Contains(t, rw.Body.String(), "successfully")

	t.Run("test bad request", func(t *testing.T) {
		user := models.User{
			Email: "cece@decagon.dev",
			Token: "*acc",
		}
		beneficiary := &models.FoodBeneficiary{
			User: user,
		}

		newUser, err := json.Marshal(beneficiary)
		if err != nil {
			t.Fail()
		}
		rw := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodPost, "/api/v1/benefactor/beneficiarylogout", strings.NewReader(string(newUser)))
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", "hdhhdhddhdh"))
		router.ServeHTTP(rw, req)
		assert.Equal(t, http.StatusUnauthorized, rw.Code)
		//assert.Contains(t, rw.Body.String(), "authorize access token error")
	})
}

// TestLogoutKitchenStaff tests kitchen staff logout handler
func TestLogoutKitchenStaff(t *testing.T) {
	err := godotenv.Load("../../.env")
	if err != nil {
		log.Println(err.Error())
	}
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockDb := mocks.NewMockUserRepository(ctrl)

	r := &api.HTTPHandler{
		UserService: mockDb,
	}
	router := server.SetupRouter(r, mockDb)

	accClaim, _ := middleware.GenerateClaims("cece@decagon.dev")
	secret := os.Getenv("JWT_SECRET")
	acc, err := middleware.GenerateToken(jwt.SigningMethodHS256, accClaim, &secret)
	if err != nil {
		t.Fail()
	}
	user := models.User{
		Email: "cece@decagon.dev",
		Token: *acc,
	}
	staff := &models.KitchenStaff{
		User: user,
	}

	mockDb.EXPECT().TokenInBlacklist(acc).Return(false).Times(1)
	mockDb.EXPECT().FindKitchenStaffByEmail(staff.Email).Return(staff, nil).Times(1)
	mockDb.EXPECT().AddTokenToBlacklist(staff.Email, staff.Token).Return(nil).Times(1)

	rw := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/api/v1/staff/kitchenstafflogout", nil)
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", *acc))
	router.ServeHTTP(rw, req)
	assert.Equal(t, http.StatusOK, rw.Code)
	assert.Contains(t, rw.Body.String(), "successfully")

	t.Run("test bad request", func(t *testing.T) {
		user := models.User{
			Email: "cece@decagon.dev",
			Token: "*acc",
		}
		kitchenStaff := &models.FoodBeneficiary{
			User: user,
		}

		newUser, err := json.Marshal(kitchenStaff)
		if err != nil {
			t.Fail()
		}
		rw := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodPost, "/api/v1/benefactor/beneficiarylogout", strings.NewReader(string(newUser)))
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", "hdhhdhddhdh"))
		router.ServeHTTP(rw, req)
		assert.Equal(t, http.StatusUnauthorized, rw.Code)
		//assert.Contains(t, rw.Body.String(), "authorize access token error")
	})
}
