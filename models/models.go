package models

import "time"

type User struct {
	ID        int        `bson:"_id,omitempty"`
	Fname     *string    `json:"fname" bson:"fname" validate:"required,min=2,max=100"`
	Lname     *string    `json:"lname" bson:"lname" validate:"required,min=2,max=100"`
	Password  *string    `json:"password" bson:"password" validate:"required,min=6,max=100"`
	Email     *string    `json:"email" bson:"email" validate:"required,email"`
	Phone     *string    `json:"phone" bson:"phone" validate:"required,e164"` // or "required,len=10,numeric" for Indian numbers
	Role      *string    `json:"role" bson:"role" validate:"required,oneof=admin user guest"`
	CreatedAt time.Time  `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time  `json:"updated_at" bson:"updated_at"`
}

