package api

import (
	"github.com/decadevs/lunch-api/internal/core/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (u HTTPHandler) FoodBeneficiaryForgotPassword(c *gin.Context) {
	var forgotPassword models.ResetPasswordRequest

	err := c.ShouldBindJSON(&forgotPassword)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "please fill all fields"})
		return
	}
	beneficiary, berr := u.UserService.FindFoodBenefactorByEmail(forgotPassword.Email)
	if berr != nil {

	}
}
