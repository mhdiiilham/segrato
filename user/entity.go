package user

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID               primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Username         string             `bson:"username" json:"username"`
	Password         string             `bson:"password" json:"-"`
	IsPremium        bool               `bson:"isPremium" json:"isPremium"`
	BlockedWords     []string           `bson:"blockedWords" json:"blockedWords"`
	PremiumExpiredAt time.Time          `bson:"premiumExpiredAt" json:"premiumExpiredAt"`
}

type RegisterUserResponse struct {
	ID          string `json:"id"`
	AccessToken string `json:"accessToken"`
	User        User   `json:"user"`
}

type RegisterUserPayload struct {
	Username     string   `json:"username"`
	Password     string   `json:"password"`
	BlockedWords []string `json:"BlockedWords"`
}
