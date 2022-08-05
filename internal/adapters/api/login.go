package api

import (
	"fmt"
	"github.com/decadevs/lunch-api/internal/core/helpers"
	"github.com/decadevs/lunch-api/internal/core/middleware"
	"github.com/decadevs/lunch-api/internal/core/models"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
	"os"
)

// LoginKitchenStaffHandler handles login for Kitchen Staff
func (u *HTTPHandler) LoginKitchenStaffHandler(c *gin.Context) {
	kitchenStaff := &models.KitchenStaff{}
	KitchenStaffLoginRequest := &struct {
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	}{}

	fmt.Println("here", kitchenStaff)
	err := c.ShouldBindJSON(KitchenStaffLoginRequest)
	if err != nil {
		helpers.JSON(c, "bad request", 400, nil, []string{"bad request"})
		return
	}

	kitchenStaff, sqlErr := u.UserService.FindKitchenStaffByEmail(KitchenStaffLoginRequest.Email)

	if sqlErr != nil {
		fmt.Println(sqlErr)
		helpers.JSON(c, "user not found, sign up", http.StatusInternalServerError, nil, []string{"internal server error"})
		return
	}

	if !kitchenStaff.IsActive {
		helpers.JSON(c, "please activate your account", http.StatusInternalServerError, nil, []string{"please activate your account"})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(kitchenStaff.PasswordHash), []byte(KitchenStaffLoginRequest.Password)); err != nil {
		helpers.JSON(c, "invalid Password", http.StatusInternalServerError, nil, []string{"interval server error"})
		return
	}

	// Generates access claims and refresh claims
	accessClaims, refreshClaims := middleware.GenerateClaims(kitchenStaff.Email)

	secret := os.Getenv("JWT_SECRET")
	accToken, err := middleware.GenerateToken(jwt.SigningMethodHS256, accessClaims, &secret)
	if err != nil {
		log.Printf("token generation error err: %v\n", err)
		helpers.JSON(c, "", http.StatusInternalServerError, nil, []string{"internal server error"})
		return
	}

	refreshToken, err := middleware.GenerateToken(jwt.SigningMethodHS256, refreshClaims, &secret)
	if err != nil {
		log.Printf("token generation error err: %v\n", err)
		helpers.JSON(c, "", http.StatusInternalServerError, nil, []string{"internal server error"})
		return
	}
	c.Header("refresh_token", *refreshToken)
	c.Header("access_token", *accToken)

	helpers.JSON(c, "login successful", http.StatusOK, gin.H{
		"user":          kitchenStaff,
		"access_token":  *accToken,
		"refresh_token": *refreshToken,
	}, nil)

}

// LoginFoodBenefactorHandler handles login for food benefactors eg. Decadevs
func (u *HTTPHandler) LoginFoodBenefactorHandler(c *gin.Context) {
	benefactor := &models.FoodBeneficiary{}
	benefactorLoginRequest := &struct {
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	}{}

	err := c.ShouldBindJSON(&benefactorLoginRequest)
	if err != nil {
		helpers.JSON(c, "bad request", 400, nil, []string{"bad request"})
		return
	}

	benefactor, sqlErr := u.UserService.FindFoodBenefactorByEmail(benefactorLoginRequest.Email)

	if sqlErr != nil {
		fmt.Println(sqlErr)
		helpers.JSON(c, "email exists", http.StatusInternalServerError, nil, []string{"internal server error"})
		return
	}

	if !benefactor.IsActive {
		helpers.JSON(c, "please activate your account", http.StatusInternalServerError, nil, []string{"please activate your account"})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(benefactor.PasswordHash), []byte(benefactorLoginRequest.Password)); err != nil {
		helpers.JSON(c, "invalid Password", http.StatusInternalServerError, nil, []string{"interval server error"})
		return
	}

	// Generates access claims and refresh claims
	accessClaims, refreshClaims := middleware.GenerateClaims(benefactor.Email)

	secret := os.Getenv("JWT_SECRET")
	accToken, err := middleware.GenerateToken(jwt.SigningMethodHS256, accessClaims, &secret)
	if err != nil {
		log.Printf("token generation error err: %v\n", err)
		helpers.JSON(c, "", http.StatusInternalServerError, nil, []string{"internal server error"})
		return
	}

	refreshToken, err := middleware.GenerateToken(jwt.SigningMethodHS256, refreshClaims, &secret)
	if err != nil {
		log.Printf("token generation error err: %v\n", err)
		helpers.JSON(c, "", http.StatusInternalServerError, nil, []string{"internal server error"})
		return
	}
	c.Header("refresh_token", *refreshToken)
	c.Header("access_token", *accToken)

	helpers.JSON(c, "login successful", http.StatusOK, gin.H{
		"user":          benefactor,
		"access_token":  *accToken,
		"refresh_token": *refreshToken,
	}, nil)

}

// LoginAdminHandler handles login for the Admin
func (u *HTTPHandler) LoginAdminHandler(c *gin.Context) {
	admin := &models.Admin{}
	adminLoginRequest := &struct {
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	}{}

	err := c.ShouldBindJSON(&adminLoginRequest)
	if err != nil {
		helpers.JSON(c, "bad request", 400, nil, []string{"bad request"})
		return
	}

	admin, sqlErr := u.UserService.FindAdminByEmail(adminLoginRequest.Email)

	if sqlErr != nil {
		fmt.Println(sqlErr)
		helpers.JSON(c, "email exists", http.StatusInternalServerError, nil, []string{"internal server error"})
		return
	}

	if !admin.IsActive {
		helpers.JSON(c, "please activate your account", http.StatusInternalServerError, nil, []string{"please activate your account"})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(admin.PasswordHash), []byte(adminLoginRequest.Password)); err != nil {
		helpers.JSON(c, "invalid Password", http.StatusInternalServerError, nil, []string{"interval server error"})
		return
	}

	// Generates access claims and refresh claims
	accessClaims, refreshClaims := middleware.GenerateClaims(admin.Email)

	secret := os.Getenv("JWT_SECRET")
	accToken, err := middleware.GenerateToken(jwt.SigningMethodHS256, accessClaims, &secret)
	if err != nil {
		log.Printf("token generation error err: %v\n", err)
		helpers.JSON(c, "", http.StatusInternalServerError, nil, []string{"internal server error"})
		return
	}

	refreshToken, err := middleware.GenerateToken(jwt.SigningMethodHS256, refreshClaims, &secret)
	if err != nil {
		log.Printf("token generation error err: %v\n", err)
		helpers.JSON(c, "", http.StatusInternalServerError, nil, []string{"internal server error"})
		return
	}
	c.Header("refresh_token", *refreshToken)
	c.Header("access_token", *accToken)

	helpers.JSON(c, "login successful", http.StatusOK, gin.H{
		"user":          admin,
		"access_token":  *accToken,
		"refresh_token": *refreshToken,
	}, nil)

}
