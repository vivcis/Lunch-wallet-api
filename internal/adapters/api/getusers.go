package api

import (
	"github.com/decadevs/lunch-api/internal/core/helpers"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (u *HTTPHandler) GetUsers(c *gin.Context) {
	_, err := u.GetBenefactorFromContext(c)
	if err != nil {
		helpers.JSON(c, "internal server error", 500, nil, []string{"internal server error"})
		return
	}

	user, err := u.UserService.FindAllFoodBeneficiary(nil)
	if err != nil {
		helpers.JSON(c, "internal server error", 500, nil, []string{"internal server error"})
		return
	}
	helpers.JSON(c, "food beneficiary found successfully", http.StatusOK, user, nil)
}
