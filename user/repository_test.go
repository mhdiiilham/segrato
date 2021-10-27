package user_test

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	mock "github.com/mhdiiilham/segrato/mock/user"
	"github.com/mhdiiilham/segrato/user"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func Test_repository_Create(t *testing.T) {
	ctx := context.Background()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	insertUser := user.User{}

	t.Run("create user success", func(t *testing.T) {
		mockMongoCollection := mock.NewMockMongoCollection(ctrl)

		mockMongoCollection.
			EXPECT().
			InsertOne(ctx, &insertUser).
			Return(&mongo.InsertOneResult{
				InsertedID: primitive.NewObjectID(),
			}, nil).
			Times(1)

		r := user.NewRepository(mockMongoCollection)
		_, err := r.Create(ctx, insertUser)
		assert.NoError(t, err)
	})

	t.Run("create user failed", func(t *testing.T) {
		mockMongoCollection := mock.NewMockMongoCollection(ctrl)

		mockMongoCollection.
			EXPECT().
			InsertOne(ctx, &insertUser).
			Return(nil, mongo.ErrNilDocument).
			Times(1)

		r := user.NewRepository(mockMongoCollection)
		_, err := r.Create(ctx, insertUser)
		assert.NotNil(t, err)
	})

	t.Run("failed to cast object", func(t *testing.T) {
		mockMongoCollection := mock.NewMockMongoCollection(ctrl)

		mockMongoCollection.
			EXPECT().
			InsertOne(ctx, &insertUser).
			Return(&mongo.InsertOneResult{
				InsertedID: "",
			}, nil).
			Times(1)

		r := user.NewRepository(mockMongoCollection)
		_, err := r.Create(ctx, insertUser)
		assert.NotNil(t, err)
	})

}
