package main

import (
	"dayalubajpai/pkg/routes"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {

	//port setup to run server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8089"
	}

	// Initialize the server
	router := gin.New()

	// Logger
	router.Use(gin.Logger())

	routes.RegisterBookRoutes(router)
	router.Use(gin.Recovery())

	router.Run(":" + port)
	fmt.Println("Server is running on port: " + port)

}
