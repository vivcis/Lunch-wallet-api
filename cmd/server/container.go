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

	mailerRepository := repository.NewMail()
	mailerService := service.NewMailerService(mailerRepository)

	Handler := api.NewHTTPHandler(userService, mailerService)
	router := SetupRouter(Handler, userService)

	_ = router.Run(":" + helpers.Instance.Port)
}
