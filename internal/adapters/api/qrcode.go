package api

import (
	"github.com/decadevs/lunch-api/internal/core/helpers"
	"github.com/decadevs/lunch-api/internal/core/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"time"
)

// GetMeal godoc
// @Summary      Gets the food in the database required to generate QR code
// @Description  This should be used to get the food in the database to generate QR code meant for the day.
// @Tags         Food
//params		 mealType query string true "mealType"
// @Accept       json
// @Produce      json
// @Success      200  {object} models.Food string "success"
// @Failure      500  {string}  string "internal server error"
// @Failure      400  {string}  string "invalid meal type"
// @Router       /staff/generateqrcode [get]
func (u *HTTPHandler) GetMeal(c *gin.Context) {
	_, err := u.GetKitchenStaffFromContext(c)
	if err != nil {
		helpers.JSON(c, "internal server error", 500, nil, []string{"internal server error"})
		return
	}
	mealType := c.Query("mealType")
	mealtype := strings.ToLower(mealType)

	year, month, day := time.Now().Date()

	meals, mealErr := u.UserService.FindAllFoodByDate(year, int(month), day)
	if mealErr != nil {
		helpers.JSON(c, "internal server error", 500, nil, []string{"internal server error"})
		return
	}

	var food models.Food
	for _, meal := range meals {
		if meal.Type == mealtype {
			food = meal
			helpers.JSON(c, "success", 200, food, nil)
			return
		}
	}
	helpers.JSON(c, "bad request", 400, nil, []string{"invalid meal type"})
	return
}

// MealRecord godoc
// @Summary      Logs meal records in the database after successful QR code scan
// @Description  This is used to tell if a food beneficiary has scanned the QR code previously or not,it then logs the information .
// @Tags         user
// @Accept       json
// @Produce      json
// @Success      200  {string}  string "success"
// @Failure      500  {string}  string "internal server error"
// @Failure      400  {string}  string "meal already served"
// @Router       /benefactor/qrmealrecords [post]
func (u *HTTPHandler) MealRecord(c *gin.Context) {
	_, err := u.GetBenefactorFromContext(c)
	if err != nil {
		helpers.JSON(c, "internal server error", 500, nil, []string{"internal server error"})
		return
	}

	var mealRecord *models.QRCodeMealRecords
	err = c.ShouldBind(&mealRecord)
	if err != nil {
		helpers.JSON(c, "error", 500, nil, []string{"internal server error"})
		return
	}
	record, recordErr := u.UserService.FindFoodBenefactorQRCodeMealRecord(mealRecord.MealId, mealRecord.UserId)
	if recordErr != nil {
		Cerr := u.UserService.CreateFoodBenefactorQRMealRecord(mealRecord)
		if Cerr != nil {
			helpers.JSON(c, "internal server error", http.StatusInternalServerError, nil, []string{"internal server error"})
			return
		}
		helpers.JSON(c, "success", http.StatusOK, nil, []string{"success"})
		return
	}
	helpers.JSON(c, "error", http.StatusBadRequest, record, []string{"meal already served"})
	return
}
