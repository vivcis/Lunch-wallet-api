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

func TestUpdateBrunchFoodStatus(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockDb := mocks.NewMockUserRepository(ctrl)

	r := &api.HTTPHandler{
		UserService: mockDb,
	}

	router := server.SetupRouter(r, mockDb)

	user := models.User{}

	year, mon, day := time.Now().Date()
	month := int(mon)

	kitchenStaff := models.KitchenStaff{
		User: user,
	}

	createFood := models.Food{
		Name:      "Fried Rice",
		Type:      "BRUNCH",
		AdminName: "Orji Cecilia",
		Year:      year,
		Month:     month,
		Day:       day,
		Weekday:   "Sunday",
		Status:    "SERVED",
	}
	foods := []models.Food{
		createFood,
	}
	bytes, _ := json.Marshal(createFood)

	secret := os.Getenv("JWT_SECRET")
	accessClaims, _ := middleware.GenerateClaims(kitchenStaff.Email)
	accToken, _ := middleware.GenerateToken(jwt.SigningMethodHS256, accessClaims, &secret)

	notification := models.Notification{
		Message: createFood.Status,
		Year:    year,
		Month:   int(month),
		Day:     day,
	}

	t.Run("testing bad request", func(t *testing.T) {
		mockDb.EXPECT().TokenInBlacklist(gomock.Any()).Return(false)
		mockDb.EXPECT().FindKitchenStaffByEmail(kitchenStaff.Email).Return(&kitchenStaff, nil)
		mockDb.EXPECT().FindBrunchByDate(year, month, day).Return(nil, errors.New("internal server error"))
		rw := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodPut, "/api/v1/staff/changebrunchstatus", strings.NewReader(string(bytes)))
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", *accToken))
		router.ServeHTTP(rw, req)
		assert.Equal(t, http.StatusInternalServerError, rw.Code)
		assert.Contains(t, rw.Body.String(), "internal server error")
	})

	t.Run("testing Successful request", func(t *testing.T) {
		mockDb.EXPECT().TokenInBlacklist(gomock.Any()).Return(false)
		mockDb.EXPECT().FindKitchenStaffByEmail(kitchenStaff.Email).Return(&kitchenStaff, nil)
		mockDb.EXPECT().FindBrunchByDate(year, createFood.Month, day).Return(foods, nil)
		mockDb.EXPECT().UpdateStatus(foods, createFood.Status).Return(nil)
		mockDb.EXPECT().CreateNotification(notification).Return(nil)
		mockDb.EXPECT().FindBrunchByDate(year, month, day).Return(foods, nil)
		rw := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodPut, "/api/v1/staff/changebrunchstatus", strings.NewReader(string(bytes)))
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", *accToken))
		router.ServeHTTP(rw, req)
		assert.Equal(t, http.StatusOK, rw.Code)
		assert.Contains(t, rw.Body.String(), "food status updated")
	})
}
