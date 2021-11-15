package user

import (
	"errors"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Username string             `bson:"username" json:"username"`
	Email    string             `bson:"email" json:"email"`
	Password string             `bson:"password" json:"-"`
}

var (
	ErrInvalidUsernamePassword error = errors.New("username or password is wrong")
	ErrUsernameEmailRegistered error = errors.New("username or email already registered")
	ErrAccessTokenInvalid      error = errors.New("access token not valid")
)
