package api

import (
	"errors"
	"net/http"

	"github.com/decadevs/lunch-api/internal/core/helpers"
	"github.com/decadevs/lunch-api/internal/core/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func (u HTTPHandler) GetByID(c *gin.Context) {
	id := c.Param("id")
	if err := c.ShouldBindUri(id); err != nil {
		helpers.JSON(c, "", http.StatusBadRequest, nil, []string{err.Error()})
		return
	}

	user, err := u.UserService.GetByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			helpers.JSON(c, "User with the id does not exist", http.StatusNotFound, nil, []string{err.Error()})
			return
		}
		helpers.JSON(c, "Error getting user", http.StatusInternalServerError, nil, []string{err.Error()})
		return
	}

	helpers.JSON(c, "User found successfully", http.StatusOK, user, nil)
}

func (u HTTPHandler) FoodBeneficiarySignUp(c *gin.Context) {
	//user := &models.FoodBeneficiary{}
	var user *models.FoodBeneficiary
	err := c.ShouldBindJSON(&user)
	if err != nil {
		helpers.JSON(c, "Unable to bind request", 400, nil, []string{"unable to bind request: validation error"})
		return
	}

	if user.FullName == "" || user.Password == "" || user.Email == "" || user.Location == "" || user.Stack == "" {
		helpers.JSON(c, "Enter all fields", 400, nil, []string{err.Error()})
		return
	}

	validDecagonEmail := user.ValidateDecagonEmail()
	if !validDecagonEmail {
		helpers.JSON(c, "Enter valid decagon email", 400, nil, []string{err.Error()})
		return
	}

	_, Emailerr := u.UserService.FindUserByEmail(user.Email)
	if Emailerr == nil {
		helpers.JSON(c, "Email already exists", 400, nil, []string{"email exists"})
		return
	}
	if err = user.HashPassword(); err != nil {
		helpers.JSON(c, "Unable to hash password", 400, nil, []string{err.Error()})
		return
	}
	_, err = u.UserService.CreateUser(user)
	if err != nil {
		helpers.JSON(c, "Unable to create user", 400, nil, []string{"unable to create user"})
		return
	}
	helpers.JSON(c, "Signup Successful", 201, nil, nil)

}

func (u *HTTPHandler) KitchenStaffSignUp(c *gin.Context) {
	staff := &models.KitchenStaff{}
	err := c.ShouldBindJSON(staff)
	if err != nil {
		helpers.JSON(c, "Unable to bind request", 400, nil, []string{err.Error()})
		return
	}
	if staff.FullName == "" || staff.Password == "" || staff.Email == "" || staff.Location == "" {
		helpers.JSON(c, "Enter all fields", 400, nil, []string{err.Error()})
		return
	}

	validDecagonEmail := staff.ValidateDecagonEmail()
	if !validDecagonEmail {
		helpers.JSON(c, "Enter valid decagon email", 400, nil, []string{err.Error()})
		return
	}

	_, err = u.UserService.FindStaffByEmail(staff.Email)
	if err == nil {
		helpers.JSON(c, "Email exist", 400, nil, []string{"email exists"})
		return
	}

	if err = staff.HashPassword(); err != nil {
		helpers.JSON(c, "Unable to hash password", 400, nil, []string{err.Error()})
		return
	}
	_, err = u.UserService.CreateStaff(staff)
	if err != nil {
		helpers.JSON(c, "Unable to create user", 400, nil, []string{err.Error()})
		return
	}
	helpers.JSON(c, "Staff Signup Successful", 201, nil, nil)

}
