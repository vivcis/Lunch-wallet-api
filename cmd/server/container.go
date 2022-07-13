package server

import (
	"github.com/decadevs/lunch-api/internal/adapters/api"
	"github.com/decadevs/lunch-api/internal/adapters/repository"
	"github.com/decadevs/lunch-api/internal/core/helpers"
	"github.com/decadevs/lunch-api/internal/core/service"
	"gorm.io/gorm"
)

// Injection inject all dependencies
func Injection(db *gorm.DB) {
	userRepository := repository.NewUser(db)
	userService := service.NewUserService(userRepository)

	Handler := api.NewHTTPHandler(userService)
	router := SetupRouter(helpers.Instance.ServiceAddress, helpers.Instance.Port, Handler)

	_ = router.Run(":" + helpers.Instance.Port)
}
