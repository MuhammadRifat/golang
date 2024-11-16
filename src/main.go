package main

import (
	"log"
	"os"
	"url-shortner/src/auth"
	"url-shortner/src/url"
	"url-shortner/util"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// Access environment variables
	port := os.Getenv("PORT")
	dbURL := os.Getenv("DATABASE_URL")

	// Connect to the database using the utility function
	if err := util.ConnectDB(dbURL); err != nil {
		log.Fatalf("Could not connect to the database: %v", err)
	}

	util.DB.AutoMigrate(&auth.User{}, &url.URL{})

	router := gin.Default()

	// Register routes
	auth.RegisterAuthRoutes(router)
	url.RegisterUrlRoutes(router)

	router.Run(":" + port)
}
