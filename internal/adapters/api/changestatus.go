package api

import (
	"fmt"
	"github.com/decadevs/lunch-api/internal/core/helpers"
	"github.com/decadevs/lunch-api/internal/core/models"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

//func (u *HTTPHandler) ChangeFoodStatus(c *gin.Context) {
//	user, err := u.GetFoodFromContext(c)
//	if err != nil {
//		helpers.JSON(c, "An internal server error", 500, nil, []string{"internal server error"})
//		return
//	}
//	status := c.Param("status")
//	Id := c.Param("id")
//	foodStatus, err := helpers.CheckFoodStatus(status)
//	if err != nil {
//		helpers.JSON(c, "Bad request", 400, nil, []string{"bad request"})
//		return
//	}
//	food, er := u.UserService.GetFoodByID(Id)
//	if er != nil {
//		helpers.JSON(c, "This is an internal server error", 500, nil, []string{"internal server error"})
//		return
//	}
//	foodError := u.UserService.UpdateFoodStatusById(food.ID, foodStatus)
//	if foodError != nil {
//		helpers.JSON(c, "internal server error", 500, nil, []string{"internal server error"})
//		return
//	}
//	helpers.JSON(c, "food status updated successfully", http.StatusOK, user, nil)
//}

func (u *HTTPHandler) UpdateFoodStatus(c *gin.Context) {
	user, err := u.GetKitchenStaffFromContext(c)
	log.Println("user in context", user)
	if err != nil {
		helpers.JSON(c, "An internal server error", 500, nil, []string{"internal server error"})
		return
	}

	year, month, day := time.Now().Date()

	food, err := u.UserService.FindBrunchByDate(year, month, day)
	if err != nil {
		helpers.JSON(c, "internal server error", 500, nil, []string{"internal server error"})
		return
	}
	//status := c.Param("status")
	//Id := c.Param("id")
	//foodId, _ := strconv.Atoi(Id)
	//foodID := uint(foodId)

	var food = models.Food{}
	if err := c.BindJSON(&food); err != nil {
		helpers.JSON(c, "An internal server error", 500, nil, []string{"internal server error"})
		return
	}
	foodRecord, statusError := u.UserService.GetFoodByID(user.ID)
	if statusError != nil {
		helpers.JSON(c, "An internal server error", 500, nil, []string{"Food does not exist"})
		return
	}

	fmt.Println(food.Name)

	errS := u.UserService.UpdateFoodStatusById(foodRecord, status)
	if errS != nil {
		helpers.JSON(c, "An internal server error", 500, nil, []string{"error updating food"})
		return
	}
	helpers.JSON(c, "food status updated successfully", http.StatusOK, food, nil)
}
