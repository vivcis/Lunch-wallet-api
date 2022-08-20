package api

import (
	"github.com/decadevs/lunch-api/internal/core/helpers"
	"github.com/decadevs/lunch-api/internal/core/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

// DeleteMeal godoc
// @Summary      Admin deletes meal
// @Description  Delete meal by clicking on the particular food. The id of each food is attached to the link/endpoint. When you click on the food to be deleted, the id is gotten from the link/endpoint and the food is deleted. It is an authorized route to only ADMIN
// @Tags         Food
// @Accept       json
// @Success      200  {string} string "Successfully Deleted"
// @Failure      500  {string}  string "internal server error"
// @Failure      401  {string}  string "not authorized"
// @Router       /admin/deletemeal/:id [delete]
func (u *HTTPHandler) DeleteMeal(c *gin.Context) {
	user, err := u.GetAdminFromContext(c)
	if err != nil {
		helpers.JSON(c, "not authorized", http.StatusUnauthorized, nil, []string{"not authorized"})
		return
	}

	id := c.Param("id")
	err = u.UserService.DeleteMeal(id)
	if err != nil {
		helpers.JSON(c, "internal server error", http.StatusInternalServerError, nil, []string{"internal server error"})
		return
	}

	year, month, day := time.Now().Date()

	notification := models.Notification{
		Message: user.FullName + " updated timetable",
		Year:    year,
		Month:   int(month),
		Day:     day,
	}

	err = u.UserService.CreateNotification(notification)
	if err != nil {
		helpers.JSON(c, "Error in getting Notification", http.StatusInternalServerError, nil, []string{"Error in getting Notification"})
		return
	}

	helpers.JSON(c, "Successfully Deleted", http.StatusOK, "Successfully Deleted", nil)

}
