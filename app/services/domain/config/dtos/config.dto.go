package dtos

import "go.mongodb.org/mongo-driver/bson/primitive"

// Contact Request DTO with Validation
type ContactReqDto struct {
	ID    primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name  string             `json:"name" binding:"required"` // Required field
	Email string             `json:"email" binding:"required,email"` // Required & must be a valid email
	Msg   string             `json:"msg" binding:"required,min=10"` // Required & min length of 10 characters
}