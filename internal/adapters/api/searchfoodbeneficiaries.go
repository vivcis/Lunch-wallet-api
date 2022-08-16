package api

import (
	"github.com/decadevs/lunch-api/internal/adapters/repository"
	"github.com/decadevs/lunch-api/internal/core/helpers"
	"github.com/gin-gonic/gin"
	"net/http"
)

// SearchFoodBeneficiaries godoc
// @Summary      Search Food Beneficiary
// @Description  Kitchen staff can search for a food beneficiary by name, location or stack. It is an authorized route to only KITCHEN STAFF
// @Tags         Users
// @Accept       json
// @Produce      json
// @Success      200  {number} string "information gotten"
// @Failure      500  {string}  string "internal server error"
// @Failure      400  {string}  string "bad request"
// @Router       /staff/searchbeneficiary/:text [get]
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
