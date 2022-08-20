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

func TestUpdateMeal(t *testing.T) {
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
	year, month, day := time.Now().Date()
	food := models.Food{
		Name:    "Afang Soup",
		Type:    "brunch",
		Year:    year,
		Month:   int(month),
		Day:     day,
		Weekday: "Wednesday",
	}

	foodJSON, err := json.Marshal(food)
	if err != nil {
		t.Fail()
	}

	notification := models.Notification{
		Message: admin.FullName + " updated timetable",
		Year:    food.Year,
		Month:   int(month),
		Day:     food.Day,
	}

	secret := os.Getenv("JWT_SECRET")
	accessClaims, _ := middleware.GenerateClaims(admin.Email)
	accToken, _ := middleware.GenerateToken(jwt.SigningMethodHS256, accessClaims, &secret)

	t.Run("Successful Request", func(t *testing.T) {
		mockDb.EXPECT().TokenInBlacklist(gomock.Any()).Return(false)
		mockDb.EXPECT().FindAdminByEmail(admin.Email).Return(&admin, nil)
		mockDb.EXPECT().UpdateMeal("4", food).Return(nil)
		mockDb.EXPECT().CreateNotification(notification).Return(nil)
		rw := httptest.NewRecorder()
		req, err := http.NewRequest(http.MethodPut,
			"/api/v1/admin/updatemeal/4",
			strings.NewReader(string(foodJSON)))
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
		req, err := http.NewRequest(http.MethodPut,
			"/api/v1/admin/updatemeal/4",
			strings.NewReader(string(foodJSON)))
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
