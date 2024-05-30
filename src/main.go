package main

import (
	"fmt"
	"itw-be/src/config"
	routers "itw-be/src/modules"
	"log"
	"os"

	"github.com/joho/godotenv"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "itw-be/src/docs" // Import thư viện swagger docs
)

// @title Your Project API
// @version 1.0
// @description This is a sample server for your project.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /v1

var (
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
	if svPort == "" {
		svPort = "8080"
	}

	r := routers.Router()
	// Đăng ký đường dẫn swagger
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	config.InitDB()

	log.Fatal(r.Run(fmt.Sprintf(":%s", svPort)))
}
