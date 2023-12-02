package database

import (
	"fmt"
	"os"
	"strconv"

	"github.com/gcors88/read_files_go/src/infra/database/entities"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Database *gorm.DB

type DbConfig struct {
	User     string
	Host     string
	Port     int
	Password string
	DbName   string
}

func ConnectInDatabase() (err error) {
	configDb := createDbConfig()
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", configDb.Host, configDb.Port, configDb.User, configDb.Password, configDb.DbName)

	db, err := gorm.Open(openPostgresConnection(connStr), &gorm.Config{})

	if err != nil {
		return err
	}

	err = db.AutoMigrate(&entities.User{})
	if err != nil {
		return err
	}

	fmt.Println("Connection has be successfully")

	Database = db

	return nil
}

func GetConnection() *gorm.DB {
	return Database
}

func openPostgresConnection(connStr string) gorm.Dialector {
	return postgres.Open(connStr)
}

func createDbConfig() *DbConfig {
	port, _ := strconv.Atoi(os.Getenv("POSTGRES_PORT"))

	return &DbConfig{
		Port:     port,
		DbName:   os.Getenv("POSTGRES_DB"),
		User:     os.Getenv("POSTGRES_USER"),
		Host:     os.Getenv("POSTGRES_HOST"),
		Password: os.Getenv("POSTGRES_PASSWORD"),
	}
}
