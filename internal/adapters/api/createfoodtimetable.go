package api

import (
	"github.com/decadevs/lunch-api/internal/core/models"
	"github.com/gin-gonic/gin"
	"net/http"
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
		Name  string `json:"name" binding:"required"`
		Type  string `json:"type" binding:"required"`
		Date  int    `json:"date" binding:"required"`
		Month int    `json:"month" binding:"required"`
		Year  int    `json:"year" binding:"required"`
	}{}

	err = c.ShouldBindJSON(&foodTimetable)
	if err != nil {
		c.JSON(400, gin.H{"message": "bad request"})
		return
	}

	food.CreatedAt = time.Now()
	var location time.Location
	food.Date = time.Date(foodTimetable.Year, time.Month(foodTimetable.Month), foodTimetable.Date, 0, 0, 0, 0, &location)
	food.AdminName = admin.FullName
	food.Name = foodTimetable.Name
	food.Type = foodTimetable.Type
	err = u.UserService.CreateFoodTimetable(food)
	if err != nil {
		c.JSON(400, gin.H{"message": "bad request"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Successfully Created"})
}
