package tests

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/decadevs/lunch-api/cmd/server"
	"github.com/decadevs/lunch-api/internal/adapters/api"
	"github.com/decadevs/lunch-api/internal/adapters/repository/mocks"
	"github.com/decadevs/lunch-api/internal/core/models"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestStaffSignUpEmailExists(t *testing.T) {
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
	staff := models.KitchenStaff{
		User: user,
	}

	newUser, err := json.Marshal(staff)
	if err != nil {
		t.Fail()
	}
	mockDb.EXPECT().FindKitchenStaffByEmail(staff.Email).Return(&staff, nil)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/api/v1/user/kitchenstaffsignup", strings.NewReader(string(newUser)))
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Contains(t, w.Body.String(), "email exists")
}

func TestStaffSignUpBadRequest(t *testing.T) {
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

	staff := models.KitchenStaff{
		User: user[0],
	}

	newUser, err := json.Marshal(staff)
	if err != nil {
		t.Fail()
	}
	kitchenStaff := models.KitchenStaff{
		User: user[1],
	}

	newUse, err := json.Marshal(kitchenStaff)
	if err != nil {
		t.Fail()
	}

	t.Run("Bad request", func(t *testing.T) {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", "/api/v1/user/kitchenstaffsignup", strings.NewReader(string(newUser)))
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)
		assert.Contains(t, w.Body.String(), "unable to bind request: validation error")
	})

	t.Run("Correct details", func(t *testing.T) {
		mockDb.EXPECT().FindKitchenStaffByEmail(kitchenStaff.Email).Return(&kitchenStaff, nil)
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", "/api/v1/user/kitchenstaffsignup", strings.NewReader(string(newUse)))
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)
		assert.Contains(t, w.Body.String(), "email exist")
	})

}
