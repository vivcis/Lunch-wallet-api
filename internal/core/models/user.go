package models

import "net/mail"

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
	Stack string `json:"stack"`
}
type KitchenStaff struct {
	User
}
type Admin struct {
	User
}

func (user *User) ValidMailAddress() bool {
	_, err := mail.ParseAddress(user.Email)
	if err != nil {
		return false
	}
	return true
}