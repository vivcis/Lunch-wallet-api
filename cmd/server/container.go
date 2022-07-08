package server

import (
	"github.com/decadevs/lunch-api/internal/adapters/api"
	"github.com/decadevs/lunch-api/internal/adapters/repository"
	"github.com/decadevs/lunch-api/internal/core/service"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var DBConnection *gorm.DB

// Injection inject all dependencies
func Injection() {
	var (
		ginRoutes      = NewGinRouter(gin.Default())
		userRepository = repository.NewUser(DBConnection)
		userService    = service.NewUserService(userRepository)
		Handler        = api.NewHTTPHandler(userService)
	)

	router := ginRoutes.GROUP("/api/v1")
	user := router.Group("/user")
	user.GET("/:id", Handler.GetByID)

	err := ginRoutes.SERVE()

	if err != nil {
		return
	}
}
