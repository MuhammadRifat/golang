package main

import (
	"log"
	"url-shortner/src/auth"
	"url-shortner/util"

	"github.com/gin-gonic/gin"
)

func main() {
	dsn := "root:rifat1234@tcp(127.0.0.1:3306)/url_shortner?charset=utf8mb4&parseTime=True&loc=Local"
	// Connect to the database using the utility function
	if err := util.ConnectDB(dsn); err != nil {
		log.Fatalf("Could not connect to the database: %v", err)
	}

	util.DB.AutoMigrate(&auth.User{})

	router := gin.Default()

	// Register auth routes
	auth.RegisterAuthRoutes(router)

	router.Run(":3000")
}
