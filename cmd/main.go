package main

import (
	"github.com/decadevs/lunch-api/cmd/server"
	"log"
)

func main() {
	err := server.Run()
	if err != nil {
		log.Fatal(err)
		return
	}
	server.Injection()
}
