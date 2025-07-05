package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/gofiber/fiber/middleware"
)

func SetupRoutes(r * gin.Engine) {
    r.POST("/signup", controllers.Signup())
	r.POST("/login", controllers.Login())


    protected := r.Group("/") 
	protected.Use(middleware.Authenticate())
	{	protected.GET("/users", controllers.GetUsers())
		protected.GET("/users/:id", controllers.GetUser())
    }

}