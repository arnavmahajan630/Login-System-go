package auth

import (
	"log"
	"net/http"
	"strings"

	"github.com/arnavmahajan630/login-portal-go/services"
	"github.com/gin-gonic/gin"
)

func Authenticate() gin.HandlerFunc {

	return func (c * gin.Context){
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "You are unauthorized to access"})
			c.Abort()
			return
		}
		// Authorization token is sent as Bearer <..>
		// Removing Bearer

		authHeader = strings.TrimPrefix(authHeader, "Bearer")
		// check signature
		claims , err := services.Validate(authHeader)
		if err != nil {
			log.Println("The error is: ", err)
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
			c.Abort()
			return 
		}
		c.Set("claims", claims)
		c.Next() // we get validated and we are ready to move to next routes
	}
}