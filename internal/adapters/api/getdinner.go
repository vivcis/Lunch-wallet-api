package api

import (
	"github.com/decadevs/lunch-api/internal/core/helpers"
	"github.com/gin-gonic/gin"
	"time"
)

// GetDinner godoc
// @Summary      Gets all the dinner in the database using the date of the present day
// @Description  Gets all the dinner in the database meant for today. GetAllFood should be used instead for scalability purpose when rendering on the Beneficiary dashboard. Frontend can filter brunch and dinner.
// @Tags         Users
// @Accept       json
// @Produce      json
// @Success      200  {object} []models.Food string "Dinner found"
// @Failure      500  {string}  string "internal server error"
// @Router       /user/dinner [get]
func (u *HTTPHandler) GetDinnerHandle(c *gin.Context) {

	year, month, day := time.Now().Date()

	food, err := u.UserService.FindDinnerByDate(year, int(month), day)
	if err != nil {
		helpers.JSON(c, "internal server error", 500, nil, []string{"internal server error"})
		return
	}

	helpers.JSON(c, "Dinner found", 200, food, nil)

}
