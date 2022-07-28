package api

import (
	"github.com/decadevs/lunch-api/internal/core/helpers"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func (u HTTPHandler) BeneficiaryQRBrunch(c *gin.Context) {
	foodBeneficiary, err := u.GetBenefactorFromContext(c)
	if err != nil {
		helpers.JSON(c, "you are not logged in", http.StatusBadRequest, nil, []string{"bad request"})
		return
	}
	date := time.Now().Format("2006-01-02")

	mealRecord, mealErr := u.UserService.FindFoodBenefactorMealRecord(foodBeneficiary.Email, date)
	if mealErr != nil {
		Cerr := u.UserService.CreateFoodBenefactorBrunchMealRecord(foodBeneficiary)
		if Cerr != nil {
			helpers.JSON(c, "internal server error", http.StatusInternalServerError, nil, []string{"internal server error"})
			return
		}
		helpers.JSON(c, "success", http.StatusOK, nil, []string{"success"})
		return
	}

	if mealRecord.Brunch {
		helpers.JSON(c, "brunch already served", http.StatusBadRequest, nil, []string{"brunch already served"})
		return
	} else {
		Cerr := u.UserService.CreateFoodBenefactorBrunchMealRecord(foodBeneficiary)
		if Cerr != nil {
			helpers.JSON(c, "internal server error", http.StatusInternalServerError, nil, []string{"internal server error"})
			return
		}
		helpers.JSON(c, "success", http.StatusOK, nil, []string{"success"})
		return
	}

}

func (u HTTPHandler) BeneficiaryQRDinner(c *gin.Context) {

}
