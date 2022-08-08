package api

import (
	"github.com/decadevs/lunch-api/internal/adapters/repository"
	"github.com/decadevs/lunch-api/internal/core/helpers"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (u *HTTPHandler) SearchFoodBeneficiaries(c *gin.Context) {
	_, err := u.GetKitchenStaffFromContext(c)
	if err != nil {
		helpers.JSON(c, "An internal server error", 500, nil, []string{"internal server error"})
		return
	}
	pagination := repository.GeneratePaginationFromRequest(c)
	query := c.Param("text")

	beneficiaries, err := u.UserService.SearchFoodBeneficiary(query, &pagination)
	if err != nil {
		helpers.JSON(c, "An internal server error", 500, nil, []string{"internal server error"})
		return
	}
	if len(beneficiaries) == 0 {
		helpers.JSON(c, "Record Not Found", 404, nil, []string{"Record Not Found"})
		return
	}
	helpers.JSON(c, "information gotten", http.StatusOK, beneficiaries, nil)
}
