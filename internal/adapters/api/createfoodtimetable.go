package api

import (
	"github.com/decadevs/lunch-api/internal/core/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"time"
)

func (u HTTPHandler) CreateFoodTimetableHandle(c *gin.Context) {
	admin, err := u.GetAdminFromContext(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, []string{"internal server error"})
		return
	}
	var food models.Food
	foodTimetable := &struct {
		Name    string `json:"name" binding:"required"`
		Type    string `json:"type" binding:"required"`
		Date    int    `json:"date" binding:"required"`
		Month   int    `json:"month" binding:"required"`
		Year    int    `json:"year" binding:"required"`
		Weekday string `json:"weekday" binding:"required"`
	}{}

	err = c.ShouldBindJSON(&foodTimetable)
	if err != nil {
		c.JSON(400, gin.H{"message": "bad request"})
		return
	}

	foodType := strings.ToLower(foodTimetable.Type)
	food.CreatedAt = time.Now()
	food.Name = foodTimetable.Name
	food.Type = foodType
	food.AdminName = admin.FullName
	food.Year = foodTimetable.Year
	food.Month = time.Month(foodTimetable.Month)
	food.Day = foodTimetable.Date
	food.Weekday = foodTimetable.Weekday
	err = u.UserService.CreateFoodTimetable(food)
	if err != nil {
		c.JSON(400, gin.H{"message": "bad request"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Successfully Created"})
}
