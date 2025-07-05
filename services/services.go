package services

import (
	"fmt"

	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	UserID string `json:"user_id"`
	Email string `json:"email"`
	Role string `json:"role"`
	jwt.RegisteredClaims
}

var jwtKey []byte
func Setjwtkey() {
	jwtKey = []byte
}

func GetjwtKey() []byte{
	return []byte(jwtKey)
}

func ParseToken(tokenString string) (*Claims, error) {
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


