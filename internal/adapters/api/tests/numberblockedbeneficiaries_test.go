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

func TestGetNumberOfBlockedUsers(t *testing.T) {
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
	var correctNum int64 = 1

	secret := os.Getenv("JWT_SECRET")
	accessClaims, _ := middleware.GenerateClaims(admin.Email)
	accToken, _ := middleware.GenerateToken(jwt.SigningMethodHS256, accessClaims, &secret)

	t.Run("testing bad request", func(t *testing.T) {
		mockDb.EXPECT().TokenInBlacklist(gomock.Any()).Return(false)
		mockDb.EXPECT().FindAdminByEmail(admin.Email).Return(&admin, nil)
		mockDb.EXPECT().NumberOfBlockedBeneficiary().Return(num, errors.New("internal server error"))
		rw := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodGet, "/api/v1/admin/numberblocked", strings.NewReader(string(bytes)))
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", *accToken))
		router.ServeHTTP(rw, req)
		assert.Equal(t, http.StatusInternalServerError, rw.Code)
		assert.Contains(t, rw.Body.String(), "internal server error")
	})

	t.Run("successful request", func(t *testing.T) {
		mockDb.EXPECT().TokenInBlacklist(gomock.Any()).Return(false)
		mockDb.EXPECT().FindAdminByEmail(admin.Email).Return(&admin, nil)
		mockDb.EXPECT().NumberOfBlockedBeneficiary().Return(correctNum, nil)
		rw := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodGet, "/api/v1/admin/numberblocked", strings.NewReader(string(bytes)))
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", *accToken))
		router.ServeHTTP(rw, req)
		assert.Equal(t, http.StatusOK, rw.Code)
		assert.Contains(t, rw.Body.String(), "successfully gotten")
	})
}
