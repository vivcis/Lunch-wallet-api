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
	"time"
)

func TestDeleteMeal(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockDb := mocks.NewMockUserRepository(ctrl)

	r := &api.HTTPHandler{
		UserService: mockDb,
	}

	router := server.SetupRouter(r, mockDb)

	user := models.User{
		FullName: "Joseph Asuquo",
		Email:    "joseph@yahoo.com",
		Location: "Edo Tech Park",
		IsActive: true,
	}

	admin := models.Admin{User: user}

	id := "4"
	idJSON, err := json.Marshal(id)
	if err != nil {
		t.Fail()
	}

	year, month, day := time.Now().Date()

	notification := models.Notification{
		Message: user.FullName + " updated timetable",
		Year:    year,
		Month:   time.Month(month),
		Day:     day,
	}

	secret := os.Getenv("JWT_SECRET")
	accessClaims, _ := middleware.GenerateClaims(admin.Email)
	accToken, _ := middleware.GenerateToken(jwt.SigningMethodHS256, accessClaims, &secret)

	t.Run("Successful Request", func(t *testing.T) {
		mockDb.EXPECT().TokenInBlacklist(gomock.Any()).Return(false)
		mockDb.EXPECT().FindAdminByEmail(admin.Email).Return(&admin, nil)
		mockDb.EXPECT().DeleteMeal(id).Return(nil)
		mockDb.EXPECT().CreateNotification(notification).Return(nil)
		rw := httptest.NewRecorder()
		req, err := http.NewRequest(http.MethodDelete,
			"/api/v1/admin/deletemeal/4",
			strings.NewReader(string(idJSON)))
		if err != nil {
			fmt.Printf("errrr here %v \n", err)
			return
		}
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", *accToken))

		router.ServeHTTP(rw, req)
		assert.Equal(t, http.StatusOK, rw.Code)
		assert.Contains(t, rw.Body.String(), "Successful")
	})

	t.Run("testing error in context", func(t *testing.T) {
		mockDb.EXPECT().TokenInBlacklist(gomock.Any()).Return(false)
		mockDb.EXPECT().FindAdminByEmail(admin.Email).Return(nil, errors.New("user not found"))

		rw := httptest.NewRecorder()
		req, err := http.NewRequest(http.MethodDelete,
			"/api/v1/admin/deletemeal/4",
			strings.NewReader(string(idJSON)))
		if err != nil {
			fmt.Printf("errrr here %v \n", err)
			return
		}
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", *accToken))
		router.ServeHTTP(rw, req)
		assert.Equal(t, http.StatusNotFound, rw.Code)
		assert.Contains(t, rw.Body.String(), "user not found")
	})

}
