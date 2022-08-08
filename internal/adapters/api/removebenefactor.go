package api

import (
	"github.com/decadevs/lunch-api/internal/core/helpers"
	"github.com/gin-gonic/gin"
	"net/http"
)

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
