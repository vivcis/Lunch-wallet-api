package models

type User struct {
	Model
	FullName     string `json:"full_name"`
	Email        string `json:"email" gorm:"unique"`
	Location     string `json:"location"`
	Password     string `json:"-"`
	PasswordHash string `json:"password_hash"`
	IsActive     bool   `json:"is_active"`
	Status       string `json:"status"`
	Avatar       string `json:"avatar"`
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
