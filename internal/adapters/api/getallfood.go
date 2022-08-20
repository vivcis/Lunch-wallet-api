package api

import (
	"github.com/decadevs/lunch-api/internal/core/helpers"
	"github.com/gin-gonic/gin"
	"time"
)

// GetAllFood godoc
// @Summary      Gets all the food in the database using the date of the present day
// @Description  This should be used to get all the food in the database meant for today. This should be used instead of GetBrunch and GetDinner seperately for scalability purpose when rendering on the Beneficiary dashboard. Frontend can seperate dinner and brunch.
// @Tags         Users
// @Accept       json
// @Produce      json
// @Success      200  {object} []models.Food string "Food successfully gotten"
// @Failure      500  {string}  string "internal server error"
// @Router       /user/allfood [get]
func (u *HTTPHandler) GetAllFoodHandler(c *gin.Context) {
	year, month, day := time.Now().Date()

	food, err := u.UserService.FindAllFoodByDate(year, int(month), day)
	if err != nil {
		helpers.JSON(c, "internal server error", 500, nil, []string{"internal server error"})
		return
	}

	helpers.JSON(c, "Food successfully gotten", 200, food, nil)

}
