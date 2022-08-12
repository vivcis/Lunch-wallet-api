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

func TestFoodBeneficiarySignupWithIncorrectDetails(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockDb := mocks.NewMockUserRepository(ctrl)
	mockMail := mocks.NewMockMailerRepository(ctrl)

	r := &api.HTTPHandler{
		UserService:   mockDb,
		MailerService: mockMail,
	}

	router := server.SetupRouter(r, mockDb)

	user := []models.User{
		{
			FullName:     "Orji Cecilia",
			Email:        "cecilia.orjidecago.dev",
			Password:     "password",
			PasswordHash: "",
			Location:     "ETP",
		},
		{
			FullName:     "Orji Cecilia",
			Email:        "cecilia.orji@decagon.dev",
			Password:     "password",
			PasswordHash: "",
			Location:     "ETP",
		},
		{
			FullName:     "Orji Cecilia",
			Email:        "cecilia.orji@decago.dev",
			Password:     "password",
			PasswordHash: "",
			Location:     "ETP",
		},
	}

	t.Run("bad request", func(t *testing.T) {
		beneficiary := models.FoodBeneficiary{
			User:  user[0],
			Stack: "Golang",
		}

		newUser, err := json.Marshal(beneficiary)
		if err != nil {
			t.Fail()
		}

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", "/api/v1/user/beneficiarysignup", strings.NewReader(string(newUser)))
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)
		fmt.Println(w.Code)
		fmt.Println(w.Body.String())
		assert.Contains(t, w.Body.String(), fmt.Sprintf("validation failed on field 'Email', condition: email, actual: %s", beneficiary.Email))
	})

	t.Run("invalid decadev email", func(t *testing.T) {
		beneficiary := models.FoodBeneficiary{
			User:  user[2],
			Stack: "Golang",
		}

		newUser, err := json.Marshal(beneficiary)
		if err != nil {
			t.Fail()
		}
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", "/api/v1/user/beneficiarysignup", strings.NewReader(string(newUser)))
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)
		assert.Contains(t, w.Body.String(), "enter valid decagon email")
	})

	t.Run("email exists", func(t *testing.T) {
		beneficiary := models.FoodBeneficiary{
			User:  user[1],
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
	})
}

//func TestFoodBeneficiarySignUpEmailExists(t *testing.T) {
//	ctrl := gomock.NewController(t)
//	mockDb := mocks.NewMockUserRepository(ctrl)
//	mockMail := mocks.NewMockMailerRepository(ctrl)
//
//	r := &api.HTTPHandler{
//		UserService:   mockDb,
//		MailerService: mockMail,
//	}
//
//	router := server.SetupRouter(r, mockDb)
//
//	user := models.User{
//		FullName:     "Orji Cecilia",
//		Email:        "cecilia.orji@decagon.dev",
//		Password:     "password",
//		PasswordHash: "",
//		Location:     "ETP",
//	}
//	beneficiary := models.FoodBeneficiary{
//		User:  user,
//		Stack: "Golang",
//	}
//
//	newUser, err := json.Marshal(beneficiary)
//	if err != nil {
//		t.Fail()
//	}
//	mockDb.EXPECT().FindFoodBenefactorByEmail(beneficiary.Email).Return(&beneficiary, nil)
//	w := httptest.NewRecorder()
//	req, _ := http.NewRequest("POST", "/api/v1/user/beneficiarysignup", strings.NewReader(string(newUser)))
//	router.ServeHTTP(w, req)
//
//	assert.Equal(t, http.StatusBadRequest, w.Code)
//	assert.Contains(t, w.Body.String(), "email exists")
//}
//
//func TestFoodBeneficiarySignUpBadRequest(t *testing.T) {
//	ctrl := gomock.NewController(t)
//	mockDb := mocks.NewMockUserRepository(ctrl)
//
//	defer ctrl.Finish()
//
//	r := &api.HTTPHandler{
//		UserService: mockDb,
//	}
//	router := server.SetupRouter(r, mockDb)
//	user := []models.User{
//		{
//			FullName:     "",
//			Email:        "cece",
//			Password:     "password",
//			PasswordHash: "",
//			Location:     "ETP",
//		},
//		{
//			FullName:     "Dede",
//			Email:        "cece@decagon.dev",
//			Password:     "password",
//			PasswordHash: "",
//			Location:     "ETP",
//		},
//	}
//
//	foodBeneficiary := models.FoodBeneficiary{
//		User:  user[0],
//		Stack: "java",
//	}
//
//	newUser, err := json.Marshal(foodBeneficiary)
//	if err != nil {
//		t.Fail()
//	}
//	beneficiary := models.FoodBeneficiary{
//		User:  user[1],
//		Stack: "gOLANG",
//	}
//
//	newUse, err := json.Marshal(beneficiary)
//	if err != nil {
//		t.Fail()
//	}
//
//	t.Run("Bad request", func(t *testing.T) {
//		w := httptest.NewRecorder()
//		req, _ := http.NewRequest("POST", "/api/v1/user/beneficiarysignup", strings.NewReader(string(newUser)))
//		router.ServeHTTP(w, req)
//
//		assert.Equal(t, http.StatusBadRequest, w.Code)
//		assert.Contains(t, w.Body.String(), "unable to bind request: validation error")
//	})
//
//	t.Run("Correct details", func(t *testing.T) {
//		mockDb.EXPECT().FindFoodBenefactorByEmail(beneficiary.Email).Return(&beneficiary, nil)
//		w := httptest.NewRecorder()
//		req, _ := http.NewRequest("POST", "/api/v1/user/beneficiarysignup", strings.NewReader(string(newUse)))
//		router.ServeHTTP(w, req)
//
//		assert.Equal(t, http.StatusBadRequest, w.Code)
//		assert.Contains(t, w.Body.String(), "email exist")
//	})
//
//}

