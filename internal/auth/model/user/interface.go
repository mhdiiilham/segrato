package user

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Repository interface {
	Create(context.Context, User) (user User, err error)
	FindOne(ctx context.Context, username string) (user User, err error)
	FindByID(ctx context.Context, id string) (user User, err error)
	CheckUniqueness(ctx context.Context, username, email string) (unique bool)
}

type Service interface {
	RegisterUser(ctx context.Context, username, email, plainPassword string) (user User, accessToken string, err error)
	GetUser(ctx context.Context, id string) (user User, err error)
	Login(ctx context.Context, username, password string) (user User, accessToken string, err error)
	GetUserByAccessToken(ctx context.Context, accessToken string) (u User, err error)
}

type MongoCollection interface {
	InsertOne(ctx context.Context, document interface{}, opts ...*options.InsertOneOptions) (*mongo.InsertOneResult, error)
	FindOne(ctx context.Context, filter interface{}, opts ...*options.FindOneOptions) *mongo.SingleResult
}
