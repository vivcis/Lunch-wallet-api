package api

import (
	"github.com/decadevs/lunch-api/internal/adapters/repository"
	"github.com/decadevs/lunch-api/internal/core/helpers"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (u *HTTPHandler) PingHandler(c *gin.Context) {
	// healthcheck
	helpers.JSON(c, "pong", 200, nil, nil)
}

func (u *HTTPHandler) GetFoodBeneficiaries(c *gin.Context) {
	_, err := u.GetKitchenStaffFromContext(c)
	if err != nil {
		helpers.JSON(c, "internal server error", 500, nil, []string{"This is an internal server error"})
		return
	}
	pagination := repository.GeneratePaginationFromRequest(c)
	userLists, err := u.UserService.FindAllFoodBeneficiary(&pagination)

	if err != nil {
		helpers.JSON(c, "internal server error", 500, nil, []string{"An internal server error"})
		return
	}
	helpers.JSON(c, "food beneficiaries found successfully", http.StatusOK, userLists, nil)
}
