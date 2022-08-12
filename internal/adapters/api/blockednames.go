package api

import (
	"github.com/decadevs/lunch-api/internal/core/helpers"
	"github.com/gin-gonic/gin"
	"net/http"
)

// NamesOfBlockedBeneficiaries godoc
// @Summary      This endpoint enables admin to see all blocked users
// @Description  Admin gets to see all blocked users with this endpoint. It is an authorized route to only ADMIN
// @Tags         Users
// @Accept       json
// @Produce      json
// @Success      200  {object} []models.FoodBeneficiary string "blocked users successfully gotten"
// @Failure      500  {string}  string "internal server error"
// @Failure      400  {string}  string "bad request"
// @Router       /admin/blockedusers [get]
func (u *HTTPHandler) GetBlockedUsers(c *gin.Context) {

	_, err := u.GetAdminFromContext(c)
	if err != nil {
		helpers.JSON(c, "not authorized", http.StatusUnauthorized, nil, []string{"not authorized"})
		return
	}

	users, err := u.UserService.GetBlockedBeneficiary()
	if err != nil {
		helpers.JSON(c, "internal server error", http.StatusInternalServerError, nil, []string{"an error occurred while getting blocked users"})
		return
	}

	helpers.JSON(c, "blocked users successfully gotten", http.StatusOK, users, nil)

}
