package api

import (
	"fmt"
	"github.com/decadevs/lunch-api/internal/core/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func (u HTTPHandler) GetAdminFromContext(c *gin.Context) (*models.Admin, error) {
	userI, exists := c.Get("user")
	if !exists {
		return nil, fmt.Errorf("error getting user from context")
	}
	user, ok := userI.(*models.Admin)
	if !ok {
		return nil, fmt.Errorf("an error occurred")
	}
	return user, nil
}

func (u HTTPHandler) CreateFoodTimetableHandle(c *gin.Context) {
	admin, err := u.GetAdminFromContext(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, []string{"internal server error"})
		return
	}
	var food models.Food
	foodTimetable := &struct {
		Name  string `json:"name" binding:"required"`
		Type  string `json:"type" binding:"required"`
		Date  int    `json:"date" binding:"required"`
		Month int    `json:"month" binding:"required"`
		Year  int    `json:"year" binding:"required"`
	}{}
	//today,_,_:=time.Now().Date()
	food.CreatedAt = time.Now()
	food.Date = time.Now().AddDate(foodTimetable.Year, foodTimetable.Month, foodTimetable.Date)
	food.AdminName = admin.FullName

	err = u.UserService.CreateFoodTimetable(food)
	if err != nil {
		c.JSON(400, gin.H{"message": "bad request"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Successfully Created"})
}
