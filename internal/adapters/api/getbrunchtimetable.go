package api

import (
	"github.com/decadevs/lunch-api/internal/core/helpers"
	"github.com/gin-gonic/gin"
	"time"
)

// GetBrunchTimetable godoc
// @Summary      Gets food for brunch
// @Description  Kitchen staff gets the timetable for brunch for a particular date. It is an authorized route to only KITCHEN STAFF
// @Tags         Users
// @Accept       json
// @Produce      json
// @Success      200  {number} string "successfully gotten"
// @Failure      500  {string}  string "internal server error"
// @Failure      400  {string}  string "bad request"
// @Router       /staff/getbrunchtimetable [get]
func (u *HTTPHandler) GetBrunchTimetable(c *gin.Context) {
	_, err := u.GetKitchenStaffFromContext(c)
	if err != nil {
		helpers.JSON(c, "internal server error", 500, nil, []string{"internal server error"})
		return
	}
	year, month, day := time.Now().Date()
	food, err := u.UserService.FindBrunchByDate(year, int(month), day)
	if err != nil {
		helpers.JSON(c, "internal server error", 500, nil, []string{"internal server error"})
		return
	}
	if len(food) > 0 {
		helpers.JSON(c, "food found", 200, nil, []string{"food found"})
		return
	}
	helpers.JSON(c, "Brunch found", 200, food, nil)
}
