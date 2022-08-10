package api

import (
	"github.com/decadevs/lunch-api/internal/core/helpers"
	"github.com/gin-gonic/gin"
	"net/http"
)

// NumberOfBlockedBeneficiaries godoc
// @Summary      Gets number of blocked benefeciary
// @Description  Admin gets to see how manuy beneficiaries blocked. It is an authorized route to only ADMIN
// @Tags         Users
// @Accept       json
// @Produce      json
// @Success      200  {number} string "successfully gotten"
// @Failure      500  {string}  string "internal server error"
// @Failure      400  {string}  string "bad request"
// @Router       /admin/numberblocked [get]
func (u *HTTPHandler) GetNumberOfBlockedUsers(c *gin.Context) {

	_, err := u.GetAdminFromContext(c)
	if err != nil {
		helpers.JSON(c, "not authorized", http.StatusUnauthorized, nil, []string{"not authorized"})
		return
	}

	num, err := u.UserService.NumberOfBlockedBeneficiary()
	if err != nil {
		helpers.JSON(c, "internal server error", http.StatusInternalServerError, nil, []string{"an error occurred while getting total number of blocked users"})
		return
	}

	helpers.JSON(c, "successfully gotten", http.StatusOK, num, nil)

}
