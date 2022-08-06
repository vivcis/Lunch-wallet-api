package api

import (
	"github.com/decadevs/lunch-api/internal/core/helpers"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func (u *HTTPHandler) GetNotification(c *gin.Context) {
	year, month, day := time.Now().Date()
	notification, err := u.UserService.FindNotificationDate(year, month, day)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error in getting Notification"})
		return
	}

	helpers.JSON(c, "notification successfully loaded", http.StatusOK, notification, nil)

}
