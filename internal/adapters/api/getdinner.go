package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func (u *HTTPHandler) GetDinnerHandle(c *gin.Context) {
	_, err := u.GetBenefactorFromContext(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, []string{"internal server error"})
		return
	}

	year, month, day := time.Now().Date()

	food, err := u.UserService.FindDinnerByDate(year, month, day)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, food)

}
