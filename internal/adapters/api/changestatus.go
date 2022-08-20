package api

import (
	"github.com/decadevs/lunch-api/internal/core/helpers"
	"github.com/decadevs/lunch-api/internal/core/models"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"

	"github.com/decadevs/lunch-api/internal/adapters/repository"
)

// UpdaateFoodStatus godoc
// @Summary      Update Food
// @Description  Change brunch food status by kitchen staff from NOT SERVING to either SERVED or SERVING
// @Tags         Food
// @Accept       json
// @Produce      json
// @Param food body []models.Food true "status"
// @Success      200  {string}  string "success"
// @Failure      400  {string}  string "error"
// @Failure      404  {string}  string "error"
// @Failure      500  {string}  string "error"
// @Router       /staff/changebrunchstatus [put]
func (u *HTTPHandler) UpdateBrunchFoodStatus(c *gin.Context) {
	user, err := u.GetKitchenStaffFromContext(c)
	log.Println("user in context", user)
	if err != nil {
		helpers.JSON(c, "An internal server error", 500, nil, []string{"internal server error"})
		return
	}
	year, mon, day := time.Now().Date()
	month := int(mon)

	food, err := u.UserService.FindBrunchByDate(year, month, day)
	if err != nil {
		helpers.JSON(c, "internal server error", 500, nil, []string{"internal server error"})
		return
	}

	type StatusUpdate struct {
		Status string `json:"status"`
	}

	var status StatusUpdate

	if err := c.BindJSON(&status); err != nil {
		helpers.JSON(c, "An internal server error", 500, nil, []string{"internal server error"})
		return
	}

	inputField, err := repository.StatusEnum(status.Status)
	if err != nil {
		helpers.JSON(c, "This is an internal server error", 500, inputField, []string{"incorrect status field"})
		return
	}

	errS := u.UserService.UpdateStatus(food, status.Status)
	if errS != nil {
		helpers.JSON(c, "An internal server error", 500, nil, []string{"error updating food"})
		return
	}

	notification := models.Notification{
		Message: status.Status,
		Year:    year,
		Month:   month,
		Day:     day,
	}

	err = u.UserService.CreateNotification(notification)

	currentFood, _ := u.UserService.FindBrunchByDate(year, int(month), day)
	helpers.JSON(c, "food status updated successfully", http.StatusOK, currentFood, nil)
}

// UpdaateFoodStatus godoc
// @Summary      Update Food
// @Description  Change dinner food status by kitchen staff from NOT SERVING to either SERVED or SERVING
// @Tags         Food
// @Accept       json
// @Produce      json
// @Param food body []models.Food true "status"
// @Success      200  {string}  string "success"
// @Failure      400  {string}  string "error"
// @Failure      404  {string}  string "error"
// @Failure      500  {string}  string "error"
// @Router       /changedinnerstatus [put]
func (u *HTTPHandler) UpdateDinnerFoodStatus(c *gin.Context) {
	user, err := u.GetKitchenStaffFromContext(c)
	log.Println("user in context", user)
	if err != nil {
		helpers.JSON(c, "An internal server error", 500, nil, []string{"internal server error"})
		return
	}
	year, mon, day := time.Now().Date()
	month := int(mon)

	food, err := u.UserService.FindDinnerByDate(year, month, day)
	if err != nil {
		helpers.JSON(c, "internal server error", 500, nil, []string{"internal server error"})
		return
	}

	type StatusUpdate struct {
		Status string `json:"status"`
	}

	var status StatusUpdate

	if err := c.BindJSON(&status); err != nil {
		helpers.JSON(c, "internal server error", 500, nil, []string{"internal server error"})
		return
	}

	errS := u.UserService.UpdateStatus(food, status.Status)
	if errS != nil {
		helpers.JSON(c, "internal server error", 500, nil, []string{"error updating food"})
		return
	}

	notification := models.Notification{
		Message: status.Status,
		Year:    year,
		Month:   month,
		Day:     day,
	}

	err = u.UserService.CreateNotification(notification)
	if err != nil {
		helpers.JSON(c, "Error in getting Notification", http.StatusInternalServerError, nil, []string{"Error in getting Notification"})
		return
	}

	currentFood, _ := u.UserService.FindDinnerByDate(year, month, day)
	helpers.JSON(c, "food status updated successfully", http.StatusOK, currentFood, nil)
}
