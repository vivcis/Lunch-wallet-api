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

func (u *HTTPHandler) FoodBeneficiarySignUpHandler(c *gin.Context) {
	user := &models.User{}
	err := c.ShouldBindJSON(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Unable to bind JSON",
		})
		return
	}
	if user.FullName == "" || user.Password == "" || user.Email == "" || user.Location == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Enter all fields",
		})
		return
	}
	validEmail := user.ValidMailAddress()
	if validEmail == false {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "enter valid email",
		})
		return
	}

	_, err = u.DB.FindUserByFullName(user.FullName)
	if err == nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "fullname exists",
		})
		return
	}
	_, err = u.DB.FindUserByEmail(user.Email)
	if err == nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "email exists",
		})
		return
	}

	_, err = u.DB.FindUserByLocation(user.Location)
	if err == nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "select a location",
		})
		return
	}
	_, err = u.DB.CreateUser(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "could not create user",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Sign Up Successful",
	})

}