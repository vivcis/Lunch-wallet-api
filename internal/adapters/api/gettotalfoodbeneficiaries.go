package api

import (
	"github.com/decadevs/lunch-api/internal/core/helpers"
	"github.com/gin-gonic/gin"
	"net/http"
)

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
