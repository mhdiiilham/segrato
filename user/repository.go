package user

import (
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

type Repository interface {
	Create(ctx context.Context, username, plainPassword string) (user User, err error)
	FindOne(ctx context.Context, username string) (user User, err error)
}

type repository struct {
	collection *mongo.Collection
}

func NewRepository(collection *mongo.Collection) Repository {
	return &repository{
		collection: collection,
	}
}

func (r repository) Create(ctx context.Context, username, plainPassword string) (user User, err error) {
	var insertResult *mongo.InsertOneResult
	var bytePassword []byte
	bytePassword, err = bcrypt.GenerateFromPassword([]byte(plainPassword), bcrypt.MinCost)
	if err != nil {
		return
	}

	if !r.checkUniqueness(ctx, username) {
		err = errors.New("username already taken")
		return
	}

	password := string(bytePassword)
	user.Username = username
	user.Password = password
	insertResult, err = r.collection.InsertOne(ctx, user)
	if err != nil {
		return
	}

	oid, ok := insertResult.InsertedID.(primitive.ObjectID)
	if !ok {
		return
	}

	user.ID = oid
	return
}

func (r repository) FindOne(ctx context.Context, username string) (user User, err error) {
	err = r.collection.FindOne(ctx, bson.M{"username": username}).Decode(&user)
	return
}

func (r repository) checkUniqueness(ctx context.Context, username string) (unique bool) {
	unique = false
	err := r.collection.FindOne(ctx, bson.M{"username": username}).Err()
	return err == mongo.ErrNoDocuments
}
