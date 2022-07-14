package api

import (
	"github.com/decadevs/lunch-api/internal/core/helpers"
	"github.com/decadevs/lunch-api/internal/core/models"
	"github.com/gin-gonic/gin"
)

// FoodBeneficiarySignUp creates a new food benefactor
func (u HTTPHandler) FoodBeneficiarySignUp(c *gin.Context) {
	var user *models.FoodBeneficiary
	err := c.ShouldBindJSON(&user)
	if err != nil {
		helpers.JSON(c, "Unable to bind request", 400, nil, []string{"unable to bind request: validation error"})
		return
	}

	validDecagonEmail := user.ValidateDecagonEmail()
	if !validDecagonEmail {
		helpers.JSON(c, "Enter valid decagon email", 400, nil, []string{err.Error()})
		return
	}

	_, Emailerr := u.UserService.FindFoodBenefactorByEmail(user.Email)
	if Emailerr == nil {
		helpers.JSON(c, "Email already exists", 400, nil, []string{"email exists"})
		return
	}
	if err = user.HashPassword(); err != nil {
		helpers.JSON(c, "Unable to hash password", 400, nil, []string{err.Error()})
		return
	}
	_, err = u.UserService.CreateFoodBenefactor(user)
	if err != nil {
		helpers.JSON(c, "Unable to create user", 400, nil, []string{"unable to create user"})
		return
	}
	helpers.JSON(c, "Signup Successful", 201, nil, nil)

}

// KitchenStaffSignUp creates a new kitchen staff
func (u *HTTPHandler) KitchenStaffSignUp(c *gin.Context) {
	staff := &models.KitchenStaff{}
	err := c.ShouldBindJSON(staff)
	if err != nil {
		helpers.JSON(c, "Unable to bind request", 400, nil, []string{"unable to bind request: validation error"})
		return
	}

	validDecagonEmail := staff.ValidateDecagonEmail()
	if !validDecagonEmail {
		helpers.JSON(c, "Enter valid decagon email", 400, nil, []string{err.Error()})
		return
	}

	_, err = u.UserService.FindKitchenStaffByEmail(staff.Email)
	if err == nil {
		helpers.JSON(c, "Email exist", 400, nil, []string{"email exists"})
		return
	}

	if err = staff.HashPassword(); err != nil {
		helpers.JSON(c, "Unable to hash password", 400, nil, []string{err.Error()})
		return
	}
	_, err = u.UserService.CreateKitchenStaff(staff)
	if err != nil {
		helpers.JSON(c, "Unable to create user", 400, nil, []string{err.Error()})
		return
	}
	helpers.JSON(c, "Staff Signup Successful", 201, nil, nil)

}
