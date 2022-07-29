package api

import (
	"github.com/decadevs/lunch-api/internal/core/helpers"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (u *HTTPHandler) ChangeFoodStatus(c *gin.Context) {
	user, err := u.GetFoodFromContext(c)
	if err != nil {
		helpers.JSON(c, "internal server error", 500, nil, []string{"internal server error"})
		return
	}
	status := c.Param("status")
	Id := c.Param("id")
	foodStatus, err := helpers.CheckFoodStatus(status)
	if err != nil {
		helpers.JSON(c, "Bad request", 400, nil, []string{"bad request"})
		return
	}
	food, er := u.UserService.GetFoodByID(Id)
	if er != nil {
		helpers.JSON(c, "internal server error", 500, nil, []string{"internal server error"})
		return
	}
	foodError := u.UserService.UpdateFoodStatusById(food.ID, foodStatus)
	if foodError != nil {
		helpers.JSON(c, "internal server error", 500, nil, []string{"internal server error"})
		return
	}
	helpers.JSON(c, "food status updated successfully", http.StatusOK, user, nil)
}
