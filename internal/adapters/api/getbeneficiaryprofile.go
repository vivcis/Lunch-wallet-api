package api

import (
	"github.com/decadevs/lunch-api/internal/core/helpers"
	"github.com/decadevs/lunch-api/internal/core/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (u *HTTPHandler) GetFoodBeneficiaryProfile(c *gin.Context) {
	_, err := u.GetAdminFromContext(c)
	if err != nil {
		helpers.JSON(c, "internal server error", 500, nil, []string{"internal server error"})
		return
	}

	userID := c.Param("id")

	user, err := u.UserService.GetFoodBenefactorById(userID)
	if err != nil {
		helpers.JSON(c, "This is an internal server error", 500, nil, []string{"internal server error"})
		return
	}

	var userprofile models.UserProfile
	userprofile = helpers.GetUserProfileFromBeneficiary(user)
	helpers.JSON(c, "food beneficiary details retrieved correctly", http.StatusOK, userprofile, nil)
}
