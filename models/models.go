package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID           primitive.ObjectID `bson:"_id,omitempty"`
	UserID       string                `json:"user_id" bson:"user_id"` 
	Fname        *string            `json:"fname" bson:"fname" validate:"required,min=2,max=100"`
	Lname        *string            `json:"lname" bson:"lname" validate:"required,min=2,max=100"`
	Password     *string            `json:"password" bson:"password" validate:"required,min=6,max=100"`
	Email        *string            `json:"email" bson:"email" validate:"required,email"`
	Phone        *string            `json:"phone" bson:"phone" validate:"required,e164"`
	Role         *string            `json:"role" bson:"role" validate:"required,oneof=admin user guest"`
	AccessToken  string             `json:"access_token,omitempty" bson:"access_token,omitempty"`   
	RefreshToken string             `json:"refresh_token,omitempty" bson:"refresh_token,omitempty"`
	CreatedAt    time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt    time.Time          `json:"updated_at" bson:"updated_at"`
}

