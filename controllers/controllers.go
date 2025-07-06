package controllers

import (
	"context"
	"net/http"
	"time"

	"github.com/arnavmahajan630/login-portal-go/config"
	"github.com/arnavmahajan630/login-portal-go/models"
	"github.com/arnavmahajan630/login-portal-go/services"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)



var validate = validator.New()
var userCollection = config.Opencollection("users")


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
	   // duplicate sign ins
	   count , err := userCollection.CountDocuments(ctx, bson.M{
		 "$or": []bson.M{
			{"email": user.Email},
			{"phone": user.Phone},

		},
	   })
	   if err != nil {
		  c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	   }
	   
	   if count > 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Email or phone already exists"})
	   }
       
	   // hash password
	   user.Password = services.HashPassword(user.Password)
	   user.CreatedAt = time.Now()
	   user.UpdatedAt = time.Now()
	   user.ID = primitive.NewObjectID()
	   user.UserID = user.ID.Hex()

	   // insert in the db
	   if _, ierr := userCollection.InsertOne(ctx, user); ierr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": ierr.Error()})
	   }

	   c.JSON(http.StatusOK, gin.H{"message": "User Signed up successfully", "user_id": user.UserID})
       
	}
}

func Login() gin.HandlerFunc {
	return func( c * gin.Context) {
        ctx , cancel := context.WithTimeout(context.Background(), 10 * time.Second)
		defer cancel()
		var user models.User
		var foundUser models.User

		if err := c.BindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return;
		}
		err := userCollection.FindOne(ctx, bson.M{"email": user.Email}).Decode(&foundUser)
		if err != nil {
			 c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
		passowrdIsValid, msg := services.VerifyPass(*foundUser.Password, *user.Password)
		if !passowrdIsValid {
			 c.JSON(http.StatusUnauthorized, gin.H{"error": msg})
		}

	   accessToken , refreshToken := services.GenerateTokens(*user.Email, user.UserID, *user.Role);
	   user.AccessToken = accessToken
	   user.RefreshToken = refreshToken
	   services.UpdateAllTokens(accessToken, refreshToken, foundUser.UserID)
	   c.JSON(http.StatusOK, gin.H{
		   "user": foundUser,
		   "access_token": accessToken,
		    "refresh_token": refreshToken,
	   })
	}
}

func GetUsers() gin.HandlerFunc {
	return func (c * gin.Context) {
		claims, exist := c.Get("claims")
		if !exist {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "User Unauthorized"})
			return
		}
		tokenClaims, ok := claims.(*services.Claims)
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "User Unauthorized"})
			return
		}
		if tokenClaims.Role !=  "admin" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "User Unauthorized"})
			return
		}
		// admin is authorized. returning list of users

		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		cursor, err := userCollection.Find(ctx, bson.M{})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		defer cursor.Close(ctx)
		var users []models.User
		if err := cursor.All(ctx ,&users); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		c.JSON(http.StatusOK, users)

	}
}

func GetUser() gin.HandlerFunc {
	return func (c * gin.Context) {
		requestid := c.Param("id")
		
		claims, exists := c.Get("claims")
		if !exists {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "User Unauthorized"})
		}
		tokenClaims, ok := claims.(*services.Claims)
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "User Unauthorized"})
			return
		}

		if(tokenClaims.ID != requestid) {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "User Unauthorized"})
			return
		}

		if(tokenClaims.Role != "admin") {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "User Unauthorized"})
			return
		}

		// querry db to get the single user
		var user models.User
		ctx , cancel := context.WithTimeout(context.Background(), 10* time.Second)
		defer cancel()
		if err := userCollection.FindOne(ctx, bson.M{"user_id": requestid}).Decode(&user); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, user)

	}
}


func HelloWorld() gin.HandlerFunc {
	return func (c * gin.Context) {
		c.JSON(http.StatusOK, "Hello Welcome to Authentication in Golang")
	}
}
