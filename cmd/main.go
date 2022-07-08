package main

import "log"

func main() {
	err := server.Run()
	if err != nil {
		log.Fatal(err)
		return
	}
	server.Injection()
}
