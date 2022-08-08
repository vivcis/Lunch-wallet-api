package api

import (
	"fmt"
	"net/http"

	"github.com/decadevs/lunch-api/internal/core/helpers"
	"github.com/gin-gonic/gin"
)

func (u *HTTPHandler) BlockFoodBeneficiary(c *gin.Context) {
	_, err := u.GetAdminFromContext(c)
	if err != nil {
		helpers.JSON(c, "internal server error", 500, nil, []string{"internal server error"})
		return
	}

	id := c.Param("id")
	errBlock := u.UserService.AdminBlockFoodBeneficiary(id)
	if errBlock != nil {
		helpers.JSON(c, "This is an internal server error", 500, nil, []string{"internal server error"})
		fmt.Println(errBlock)
		return
	}
	helpers.JSON(c, "food beneficiary blocked", http.StatusOK, nil, nil)
}
