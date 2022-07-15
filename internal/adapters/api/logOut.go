package api

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/decadevs/lunch-api/internal/core/helpers"
	"github.com/decadevs/lunch-api/internal/core/middleware"
	"github.com/decadevs/lunch-api/internal/core/models"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func (u HTTPHandler) FoodBeneficiaryLogout(c *gin.Context) {
	//create token blacklist struct
	tokenI, exists := c.Get("access_token")
	if !exists {
		helpers.JSON(c, "error getting access token", http.StatusBadRequest, nil, []string{"error getting access token"})
	}

	foodBeneficiary, exists := c.Get("user")
	if !exists {
		helpers.JSON(c, "error getting user from context",
			http.StatusBadRequest, nil, []string{"error getting user from context"})
	}
	beneficiary := foodBeneficiary.(*models.FoodBeneficiary)
	tokenStr := tokenI.(string)

	token, err := jwt.ParseWithClaims(tokenStr, &middleware.Claims{}, func(t *jwt.Token) (interface{}, error) {

		if t.Method.Alg() != jwt.SigningMethodHS256.Alg() {
			return nil, fmt.Errorf("invalid signing algorithm")
		}
		return os.Getenv("JWT_SECRET"), nil
	})
	if claims, ok := token.Claims.(*middleware.Claims); !ok && !token.Valid {
		helpers.JSON(c, "error inserting claims",
			http.StatusBadRequest, nil, []string{"Claims not valid type"})

	} else {
		claims.StandardClaims.ExpiresAt = time.Now().Add(-time.Hour).Unix()
	}

	err = u.UserService.AddTokenToBlacklist(beneficiary.Email, tokenStr)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error inserting Token into database", "Error": err})

	}

	helpers.JSON(c, "signed out successfully", 200, nil, nil)

}

func (u HTTPHandler) KitchenStaffLogout(c *gin.Context) {
	//create token blacklist struct
	tokenI, exists := c.Get("access_token")
	if !exists {
		helpers.JSON(c, "error getting access token", http.StatusBadRequest, nil, []string{"error getting access token"})
	}

	kitchenStaff, exists := c.Get("user")
	if !exists {
		helpers.JSON(c, "error getting user from context",
			http.StatusBadRequest, nil, []string{"error getting user from context"})
	}
	staff := kitchenStaff.(*models.KitchenStaff)
	tokenStr := tokenI.(string)

	token, err := jwt.ParseWithClaims(tokenStr, &middleware.Claims{}, func(t *jwt.Token) (interface{}, error) {

		if t.Method.Alg() != jwt.SigningMethodHS256.Alg() {
			return nil, fmt.Errorf("invalid signing algorithm")
		}
		return os.Getenv("JWT_SECRET"), nil
	})
	if claims, ok := token.Claims.(*middleware.Claims); !ok && !token.Valid {
		helpers.JSON(c, "error inserting claims",
			http.StatusBadRequest, nil, []string{"Claims not valid type"})

	} else {
		claims.StandardClaims.ExpiresAt = time.Now().Add(-time.Hour).Unix()
	}

	err = u.UserService.AddTokenToBlacklist(staff.Email, tokenStr)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error inserting Token into database", "Error": err})

	}

	helpers.JSON(c, "signed out successfully", 200, nil, nil)

}
