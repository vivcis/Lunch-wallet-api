package tests

import (
	"encoding/json"
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

func TestFoodBeneficiarySignUpEmailExists(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockDb := mocks.NewMockUserRepository(ctrl)

	r := &api.HTTPHandler{
		UserService: mockDb,
	}

	router := server.SetupRouter(r, mockDb)

	user := models.User{
		FullName:     "Orji Cecilia",
		Email:        "cece@decagon.dev",
		Password:     "password",
		PasswordHash: "",
		Location:     "ETP",
	}
	beneficiary := models.FoodBeneficiary{
		User:  user,
		Stack: "Golang",
	}

	newUser, err := json.Marshal(beneficiary)
	if err != nil {
		t.Fail()
	}
	mockDb.EXPECT().FindFoodBenefactorByEmail(beneficiary.Email).Return(&beneficiary, nil)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/api/v1/user/beneficiarysignup", strings.NewReader(string(newUser)))
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Contains(t, w.Body.String(), "email exists")
}

func TestFoodBeneficiarySignUpBadRequest(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockDb := mocks.NewMockUserRepository(ctrl)

	defer ctrl.Finish()

	r := &api.HTTPHandler{
		UserService: mockDb,
	}
	router := server.SetupRouter(r, mockDb)
	user := []models.User{
		{
			FullName:     "",
			Email:        "cece",
			Password:     "password",
			PasswordHash: "",
			Location:     "ETP",
		},
		{
			FullName:     "Dede",
			Email:        "cece@decagon.dev",
			Password:     "password",
			PasswordHash: "",
			Location:     "ETP",
		},
	}

	foodBeneficiary := models.FoodBeneficiary{
		User:  user[0],
		Stack: "java",
	}

	newUser, err := json.Marshal(foodBeneficiary)
	if err != nil {
		t.Fail()
	}
	beneficiary := models.FoodBeneficiary{
		User:  user[1],
		Stack: "gOLANG",
	}

	newUse, err := json.Marshal(beneficiary)
	if err != nil {
		t.Fail()
	}

	t.Run("Bad request", func(t *testing.T) {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", "/api/v1/user/beneficiarysignup", strings.NewReader(string(newUser)))
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)
		assert.Contains(t, w.Body.String(), "unable to bind request: validation error")
	})

	t.Run("Correct details", func(t *testing.T) {
		mockDb.EXPECT().FindFoodBenefactorByEmail(beneficiary.Email).Return(&beneficiary, nil)
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", "/api/v1/user/beneficiarysignup", strings.NewReader(string(newUse)))
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)
		assert.Contains(t, w.Body.String(), "email exist")
	})

}
