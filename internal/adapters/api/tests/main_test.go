package tests

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	if err := godotenv.Load("../../../../.env"); err != nil {
		log.Println(err.Error())
	}
	exitcode := m.Run()

	os.Exit(exitcode)
}
