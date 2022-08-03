package api

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/decadevs/lunch-api/internal/core/helpers"
	"github.com/decadevs/lunch-api/internal/core/middleware"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func (u *HTTPHandler) FoodBeneficiaryLogout(c *gin.Context) {
	tokenstr, err := u.GetTokenFromContext(c)
	if err != nil {
		helpers.JSON(c, "error getting access token", http.StatusBadRequest, nil, []string{"bad request"})
		return
	}

	foodBeneficiary, err := u.GetBenefactorFromContext(c)
	if err != nil {
		helpers.JSON(c, "error getting access token", http.StatusBadRequest, nil, []string{"bad request"})
		return
	}

	token, err := jwt.ParseWithClaims(tokenstr, &middleware.Claims{}, func(t *jwt.Token) (interface{}, error) {

		if t.Method.Alg() != jwt.SigningMethodHS256.Alg() {
			return nil, fmt.Errorf("invalid signing algorithm")
		}
		return os.Getenv("JWT_SECRET"), nil
	})
	if claims, ok := token.Claims.(*middleware.Claims); !ok && !token.Valid {
		helpers.JSON(c, "error inserting claims", http.StatusBadRequest, nil, []string{"bad request"})
		return

	} else {
		claims.StandardClaims.ExpiresAt = time.Now().Add(-time.Hour).Unix()
	}

	err = u.UserService.AddTokenToBlacklist(foodBeneficiary.Email, tokenstr)
	if err != nil {
		helpers.JSON(c, "error inserting token into database", http.StatusInternalServerError, nil, []string{"Claims not valid type"})
		return
	}

	helpers.JSON(c, "signed out successfully", 200, nil, nil)

}

func (u *HTTPHandler) KitchenStaffLogout(c *gin.Context) {
	tokenstr, err := u.GetTokenFromContext(c)
	if err != nil {
		helpers.JSON(c, "error getting access token", http.StatusBadRequest, nil, []string{"bad request"})
		return
	}

	kitchenStaff, err := u.GetKitchenStaffFromContext(c)
	if err != nil {
		helpers.JSON(c, "error getting access token", http.StatusBadRequest, nil, []string{"bad request"})
		return
	}
	token, err := jwt.ParseWithClaims(tokenstr, &middleware.Claims{}, func(t *jwt.Token) (interface{}, error) {

		if t.Method.Alg() != jwt.SigningMethodHS256.Alg() {
			return nil, fmt.Errorf("invalid signing algorithm")
		}
		return os.Getenv("JWT_SECRET"), nil
	})
	if claims, ok := token.Claims.(*middleware.Claims); !ok && !token.Valid {
		helpers.JSON(c, "error inserting claims", http.StatusBadRequest, nil, []string{"Claims not valid type"})
		return

	} else {
		claims.StandardClaims.ExpiresAt = time.Now().Add(-time.Hour).Unix()
	}

	err = u.UserService.AddTokenToBlacklist(kitchenStaff.Email, tokenstr)
	if err != nil {
		helpers.JSON(c, "error inserting token into database", http.StatusInternalServerError, nil, []string{"Claims not valid type"})
		return
	}

	helpers.JSON(c, "signed out successfully", 200, nil, nil)

}
