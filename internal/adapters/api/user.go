package api

import (
	"errors"
	"github.com/decadevs/lunch-api/internal/core/helpers"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

func (u HTTPHandler) GetByID(c *gin.Context) {
	id := c.Param("id")
	if err := c.ShouldBindUri(id); err != nil {
		helpers.JSON(c, "", http.StatusBadRequest, nil, []string{err.Error()})
		return
	}

	user, err := u.UserService.GetByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			helpers.JSON(c, "User with the id does not exist", http.StatusNotFound, nil, []string{err.Error()})
			return
		}
		helpers.JSON(c, "Error getting user", http.StatusInternalServerError, nil, []string{err.Error()})
		return
	}

	helpers.JSON(c, "User found successfully", http.StatusOK, user, nil)
}
