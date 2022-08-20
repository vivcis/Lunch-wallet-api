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

func TestGetFoodBeneficiaryProfile(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockDb := mocks.NewMockUserRepository(ctrl)

	r := &api.HTTPHandler{
		UserService: mockDb,
	}
	router := server.SetupRouter(r, mockDb)

	user := models.User{
		FullName: "Orji Cecilia",
		Email:    "cecilia.orji@decagonhq.com",
		Location: "Edo Tech Park",
		Avatar:   "img.png",
	}
	foodBeneficiary := models.FoodBeneficiary{
		User:  user,
		Stack: "golang",
	}

	admin := models.Admin{User: user}

	id := "1"
	bytes, _ := json.Marshal(id)

	secret := os.Getenv("JWT_SECRET")
	accessClaims, _ := middleware.GenerateClaims(admin.Email)
	accToken, _ := middleware.GenerateToken(jwt.SigningMethodHS256, accessClaims, &secret)

	t.Run("test bad request", func(t *testing.T) {
		mockDb.EXPECT().TokenInBlacklist(gomock.Any()).Return(false)
		mockDb.EXPECT().FindAdminByEmail(admin.Email).Return(&admin, nil)
		mockDb.EXPECT().GetFoodBenefactorById(gomock.Any()).Return(nil, errors.New("record not found"))
		rw := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodGet, "/api/v1/admin/getfoodbeneficiaryprofile/1", strings.NewReader(string(bytes)))
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", *accToken))
		router.ServeHTTP(rw, req)
		assert.Equal(t, http.StatusInternalServerError, rw.Code)
		assert.Contains(t, rw.Body.String(), "internal server error")
	})

	t.Run("Successful Profile Gotten", func(t *testing.T) {
		mockDb.EXPECT().TokenInBlacklist(gomock.Any()).Return(false)
		mockDb.EXPECT().FindAdminByEmail(admin.Email).Return(&admin, nil)
		mockDb.EXPECT().GetFoodBenefactorById(id).Return(&foodBeneficiary, nil)
		rw := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodGet,
			"/api/v1/admin/getfoodbeneficiaryprofile/1",
			strings.NewReader(string(bytes)))
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", *accToken))

		router.ServeHTTP(rw, req)
		assert.Equal(t, http.StatusOK, rw.Code)
		assert.Contains(t, rw.Body.String(), "food beneficiary details retrieved")
	})
}
