package models

import (
	"golang.org/x/crypto/bcrypt"
	"net/mail"
	"regexp"
	"strings"
	"unicode"
)

type User struct {
	Model
	FullName     string `json:"full_name" binding:"required"`
	Email        string `json:"email" binding:"required,email" gorm:"unique"`
	Location     string `json:"location" binding:"required"`
	Password     string `json:"password,omitempty" gorm:"-"`
	PasswordHash string `json:"password_hash"`
	IsActive     bool   `json:"is_active"`
	Status       string `json:"status"`
	Avatar       string `json:"avatar"`
	Token        string `json:"token"`
	IsBlock      bool   `json:"is_block"`
}

type UserDetails struct {
	FullName string `json:"full_name" binding:"required"`
	Stack    string `json:"stack"`
	Location string `json:"location"`
}

type UserProfile struct {
	FullName string `json:"full_name" binding:"required"`
	Stack    string `json:"stack"`
	Email    string `json:"email"`
	Location string `json:"location"`
	Avatar   string `json:"avatar"`
}

//FoodBeneficiary represents a decadev
type FoodBeneficiary struct {
	User
	Stack string `json:"stack" binding:"required"`
}

type MealRecords struct {
	Model
	MealDate  string `json:"meal_date"`
	UserID    string `json:"user_id" gorm:"foreignKey"`
	UserEmail string `json:"user_email" gorm:"foreignKey"`
	Brunch    bool   `json:"brunch"`
	Dinner    bool   `json:"dinner"`
}

type QRCodeMealRecords struct {
	Model
	MealId string `json:"meal_id" binding:"required"`
	UserId string `json:"user_id" binding:"required"`
}

type KitchenStaff struct {
	User
}
type Admin struct {
	User
}

type ForgotPassword struct {
	Email string `json:"email"`
}

type ResetPassword struct {
	NewPassword        string `json:"new_password"`
	ConfirmNewPassword string `json:"confirm_new_password"`
}

type Pagination struct {
	Limit int    `json:"limit"`
	Page  int    `json:"page"`
	Sort  string `json:"sort"`
}

func (user *User) ValidMailAddress() bool {
	_, err := mail.ParseAddress(user.Email)
	if err != nil {
		return false
	}
	return true
}

func (user *User) HashPassword() error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.PasswordHash = string(hashedPassword)
	return nil
}

func (user *User) PasswordStrength() bool {
	password := regexp.MustCompile(`^(?=.*[a-z])(?=.*[A-Z])(?=.*\d)(?=.*[@$!%*?&])[A-Za-z\d@$!%*?&]{8}$`)
	return password.MatchString(user.Password)
}

func (user *User) ValidateEmail() bool {
	emailRegexp := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{5,50}$`)
	return emailRegexp.MatchString(user.Email)

}

func (user *User) ValidateDecagonEmail() bool {
	decagon := strings.Split(user.Email, "@")
	if decagon[1] == "decagon.dev" {
		return true
	}
	return false
}

func (user *User) ValidAdminDecagonEmail() bool {
	decagonEmail := strings.Split(user.Email, "@")
	if decagonEmail[1] == "decagon.dev" {
		return true
	}
	return false
}

func (user *User) IsValid(password string) bool {
	var (
		hasMinLen  = false
		hasUpper   = false
		hasLower   = false
		hasNumber  = false
		hasSpecial = false
	)
	if len(password) >= 8 {
		hasMinLen = true
	}
	for _, char := range password {
		switch {
		case unicode.IsUpper(char):
			hasUpper = true
		case unicode.IsLower(char):
			hasLower = true
		case unicode.IsNumber(char):
			hasNumber = true
		case unicode.IsPunct(char) || unicode.IsSymbol(char):
			hasSpecial = true
		}
	}
	return hasMinLen && hasUpper && hasLower && hasNumber && hasSpecial
}
