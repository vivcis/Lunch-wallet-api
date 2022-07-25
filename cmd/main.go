package main

import (
	"github.com/decadevs/lunch-api/cmd/server"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {
	env := os.Getenv("GIN_MODE")
	if env != "release" {
		if err := godotenv.Load(); err != nil {
			log.Fatalf("couldn't load env vars: %v", err)
		}
	}
	db, err := server.Run()
	if err != nil {
		log.Fatal(err)
		return
	}
	server.Injection(db)
}
