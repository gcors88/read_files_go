package main

import (
	"log"

	"github.com/gcors88/read_files_go/src/app"
	"github.com/gcors88/read_files_go/src/infra/database"
	"github.com/joho/godotenv"
)

func main() {
	loadEnvError := godotenv.Load()

	if loadEnvError != nil {
		log.Fatal("Error on load envs")
	}

	connectInDatabaseError := database.ConnectInDatabase()

	if connectInDatabaseError != nil {
		log.Fatal("Error on connect in db")
	}

	app.ReadFiles()
}
