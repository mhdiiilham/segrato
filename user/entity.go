package user

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID               primitive.ObjectID `bson:"_id,omitempty"`
	Username         string             `bson:"username"`
	Password         string             `bson:"password"`
	IsPremium        bool               `bson:"isPremium"`
	PremiumExpiredAt time.Time          `bson:"premiumExpiredAt"`
}
