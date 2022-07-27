package models

import (
	"net/mail"
	"regexp"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Model
	FullName     string `json:"full_name" binding:"required"`
	Email        string `json:"email" binding:"required" gorm:"unique"`
	Location     string `json:"location" binding:"required"`
	Password     string `json:"password,omitempty" gorm:"-"`
	PasswordHash string `json:"password_hash"`
	IsActive     bool   `json:"is_active"`
	Status       string `json:"status"`
	Avatar       string `json:"avatar"`
	Token        string `json:"token"`
}

//FoodBeneficiary represents a decadev
type FoodBeneficiary struct {
	User
	Stack string `json:"stack" binding:"required"`
}

type MealRecords struct {
	Model
	UserID    string `json:"user_id" gorm:"foreignKey"`
	UserEmail string `json:"user_email" gorm:"foreignKey"`
	Brunch    bool   `json:"brunch"`
	Dinner    bool   `json:"dinner"`
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

func (user *User) ValidateEmail() bool {
	emailRegexp := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
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
	if decagonEmail[1] == "decagonhq.com" {
		return true
	}
	return false
}
