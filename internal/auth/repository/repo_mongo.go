package repository

import (
	"context"

	"github.com/mhdiiilham/segrato/internal/auth/model/user"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"gopkg.in/mgo.v2/bson"
)

type repository struct {
	collection *mongo.Collection
}

func NewUserRepository(collection *mongo.Collection) user.Repository {
	return &repository{
		collection: collection,
	}
}

func (r repository) Create(ctx context.Context, u user.User) (user.User, error) {
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

func (r repository) FindOne(ctx context.Context, username string) (user user.User, err error) {
	err = r.collection.FindOne(ctx, bson.M{"username": username}).Decode(&user)
	return
}

func (r repository) CheckUniqueness(ctx context.Context, username, email string) (unique bool) {
	unique = false
	err := r.collection.FindOne(ctx, bson.M{"$or": []bson.M{
		bson.M{"username": username},
		bson.M{"email": email},
	}}).Err()
	return err == mongo.ErrNoDocuments
}

func (r repository) FindByID(ctx context.Context, id string) (user user.User, err error) {
	objID, _ := primitive.ObjectIDFromHex(id)
	err = r.collection.FindOne(ctx, bson.M{"_id": objID}).Decode(&user)
	if err != nil {
		return
	}

	return
}

func (r repository) PingMongoDB(ctx context.Context) error {
	return r.collection.Database().Client().Ping(ctx, nil)
}
