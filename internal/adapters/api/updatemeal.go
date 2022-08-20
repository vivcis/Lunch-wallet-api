package api

import (
	"github.com/decadevs/lunch-api/internal/core/helpers"
	"github.com/decadevs/lunch-api/internal/core/models"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

// UpdateMeal godoc
// @Summary      Admin updates meal
// @Description  Update meal by clicking on the particular food. The id of each food is attached to the link/endpoint. When you click on the food to be updated, the id is gotten from the link/endpoint. It is an authorized route to only ADMIN
// @Tags         Food
// @Accept       json
// @Success      200  {string} string "Successfully updated"
// @Failure      500  {string}  string "internal server error"
// @Failure      401  {string}  string "not authorized"
// @Router       /admin/updatemeal/:id [put]
func (u *HTTPHandler) UpdateMeal(c *gin.Context) {
	user, err := u.GetAdminFromContext(c)
	if err != nil {
		helpers.JSON(c, "not authorized", http.StatusUnauthorized, nil, []string{"not authorized"})
		return
	}

	id := c.Param("id")

	var food models.Food

	if err = c.BindJSON(&food); err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "error binding in updateFood handler"})
		return
	}

	err = u.UserService.UpdateMeal(id, food)
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

	helpers.JSON(c, "Successful updated", http.StatusOK, "Successfully updated", nil)

}
