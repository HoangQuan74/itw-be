package config

import (
    "fmt"
    "log"
    "os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB     *gorm.DB
	err    error
	dbHost string
	dbPort string
	dbUser string
	dbPass string
	dbName string
	dbType string
)

func init() {
    err = godotenv.Load()
    if err != nil {
        log.Fatal("Error loading .env file")
    }

    dbHost = os.Getenv("DB_HOST")
    dbPort = os.Getenv("DB_PORT")
    dbUser = os.Getenv("DB_USERNAME")
    dbPass = os.Getenv("DB_PASSWORD")
    dbName = os.Getenv("DB_NAME")
}

func InitDB() {
    dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbUser, dbPass, dbHost, dbPort, dbName)
    DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatal("Failed to connect to database. \n", err)
    }
}
