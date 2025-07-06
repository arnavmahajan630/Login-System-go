package services

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/arnavmahajan630/login-portal-go/config"
	"github.com/golang-jwt/jwt/v5"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
)

type Claims struct {
	UserID string `json:"user_id"`
	Email  string `json:"email"`
	Role   string `json:"role"`
	jwt.RegisteredClaims
}

var jwtKey []byte

func Setjwtkey(key string) {
	jwtKey = []byte(key)
}

func GetjwtKey() []byte {
	return []byte(jwtKey)
}

func Validate(tokenString string) (*Claims, error) {
	claims := &Claims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		// Optional: check signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method")
		}
		return jwtKey, nil
	})

	if err != nil || !token.Valid {
		return nil, fmt.Errorf("invalid token: %v", err)
	}

	return claims, nil
}

func HashPassword(password *string) *string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(*password), bcrypt.DefaultCost)

	if err != nil {
		panic(err)
	}
	haspwd := string(bytes)
	return &haspwd
}

func GenerateTokens(email, userID, role string) (string, string) {
	accessExpiry, refreshExpiry := getTokenExpiryDuration()

	claims := &Claims{
		Email:  email,
		UserID: userID,
		Role:   role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(accessExpiry)),
		},
	}
	refreshClaims := &Claims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(refreshExpiry)),
		},
	}

	accessToken := jwt.NewWithClaims(jwt.SigningMethodES256, claims)
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodES256, refreshClaims)
	signedAccess, err := accessToken.SignedString(jwtKey)
	if err != nil {
		panic(err)
	}
	signedRefresh, err := refreshToken.SignedString(jwtKey)
	if err != nil {
		panic(err)
	}

	return signedAccess, signedRefresh
}

func getTokenExpiryDuration() (time.Duration, time.Duration) {
	accessExpiry := os.Getenv("JWT_EXPIRY")
	refreshExpiry := os.Getenv("REFRESH_EXPIRY")
	if accessExpiry == "" {
		log.Println("⚠️ TOKEN_EXPIRY not set, defaulting to 24h")
		accessExpiry = "24h"
	}
	if refreshExpiry == "" {
		log.Println("⚠️ Refresh Expiry not set, defaulting to 24h")
		refreshExpiry = "120h"
	}
	accessDuration, err := time.ParseDuration(accessExpiry)
	if err != nil {
		log.Fatalf("❌ Invalid ACCESS_EXPIRY format: %v", err)
	}
	refreshDuration, err := time.ParseDuration(refreshExpiry)

	if err != nil {
		log.Fatalf("❌ Invalid REFRESH_EXPIRY format: %v", err)
	}
	return accessDuration, refreshDuration

}

func VerifyPass(foundpass, userpass string) (bool, string) {
     if err := bcrypt.CompareHashAndPassword([]byte(foundpass), []byte(userpass)); err != nil {
		return false, err.Error()
	 } else  {
		return true, ""
	 }
}

func UpdateAllTokens(access, refresh, userID string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	usercollection := config.Opencollection("users")

	updateObj := bson.M{
		"$set": bson.M{
			"access_token":  access,
			"refresh_token": refresh,
			"updated_at":    time.Now(),
		},
	}

	filter := bson.M{"user_id": userID}
	_, err := usercollection.UpdateOne(ctx, filter, updateObj)
	return err
}
