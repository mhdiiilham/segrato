package user

import (
	"errors"
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

var (
	INVALID_USERNAME_PASSWORD error = errors.New("username or password is wrong")
)
