package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID            primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	First_name    *string            `json:"first_name" validate:"required, min=2, max=100" bson:"First_name"`
	Last_name     *string            `json:"last_name" validate:"required, min=2, max=100" bson:"Last_name"`
	Password      *string            `json:"password" validate:"required, min=8, max=100" bson:"Password"`
	Email         *string            `json:"email" validate:"required, email" bson:"Email"`
	Phone         *string            `json:"phone" validate:"required, min=10, max=10" bson:"Phone"`
	Token         *string            `json:"token" bson:"Token"`
	User_type     *string            `json:"user_type" bson:"User_type"`
	Refresh_token *string            `json:"refresh_token" bson:"Refresh_token"`
	Created_at    *time.Time         `json:"created_at" bson:"Created_at"`
	Updated_at    *time.Time         `json:"updated_at" bson:"Updated_at"`
	User_id       *string            `json:"user_id" bson:"User_id"`
}
