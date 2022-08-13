package api

import (
	"fmt"
	"net/http"

	"github.com/decadevs/lunch-api/internal/core/helpers"
	"github.com/gin-gonic/gin"
)

// BlockFoodBeneficiary godoc
// @Summary      Block a food beneficiary
// @Description  Admin blocks a food beneficiary
// @Tags         Users
// @Param name path string true "id"
// @Accept       json
// @Produce      json
// @Success      200  {object} userID string "food beneficiary blocked"
// @Failure      500  {string}  string "error"
// @Failure      400  {string}  string "error"
// @Router       /admin/blockfoodbeneficiary/:id [put]
func (u *HTTPHandler) BlockFoodBeneficiary(c *gin.Context) {
	_, err := u.GetAdminFromContext(c)
	if err != nil {
		helpers.JSON(c, "Bad request", 400, nil, []string{"This is a bad request"})
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
