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
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
	"time"
)

func TestGetMeal(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockDb := mocks.NewMockUserRepository(ctrl)

	r := &api.HTTPHandler{
		UserService: mockDb,
	}
	router := server.SetupRouter(r, mockDb)

	mealType := "brunch"
	meals := &[]models.Food{
		{
			Model:     models.Model{},
			Name:      "eba",
			Type:      mealType,
			AdminName: "bolu",
			Kitchen:   "etp",
		}, {
			Model:     models.Model{},
			Name:      "dodo",
			Type:      "dinner",
			AdminName: "bolu",
			Kitchen:   "etp",
		},
	}
	user := models.User{
		Model:    models.Model{},
		FullName: "Michael Gbenle",
		Email:    "michael.gbenle@decagon.dev",
		Location: "Edo Tech Park",
		IsActive: true,
	}
	kitchenStaff := &models.KitchenStaff{
		User: user,
	}
	kitchenStaffEmail := "bolu@decagon.dev"
	year, month, day := time.Now().Date()
	secret := os.Getenv("JWT_SECRET")
	accessClaims, _ := middleware.GenerateClaims(kitchenStaffEmail)
	accToken, _ := middleware.GenerateToken(jwt.SigningMethodHS256, accessClaims, &secret)

	t.Run("testing bad request", func(t *testing.T) {
		mockDb.EXPECT().TokenInBlacklist(gomock.Any()).Return(false)
		mockDb.EXPECT().FindFoodBenefactorByEmail(kitchenStaff.Email).Return(&kitchenStaff, nil)
		mockDb.EXPECT().FindAllFoodByDate(year, int(month), day).Return(&meals, nil)

		bytes, _ := json.Marshal(meals)
		rw := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodGet, "/api/v1/staff/generateqrcode?mealType=brunch", strings.NewReader(string(bytes)))
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", *accToken))
		router.ServeHTTP(rw, req)
		assert.Equal(t, http.StatusBadRequest, rw.Code)
		assert.Contains(t, rw.Body.String(), "internal server error")

	})
}
