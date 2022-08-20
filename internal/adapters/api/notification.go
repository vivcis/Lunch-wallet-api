package api

import (
	"github.com/decadevs/lunch-api/internal/core/helpers"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

// Notification godoc
// @Summary      Notifies users whenever there is a change worthy of notification
// @Description  Returns all notifications in the database and their dates to be rendered as will by the frontend
// @Tags         Notification
// @Accept       json
// @Produce      json
// @Success      200  {object} []models.Notification string "notification successfully loaded"
// @Failure      500  {string}  string "internal server error"
// @Router       /user/notifications [get]
func (u *HTTPHandler) GetNotification(c *gin.Context) {
	year, mnth, day := time.Now().Date()

	month := int(mnth)
	notification, err := u.UserService.FindNotificationByDate(year, month, day)
	if err != nil {
		helpers.JSON(c, "Error in getting Notification", http.StatusInternalServerError, nil, []string{"Error in getting Notification"})
		return
	}

	helpers.JSON(c, "notification successfully loaded", http.StatusOK, notification, nil)

}
