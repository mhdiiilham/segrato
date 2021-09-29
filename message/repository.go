package message

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/mgo.v2/bson"
)

type repository struct {
	collection *mongo.Collection
}

func NewRepository(collection *mongo.Collection) Repository {
	return repository{collection: collection}
}

func (r repository) Create(ctx context.Context, msg Message) (ID string, err error) {
	var result *mongo.InsertOneResult
	result, err = r.collection.InsertOne(ctx, msg)
	if err != nil {
		return
	}
	oid, ok := result.InsertedID.(primitive.ObjectID)
	if !ok {
		return
	}
	ID = oid.Hex()
	return
}

func (r repository) FindOne(ctx context.Context, id string) (msg Message, err error) {
	objID, _ := primitive.ObjectIDFromHex(id)
	if err = r.collection.FindOne(ctx, bson.M{"_id": objID}).Decode(&msg); err != nil {
		return
	}
	return
}

func (r repository) UpdateOne(ctx context.Context, msg Message) (err error) {
	objID, _ := primitive.ObjectIDFromHex(msg.ID.Hex())
	_, err = r.collection.UpdateOne(ctx, bson.M{"_id": objID}, bson.M{"$set": msg})
	if err != nil {
		return
	}
	return
}

func (r repository) FindByUserID(ctx context.Context, userID string) (messages []Message, err error) {
	var cursor *mongo.Cursor

	findOptions := options.Find()
	findOptions.SetSort(bson.M{"createdAt": -1})
	cursor, err = r.collection.Find(ctx, bson.M{"userId": userID}, findOptions)
	if err != nil {
		return
	}
	if err = cursor.All(ctx, &messages); err != nil {
		return
	}
	return
}
