package api

import (
	"github.com/decadevs/lunch-api/internal/core/helpers"
	"github.com/gin-gonic/gin"
	"net/http"
)

// RemoveFoodBeneficiary godoc
// @Summary      Remove a Food Beneficiary
// @Description  Admin deletes/removes a food beneficiary from the dashboard. It is an authorized route to only ADMIN
// @Tags         Users
// @Accept       json
// @Produce      json
// @Success      200  {number} string "food beneficiary removed"
// @Failure      500  {string}  string "internal server error"
// @Failure      400  {string}  string "bad request"
// @Router       /staff/removefoodbeneficiary/{id} [delete]
func (u *HTTPHandler) RemoveFoodBeneficiary(c *gin.Context) {
	_, err := u.GetAdminFromContext(c)
	if err != nil {
		helpers.JSON(c, "internal server error", 500, nil, []string{"internal server error"})
		return
	}
	id := c.Param("id")
	err = u.UserService.AdminRemoveFoodBeneficiary(id)
	if err != nil {
		helpers.JSON(c, "internal server error", http.StatusInternalServerError, nil, []string{"internal server error"})
		return
	}
	helpers.JSON(c, "food beneficiary removed", http.StatusOK, nil, nil)
}
