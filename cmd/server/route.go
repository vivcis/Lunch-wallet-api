package server

import (
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"log"
	"strconv"
	"time"
)

// Injection function allow dependency injection into provided methods
type ginRouter struct {
	router *gin.Engine
}

// NewGinRouter creates an instance of the gin router
func NewGinRouter(r *gin.Engine) ports.Router {
	return &ginRouter{
		r,
	}
}

func (g *ginRouter) GET(uri string, fn func(ctx *gin.Context)) {
	g.router.GET(uri, fn)
}

func (g *ginRouter) POST(uri string, fn func(ctx *gin.Context)) {
	g.router.POST(uri, fn)
}

func (g *ginRouter) PUT(uri string, fn func(ctx *gin.Context)) {
	g.router.PUT(uri, fn)
}

func (g *ginRouter) DELETE(uri string, fn func(ctx *gin.Context)) {
	g.router.DELETE(uri, fn)
}

func (g ginRouter) GROUP(path string) *gin.RouterGroup {
	return g.router.Group(path)
}

func (g *ginRouter) SERVE() error {
	var port int
	if helpers.Instance.Port != nil {

		p, err := strconv.Atoi(*helpers.Instance.Port)
		if err != nil {
			log.Fatal("Error parsing port, must be numeric")
		}
		port = p
	}

	g.router.Use(gin.Logger())
	g.router.Use(gin.Recovery())
	g.router.Use(cors.New(cors.Config{
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
	if err := g.router.SetTrustedProxies([]string{"127.0.0.1"}); err != nil {
		return err
	}

	g.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	if err := g.router.Run(fmt.Sprintf("0.0.0.0:%d", port)); err != nil {
		fmt.Println("Could not run infrastructure -> ", err)
		return err
	}

	return nil
}
