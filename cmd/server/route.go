package server

import (
	"github.com/decadevs/lunch-api/internal/adapters/api"
	"github.com/decadevs/lunch-api/internal/core/middleware"
	"github.com/decadevs/lunch-api/internal/ports"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"time"
)

func SetupRouter(appPort, address string, handler *api.HTTPHandler, userService ports.UserService) *gin.Engine {
	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"POST", "GET", "PUT", "PATCH", "DELETE"},
		AllowHeaders:     []string{"*"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		// AllowOriginFunc: func(origin string) bool {
		// 	return origin == "https://github.com"
		// },
		MaxAge: 12 * time.Hour,
	}))

	r := router.Group("/api/v1")
	{
		r.POST("/user/beneficiarysignup", handler.FoodBeneficiarySignUp)
		r.POST("/user/kitchenstaffsignup", handler.KitchenStaffSignUp)
	}

	// authorizeKitchenStaff authorizes all authorized kitchen staff handler
	authorizeKitchenStaff := router.Group("/")
	authorizeKitchenStaff.Use(middleware.AuthorizeKitchenStaff(userService.FindKitchenStaffByEmail, userService.TokenInBlacklist))
	{

	}

	// authorizeBenefactor authorizes all authorized benefactor handler
	authorizeBenefactor := router.Group("/")
	authorizeBenefactor.Use(middleware.AuthorizeFoodBenefactor(userService.FindFoodBenefactorByEmail, userService.TokenInBlacklist))
	{

	}

	// authorizeAdmin authorizes all authorized admin handler
	authorizeAdmin := router.Group("/")
	authorizeAdmin.Use(middleware.AuthorizeAdmin(userService.FindAdminByEmail, userService.TokenInBlacklist))
	{

	}
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return router
}
