package main

import (
	"github.com/decadevs/lunch-api/cmd/server"
	_ "github.com/decadevs/lunch-api/docs"
	"log"
)

// @title           Lunch Wallet Swagger API
// @version         1.0
// @description     This is a lunch wallet server.
// @termsOfService  http://swagger.io/terms/

// @contact.name   Lunch-wallet Team API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  info@lunchwallet.com

// @license.name  BSD
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /api/v1

// @securityDefinitions.basic  BasicAuth
// @in header
// @name Authorization

func main() {
	db, err := server.Run()
	if err != nil {
		log.Fatal(err)
		return
	}
	server.Injection(db)
}
