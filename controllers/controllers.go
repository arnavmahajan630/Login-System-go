package controllers

import (
	"context"
	"net/http"
	"time"

	"github.com/arnavmahajan630/login-portal-go/models"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

func Signup() gin.HandlerFunc{
	return func(c * gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		var user models.User
        

		// get inputs
		if err := c.BindJSON(&user); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		// validate input
		if validationErr := validate.Struct(user); validationErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error()})
		return
	   }
	   


	}
}