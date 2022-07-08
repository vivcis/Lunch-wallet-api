package ports

import (
	"github.com/decadevs/lunch-api/internal/core/models"
	"github.com/gin-gonic/gin"
)

type UserRepository interface {
	GetByID(id string) (*models.User, error)
}

type Router interface {
	GET(uri string, fn func(ctx *gin.Context))
	POST(uri string, fn func(ctx *gin.Context))
	PUT(uri string, fn func(ctx *gin.Context))
	DELETE(uri string, fn func(ctx *gin.Context))
	SERVE() error
	GROUP(path string) *gin.RouterGroup
}
