package main

import (
	"github.com/decadevs/lunch-api/cmd/server"
	"log"
)

// @title Swagger  lunch wallet service API
// @version 1.0
// @description This is Lunch wallet APIs.
// @termsOfService demo.com

// @contact.name API Support
// @contact.url https://deca-meal-wallet.herokuapp.com

// @host localhost:8081
// @BasePath /api/v1

// @securityDefinitions.basic BasicAuth

// @securityDefinitions.apikey ApiKeyAuth
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
