package routes

import (
	"github.com/arnavmahajan630/login-portal-go/controllers"
	auth "github.com/arnavmahajan630/login-portal-go/middleware"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(r * gin.Engine) {
    r.POST("/signup", controllers.Signup())
	r.POST("/login", controllers.Login())


    protected := r.Group("/") 
	protected.Use(auth.Authenticate())
	{	protected.GET("/users", controllers.GetUsers())
		protected.GET("/users/:id", controllers.GetUser())
    }

}