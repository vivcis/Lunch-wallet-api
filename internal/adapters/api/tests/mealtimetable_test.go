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
	"time"
)

func TestGetMealTimeTableHandle(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockDb := mocks.NewMockUserRepository(ctrl)

	r := &api.HTTPHandler{
		UserService: mockDb,
	}
	router := server.SetupRouter(r, mockDb)

	user := models.User{
		FullName: "sabi Nwa",
		Email:    "sabinwa@decagon.dev",
		IsActive: true,
	}

	benefactor := models.Admin{
		User: user,
	}
	year, month, day := time.Now().Date()

	food := models.Food{
		Name:      "Egusi and choice of swallow",
		Type:      "BRUNCH",
		AdminName: "Investor Sabinus",
		Year:      year,
		Month:     int(month),
		Day:       day,
		Weekday:   "Friday",
		Status:    "Not serve",
	}

	foods := []models.Food{
		food,
	}

	bytes, _ := json.Marshal(food)

	secret := os.Getenv("JWT_SECRET")
	accessClaims, _ := middleware.GenerateClaims(benefactor.Email)
	accToken, _ := middleware.GenerateToken(jwt.SigningMethodHS256, accessClaims, &secret)

	t.Run("testing bad request", func(t *testing.T) {
		mockDb.EXPECT().TokenInBlacklist(gomock.Any()).Return(false)
		mockDb.EXPECT().FindAdminByEmail(benefactor.Email).Return(&benefactor, nil)
		mockDb.EXPECT().FindFoodByDate(year, food.Month, day).Return(nil, errors.New("internal server error"))
		rw := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodGet, "/api/v1/admin/getTimetable", strings.NewReader(string(bytes)))
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", *accToken))
		router.ServeHTTP(rw, req)
		assert.Equal(t, http.StatusInternalServerError, rw.Code)
		assert.Contains(t, rw.Body.String(), "internal server error")
	})

	t.Run("testing Successful request", func(t *testing.T) {
		mockDb.EXPECT().TokenInBlacklist(gomock.Any()).Return(false)
		mockDb.EXPECT().FindAdminByEmail(benefactor.Email).Return(&benefactor, nil)
		mockDb.EXPECT().FindFoodByDate(year, food.Month, day).Return(foods, nil)
		rw := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodGet, "/api/v1/admin/getTimetable", strings.NewReader(string(bytes)))
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", *accToken))
		router.ServeHTTP(rw, req)
		assert.Equal(t, http.StatusOK, rw.Code)
		assert.Contains(t, rw.Body.String(), "Timetable found")
	})
}

func TestAdminTotalNumberOfUsers(t *testing.T) {
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

	admin := models.Admin{
		User: user1,
	}
	bytes, _ := json.Marshal(user1)

	secret := os.Getenv("JWT_SECRET")
	accessClaims, _ := middleware.GenerateClaims(admin.Email)
	accToken, _ := middleware.GenerateToken(jwt.SigningMethodHS256, accessClaims, &secret)

	t.Run("testing bad request", func(t *testing.T) {
		mockDb.EXPECT().TokenInBlacklist(gomock.Any()).Return(false)
		mockDb.EXPECT().FindAdminByEmail(admin.Email).Return(&admin, nil)
		mockDb.EXPECT().GetTotalUsers().Return(count, errors.New("internal server error"))
		rw := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodGet, "/api/v1/admin/getTotalNumberOfUsers", strings.NewReader(string(bytes)))
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", *accToken))
		router.ServeHTTP(rw, req)
		assert.Equal(t, http.StatusInternalServerError, rw.Code)
		assert.Contains(t, rw.Body.String(), "internal server error")
	})

	t.Run("testing success", func(t *testing.T) {
		mockDb.EXPECT().TokenInBlacklist(gomock.Any()).Return(false)
		mockDb.EXPECT().FindAdminByEmail(admin.Email).Return(&admin, nil)
		mockDb.EXPECT().GetTotalUsers().Return(count, nil)
		rw := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodGet, "/api/v1/admin/getTotalNumberOfUsers", strings.NewReader(string(bytes)))
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", *accToken))
		router.ServeHTTP(rw, req)
		assert.Equal(t, http.StatusOK, rw.Code)
		assert.Contains(t, rw.Body.String(), "Total number")
	})
}

