package main

import (
	"log"
	"os"

	"github.com/arnavmahajan630/login-portal-go/routes"
	"github.com/arnavmahajan630/login-portal-go/services"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// loading env
	if err := godotenv.Load(); err != nil {
		panic(err)
	}
	// setjwt
	services.Setjwtkey(os.Getenv("JWT_SECRET"))

	// initialize routes
	r := gin.Default()
	routes.SetupRoutes(r)

	// server staring
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000" // fallback default
	}
	r.Run(":" + port)
	log.Println("Server is running on Port" + os.Getenv("PORT"))
}
