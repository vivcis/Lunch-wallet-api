package api

import (
	"github.com/decadevs/lunch-api/internal/core/helpers"
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetTotalNumberOfUsers godoc
// @Summary      Get total number of food beneficiaries
// @Description  Kitchen staff gets total number of food beneficiaries. It is an authorized route to only KITCHEN STAFF
// @Tags         Users
// @Accept       json
// @Produce      json
// @Success      200  {number} string "Total number of users"
// @Failure      500  {string}  string "internal server error"
// @Failure      400  {string}  string "bad request"
// @Router       /staff/gettotalusers [get]
func (u *HTTPHandler) GetTotalNumberOfUsers(c *gin.Context) {
	_, err := u.GetKitchenStaffFromContext(c)
	if err != nil {
		helpers.JSON(c, "An internal server error", 500, nil, []string{"internal server error"})
		return
	}

	totalNumber, err := u.UserService.GetTotalUsers()
	if err != nil {
		helpers.JSON(c, "internal server error", 500, nil, []string{"Unable to get total number of users"})
		return
	}
	helpers.JSON(c, "Total number of users", http.StatusOK, totalNumber, nil)
}
