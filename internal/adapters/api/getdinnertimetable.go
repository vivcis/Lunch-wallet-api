package api

import (
	"github.com/decadevs/lunch-api/internal/core/helpers"
	"github.com/gin-gonic/gin"
	"time"
)

// GetDinnerTimetable godoc
// @Summary      Gets food for dinner
// @Description  Kitchen staff gets the timetable for dinner for a particular date. It is an authorized route to only KITCHEN STAFF
// @Tags         Users
// @Accept       json
// @Produce      json
// @Success      200  {number} string "successfully gotten"
// @Failure      500  {string}  string "internal server error"
// @Failure      400  {string}  string "bad request"
// @Router       /staff/getdinnertimetable [get]
func (u *HTTPHandler) GetDinnerTimetable(c *gin.Context) {
	_, err := u.GetKitchenStaffFromContext(c)
	if err != nil {
		helpers.JSON(c, "internal server error", 500, nil, []string{"internal server error"})
		return
	}
	year, month, day := time.Now().Date()
	food, err := u.UserService.FindDinnerByDate(year, int(month), day)
	if err != nil {
		helpers.JSON(c, "internal server error", 500, nil, []string{"internal server error"})
		return
	}
	helpers.JSON(c, "Dinner found", 200, food, nil)
}
