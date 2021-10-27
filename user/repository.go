package user

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type repository struct {
	collection MongoCollection
}

func NewRepository(collection MongoCollection) Repository {
	return &repository{
		collection: collection,
	}
}

func (r repository) Create(ctx context.Context, u User) (User, error) {
	var insertResult *mongo.InsertOneResult

	insertResult, err := r.collection.InsertOne(ctx, &u)
	if err != nil {
		return u, err
	}

	oid, ok := insertResult.InsertedID.(primitive.ObjectID)
	if !ok {
		return u, err
	}

	u.ID = oid
	return u, nil
}

func (r repository) FindOne(ctx context.Context, username string) (user User, err error) {
	err = r.collection.FindOne(ctx, bson.M{"username": username}).Decode(&user)
	return
}

func (r repository) CheckUniqueness(ctx context.Context, username string) (unique bool) {
	unique = false
	err := r.collection.FindOne(ctx, bson.M{"username": username}).Err()
	return err == mongo.ErrNoDocuments
}

func (r repository) FindByID(ctx context.Context, id string) (user User, err error) {
	objID, _ := primitive.ObjectIDFromHex(id)
	err = r.collection.FindOne(ctx, bson.M{"_id": objID}).Decode(&user)
	if err != nil {
		return
	}

	return
}
