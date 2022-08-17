package api

import (
	"github.com/decadevs/lunch-api/internal/core/helpers"
	"github.com/gin-gonic/gin"
	"time"
)

func (u *HTTPHandler) GetBrunchTimetable(c *gin.Context) {
	_, err := u.GetKitchenStaffFromContext(c)
	if err != nil {
		helpers.JSON(c, "internal server error", 500, nil, []string{"internal server error"})
		return
	}
	year, month, day := time.Now().Date()
	food, err := u.UserService.FindBrunchByDate(year, month, day)
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
