package helpers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"time"
)

func JSON(c *gin.Context, message string, status int, data interface{}, errs []string) {
	responsedata := gin.H{
		"message":   message,
		"data":      data,
		"errors":    errs,
		"status":    http.StatusText(status),
		"timestamp": time.Now().Format("2006-01-02 15:04:05"),
	}

	c.JSON(status, responsedata)
}

func CheckFoodStatus(status string) (string, error) {
	var stat string
	switch strings.ToLower(status) {
	case "serving":
		stat = "Serving"
	case "served":
		stat = "Served"
	default:
		return stat, errors.New("not serving")
	}
	return stat, nil
}
