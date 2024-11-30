package main

import (
	"fmt"
	"golang-comment/database"
	"golang-comment/routes"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	r := gin.Default()
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error to load file ENV")
	}
	database.ConnectDatabase()

	routes.AuthRoutes(r)
	routes.CommentRoutes(r)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	r.Run(":" + port)
}
