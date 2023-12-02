package main

import (
	"log"

	"github.com/gcors88/read_files_go/src/infra/database"
	"github.com/joho/godotenv"
)

func main() {
	loadEnvError := godotenv.Load()

	if loadEnvError != nil {
		log.Fatal("Error on load envs")
		return
	}

	connectInDatabaseError := database.ConnectInDatabase()

	if connectInDatabaseError != nil {
		log.Fatal("Error on connect in db")
	}

	connection := database.GetConnection()

	_, queryError := connection.Exec(getQueryToCreateDatabase()).Rows()

	if queryError != nil {
		log.Fatal("Error on create db ", queryError)
	}
}

func getQueryToCreateDatabase() string {
	return `CREATE DATABASE read_files`
}
