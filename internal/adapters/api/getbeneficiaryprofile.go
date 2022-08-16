package api

import (
	"github.com/decadevs/lunch-api/internal/core/helpers"
	"github.com/decadevs/lunch-api/internal/core/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetFoodBeneficiaryProfile godoc
// @Summary      Gets profile of a food beneficiary
// @Description  Admin gets to see the profile information of a food beneficiary. It is an authorized route to only ADMIN
// @Tags         Users
// @Accept       json
// @Produce      json
// @Success      200  {number} string "successfully gotten"
// @Failure      500  {string}  string "internal server error"
// @Failure      400  {string}  string "bad request"
// @Router       /admin/getfoodbeneficiaryprofile/:id [get]
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
