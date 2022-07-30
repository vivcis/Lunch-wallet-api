package api

import (
	"github.com/decadevs/lunch-api/internal/core/helpers"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (u *HTTPHandler) GetAllBeneficiaryHandle(c *gin.Context) {

	Beneficiaries, err := u.UserService.GetAllFoodBeneficiaries()
	if err != nil {
		helpers.JSON(c, "Error Exist in Getting All Sellers", http.StatusInternalServerError, err.Error(), nil)
		return
	}

	helpers.JSON(c, "Beneficiaries found", http.StatusOK, Beneficiaries, nil)

}
func (u *HTTPHandler) GetMealTimetableHandle(c *gin.Context) {

}
