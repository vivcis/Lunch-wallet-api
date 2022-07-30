package api

import (
	"github.com/decadevs/lunch-api/internal/core/helpers"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
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
	_, err := u.GetBenefactorFromContext(c)
	if err != nil {
		helpers.JSON(c, "internal server error", 500, nil, []string{"internal server error"})
		return
	}

	year, month, day := time.Now().Date()

	food, err := u.UserService.FindFoodByDate(year, month, day)
	if err != nil {
		helpers.JSON(c, "internal server error", 500, nil, []string{"internal server error"})
		return
	}

	helpers.JSON(c, "Dinner found", 200, food, nil)

}

func (u *HTTPHandler) GetTickets(c *gin.Context) {

}