func TestAdminSearchBeneficiaries(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockDb := mocks.NewMockUserRepository(ctrl)

	r := &api.HTTPHandler{
		UserService: mockDb,
	}

	router := server.SetupRouter(r, mockDb)

	user := models.User{
		Model:    models.Model{},
		FullName: "investor sabinus",
		Email:    "sabinwa@decagon.dev",
		Location: "ETP",
		Password: "problem",
		IsActive: false,
		Status:   "active",
		Avatar:   "image.png",
	}

	admin := models.Admin{
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
	accessClaims, _ := middleware.GenerateClaims(admin.Email)
	accToken, _ := middleware.GenerateToken(jwt.SigningMethodHS256, accessClaims, &secret)

	t.Run("test bad request", func(t *testing.T) {
		mockDb.EXPECT().TokenInBlacklist(gomock.Any()).Return(false)
		mockDb.EXPECT().FindAdminByEmail(admin.Email).Return(&admin, nil)
		mockDb.EXPECT().SearchFoodBeneficiary(gomock.Any(), gomock.Any()).Return(nil, errors.New("record not found"))
		rw := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodGet, "/api/v1/admin/searchBeneficiaries/python", strings.NewReader(string(bytes)))
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", *accToken))
		router.ServeHTTP(rw, req)
		assert.Equal(t, http.StatusInternalServerError, rw.Code)
		assert.Contains(t, rw.Body.String(), "internal server error")
	})

	t.Run("test successful search", func(t *testing.T) {
		mockDb.EXPECT().TokenInBlacklist(gomock.Any()).Return(false)
		mockDb.EXPECT().FindAdminByEmail(admin.Email).Return(&admin, nil)
		mockDb.EXPECT().SearchFoodBeneficiary(beneficiary.Stack, &page).Return(foodBeneficiary, nil)
		rw := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodGet, "/api/v1/admin/searchBeneficiaries/Golang", strings.NewReader(string(bytes)))
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", *accToken))
		router.ServeHTTP(rw, req)
		assert.Equal(t, http.StatusOK, rw.Code)
		assert.Contains(t, rw.Body.String(), "information gotten")
	})
}

func TestGeAllBeneficiariesHandle(t *testing.T) {
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

	admin := models.Admin{
		User: user,
	}
	page := models.Pagination{
		Page:  1,
		Limit: 10,
		Sort:  "created_at asc",
	}

	var foodBeneficiary []models.UserDetails

	bytes, _ := json.Marshal(user)

	secret := os.Getenv("JWT_SECRET")
	accessClaims, _ := middleware.GenerateClaims(admin.Email)
	accToken, _ := middleware.GenerateToken(jwt.SigningMethodHS256, accessClaims, &secret)

	t.Run("testing successful get users", func(t *testing.T) {
		mockDb.EXPECT().TokenInBlacklist(gomock.Any()).Return(false)
		mockDb.EXPECT().FindAdminByEmail(admin.Email).Return(&admin, nil)
		mockDb.EXPECT().FindAllFoodBeneficiary(&page).Return(foodBeneficiary, nil)
		rw := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodGet, "/api/v1/admin/getAllBeneficiaries", strings.NewReader(string(bytes)))
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", *accToken))
		router.ServeHTTP(rw, req)
		assert.Equal(t, http.StatusOK, rw.Code)
		assert.Contains(t, rw.Body.String(), "food beneficiaries found successfully")
	})
}

func TestGetNumberOfScannedUsers(t *testing.T) {
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

	beneficiary := models.FoodBeneficiary{
		User: models.User{
			FullName: "Tony",
			Email:    "tony@decagon.dev",
			Location: "Edo Tech Park",
			IsActive: true,
			IsBlock:  true,
		},
		Stack: "node",
	}

	beneficiaries := []models.FoodBeneficiary{
		beneficiary,
	}

	bytes, _ := json.Marshal(beneficiaries)

	var num int64 = 0
	var correctNum = 1
	date := time.Now().Format("2006-01-02")

	secret := os.Getenv("JWT_SECRET")
	accessClaims, _ := middleware.GenerateClaims(admin.Email)
	accToken, _ := middleware.GenerateToken(jwt.SigningMethodHS256, accessClaims, &secret)

	t.Run("testing bad request", func(t *testing.T) {
		mockDb.EXPECT().TokenInBlacklist(gomock.Any()).Return(false)
		mockDb.EXPECT().FindAdminByEmail(admin.Email).Return(&admin, nil)
		mockDb.EXPECT().FindNumbersOfScannedUsers(date).Return(num, errors.New("internal server error"))
		//mockDb.EXPECT().NumberOfBlockedBeneficiary().Return(num, errors.New("internal server error"))
		rw := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodGet, "/api/v1/admin/getTotalNumberOfScannedUsers", strings.NewReader(string(bytes)))
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", *accToken))
		router.ServeHTTP(rw, req)
		assert.Equal(t, http.StatusInternalServerError, rw.Code)
		assert.Contains(t, rw.Body.String(), "internal server error")
	})

	t.Run("successful request", func(t *testing.T) {
		mockDb.EXPECT().TokenInBlacklist(gomock.Any()).Return(false)
		mockDb.EXPECT().FindAdminByEmail(admin.Email).Return(&admin, nil)
		mockDb.EXPECT().FindNumbersOfScannedUsers(date).Return(num, nil)
		mockDb.EXPECT().GetTotalUsers().Return(correctNum, nil)
		rw := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodGet, "/api/v1/admin/getTotalNumberOfScannedUsers", strings.NewReader(string(bytes)))
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", *accToken))
		router.ServeHTTP(rw, req)
		assert.Equal(t, http.StatusOK, rw.Code)
		assert.Contains(t, rw.Body.String(), "Successful")
	})
}
