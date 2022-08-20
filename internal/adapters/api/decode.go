package api

import (
	"github.com/decadevs/lunch-api/internal/core/helpers"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func (u *HTTPHandler) decode(c *gin.Context, v interface{}) []string {
	if err := c.ShouldBindJSON(v); err != nil {
		var errs []string
		verr, ok := err.(validator.ValidationErrors)
		if ok {
			for _, fieldErr := range verr {
				errs = append(errs, helpers.NewFieldError(fieldErr).String())
			}
		} else {
			errs = append(errs, "internal server error")
		}
		return errs
	}
	return nil
}
