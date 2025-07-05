package main

import (
	"log"
	"os"

	"github.com/arnavmahajan630/login-portal-go/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
   // loading env
   if err := godotenv.Load(); err != nil {
       log.Fatal("Env missing / not laoded")
   }

   middleware.SetJWTkey(os.Getenv("JWT_SECRET"))

   r := gin.Default()
   routes.SetupRoutes(r)

   // server staring
   r.Run(":", os.Getenv("PORT"))
   log.Println("Server is running on Port" + os.Getenv("PORT"))
}