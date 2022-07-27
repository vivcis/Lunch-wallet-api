package api

import (
	"github.com/decadevs/lunch-api/internal/core/helpers"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (u HTTPHandler) BeneficiaryQRBrunch(c *gin.Context) {
	// get user info from token
	//query db user table with email and get user
	// get_user.id
	// use user.id to query db meal_record table
	// if nil,create a meal record for the user with user.id
	//check if meal_record.brunch is true then return error else update to true
	foodBeneficiary, err := u.GetBenefactorFromContext(c)
	if err != nil {
		helpers.JSON(c, "you are not logged in", http.StatusBadRequest, nil, []string{"bad request"})
		return
	}

	mealRecords, mealErr := u.UserService.FindFoodBenefactorMealRecord(foodBeneficiary.Email)
	if mealErr != nil {
		helpers.JSON(c, "internal server error", http.StatusInternalServerError, nil, []string{"internal server error"})
		return
	}

	if mealRecords.UserEmail == "" {
		Cerr := u.UserService.CreateFoodBenefactorBrunchMealRecord(foodBeneficiary)
		if Cerr != nil {
			helpers.JSON(c, "internal server error", http.StatusInternalServerError, nil, []string{"internal server error"})
			return
		}
		helpers.JSON(c, "success", http.StatusOK, nil, []string{"success"})
		return
	}
	if mealRecords

}

func (u HTTPHandler) BeneficiaryQRDinner(c *gin.Context) {

}
