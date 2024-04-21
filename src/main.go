package main

import (
	"fmt"
	"itw-be/src/config"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

var (
	server *gin.Engine
	err    error
	svPort string
)

func init() {
	err = godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	svPort = os.Getenv("PORT")
}

func main() {
	server = gin.Default()

	if err != nil {
		log.Fatal("Failed to connect to database. \n", err)
	}

	config.InitDB()

	log.Fatal(server.Run(fmt.Sprintf(":%s", svPort)))
}