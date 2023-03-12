package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID              primitive.ObjectID `bson:"_id"`
	First_name      *string            `json:"first_name" validate:"required,min=2,max=100"`
	Last_name       *string            `json:"last_name" validate:"required,min=2,max=100"`
	Gender          *string            `json:"gender" validate:"require"`
	Address         *string            `json:"address" validate:"require"`
	City            *string            `json:"city" validate:"require"`
	Phone           *string            `json:"phone" validate:"require min=10,max=14"`
	Password        *string            `json:"Password" validate:"required,min=6"`
	Email           *string            `json:"email" validate:"email,required"`
	Disease         *string
	Patient_history *string
	Token           *string   `json:"token"`
	User_type       *string   `json:"user_type" validate:"required,eq=ADMIN|eq=USER"`
	Refresh_token   *string   `json:"refresh_token"`
	Created_at      time.Time `json:"created_at"`
	Updated_at      time.Time `json:"updated_at"`
	User_id         string    `json:"user_id"`
}
