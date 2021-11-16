package auth_test

import (
	"context"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/mhdiiilham/segrato/config"
	"github.com/mhdiiilham/segrato/internal/auth"
	"github.com/mhdiiilham/segrato/internal/auth/model/user/mock"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/mongo"
)

func TestServer_HealthCheck(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	t.Run("all is well", func(t *testing.T) {
		t.Parallel()

		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()

		userService := mock.NewMockService(ctrl)

		userService.
			EXPECT().
			PingMongoDB(ctx).
			Return(nil).
			Times(1)

		srv := auth.NewServer(config.Config{}, userService)

		resp, err := srv.HealthCheck(ctx, nil)
		assert.NoError(t, err)
		assert.True(t, resp.ServerRunning)
		assert.True(t, resp.MongoDBConnection)
	})

	t.Run("mongo db connection is compromise", func(t *testing.T) {
		t.Parallel()

		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()

		userService := mock.NewMockService(ctrl)

		userService.
			EXPECT().
			PingMongoDB(ctx).
			Return(mongo.ErrClientDisconnected).
			Times(1)

		srv := auth.NewServer(config.Config{}, userService)

		resp, err := srv.HealthCheck(ctx, nil)
		assert.NoError(t, err)
		assert.True(t, resp.ServerRunning)
		assert.False(t, resp.MongoDBConnection)
	})
}