//func TestSignupWithInCorrectDetailsTenant(t *testing.T) {
//	ctrl := gomock.NewController(t)
//	m := db.NewMockDB(ctrl)
//
//	s := &Server{
//		DB:     m,
//		Router: router.NewRouter(),
//	}
//	r := s.setupRouter()
//
//	role := models.Role{
//		Models: models.Models{},
//		Title:  "tenant",
//	}
//
//	user := models.User{
//		FirstName: "Spankie",
//		LastName:  "Dee",
//		Password:  "password",
//		Email:     "spankie_signup",
//		Phone1:    "08909876787",
//		RoleID:    role.ID,
//		Role:      role,
//	}
//	m.EXPECT().GetRoleByName("tenant").Return(role, nil)
//
//	jsonuser, err := json.Marshal(user)
//	if err != nil {
//		t.Fail()
//		return
//	}
//
//	w := httptest.NewRecorder()
//	req, _ := http.NewRequest("POST", "/api/v1/auth/signup_tenant", strings.NewReader(string(jsonuser)))
//	r.ServeHTTP(w, req)
//
//	bodyString := w.Body.String()
//	assert.Equal(t, http.StatusBadRequest, w.Code)
//	assert.Contains(t, bodyString, fmt.Sprintf("validation failed on field 'Email', condition: email, actual: %s", user.Email))
//}
//
//func TestSignupIfEmailExistsTenant(t *testing.T) {
//	ctrl := gomock.NewController(t)
//	m := db.NewMockDB(ctrl)
//
//	s := &Server{
//		DB:     m,
//		Router: router.NewRouter(),
//	}
//	r := s.setupRouter()
//
//	role := models.Role{
//		Models: models.Models{},
//		Title:  "tenant",
//	}
//
//	user := models.User{
//		FirstName: "Spankie",
//		LastName:  "Dee",
//		Password:  "password",
//		Address:   "1, boli drive",
//		Email:     "spankie_signup@gmail.com",
//		Phone1:    "08909876787",
//		RoleID:    role.ID,
//		Role:      role,
//	}
//
//	m.EXPECT().GetRoleByName("tenant").Return(role, nil)
//	m.EXPECT().FindUserByEmail(gomock.Any()).Return(&user, nil)
//
//	jsonuser, err := json.Marshal(user)
//	if err != nil {
//		t.Fail()
//		return
//	}
//	w := httptest.NewRecorder()
//	req, _ := http.NewRequest("POST", "/api/v1/auth/signup_tenant", strings.NewReader(string(jsonuser)))
//	r.ServeHTTP(w, req)
//
//	assert.Equal(t, http.StatusNotFound, w.Code)
//	assert.Contains(t, w.Body.String(), "user email already exists")
//}
//
//func TestSignupWithCorrectDetailsTenant(t *testing.T) {
//	ctrl := gomock.NewController(t)
//	m := db.NewMockDB(ctrl)
//
//	s := &Server{
//		DB:     m,
//		Router: router.NewRouter(),
//	}
//	r := s.setupRouter()
//
//	role := models.Role{
//		Models: models.Models{},
//		Title:  "tenant",
//	}
//
//	user := models.User{
//		FirstName: "Spankie",
//		LastName:  "Dee",
//		Password:  "password",
//		Address:   "1, boli drive",
//		Email:     "spankie_signup@gmail.com",
//		Phone1:    "08909876787",
//		RoleID:    role.ID,
//		Role:      role,
//	}
//
//	m.EXPECT().GetRoleByName("tenant").Return(role, nil)
//	m.EXPECT().FindUserByEmail(user.Email).Return(&user, nil)
//	jsonuser, err := json.Marshal(user)
//	if err != nil {
//		t.Fail()
//		return
//	}
//
//	w := httptest.NewRecorder()
//	req, _ := http.NewRequest("POST", "/api/v1/auth/signup_tenant", strings.NewReader(string(jsonuser)))
//	r.ServeHTTP(w, req)
//
//	m.EXPECT().GetRoleByName("tenant").Return(role, nil)
//	m.EXPECT().FindUserByEmail(user.Email).Return(&user, nil)
//	t.Run("check if tenant_email exists in the database", func(t *testing.T) {
//		w := httptest.NewRecorder()
//		req, _ := http.NewRequest("POST", "/api/v1/auth/signup_tenant", strings.NewReader(string(jsonuser)))
//		r.ServeHTTP(w, req)
//
//		assert.Equal(t, http.StatusNotFound, w.Code)
//		assert.Contains(t, w.Body.String(), "user email already exists")
//	})
//
//}
