package auth_test

import (
	"context"
	"testing"

	"github.com/bxcodec/faker/v3"
	"github.com/golang/mock/gomock"
	"github.com/mhdiiilham/segrato/config"
	"github.com/mhdiiilham/segrato/internal/auth"
	"github.com/mhdiiilham/segrato/internal/auth/model/proto"
	"github.com/mhdiiilham/segrato/internal/auth/model/user"
	"github.com/mhdiiilham/segrato/internal/auth/model/user/mock"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestServer_GetUserByAccessToken(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	t.Run("access token is empty", func(t *testing.T) {
		t.Parallel()

		expectedErr := status.Errorf(codes.InvalidArgument, "accessToken couldn't be empty")
		ctx := context.Background()
		userService := mock.NewMockService(ctrl)
		jwt := ""

		authSrv := auth.NewServer(config.Config{}, userService)
		resp, err := authSrv.GetUserByAccessToken(ctx, &proto.GetUserByAccessTokenRequest{
			AccessToken: jwt,
		})
		assert.Nil(t, resp)
		assert.NotNil(t, err)
		assert.ErrorIs(t, err, expectedErr)
	})

	t.Run("access is invalid", func(t *testing.T) {
		t.Parallel()

		expectedErr := status.Errorf(codes.InvalidArgument, "invalid access token")
		ctx := context.Background()
		userService := mock.NewMockService(ctrl)
		jwt := faker.Jwt()

		userService.
			EXPECT().
			GetUserByAccessToken(ctx, jwt).
			Return(user.User{}, user.ErrAccessTokenInvalid).
			Times(1)

		authSrv := auth.NewServer(config.Config{}, userService)
		resp, err := authSrv.GetUserByAccessToken(ctx, &proto.GetUserByAccessTokenRequest{
			AccessToken: jwt,
		})
		assert.Nil(t, resp)
		assert.NotNil(t, err)
		assert.ErrorIs(t, err, expectedErr)
	})

	t.Run("user not found", func(t *testing.T) {
		t.Parallel()

		expectedErr := status.Errorf(codes.InvalidArgument, "invalid access token")
		ctx := context.Background()
		userService := mock.NewMockService(ctrl)
		jwt := faker.Jwt()

		userService.
			EXPECT().
			GetUserByAccessToken(ctx, jwt).
			Return(user.User{}, mongo.ErrNoDocuments).
			Times(1)

		authSrv := auth.NewServer(config.Config{}, userService)
		resp, err := authSrv.GetUserByAccessToken(ctx, &proto.GetUserByAccessTokenRequest{
			AccessToken: jwt,
		})
		assert.Nil(t, resp)
		assert.NotNil(t, err)
		assert.ErrorIs(t, err, expectedErr)
	})

	t.Run("unexptected error", func(t *testing.T) {
		t.Parallel()

		expectedErr := status.Error(codes.Internal, "internal server error")
		ctx := context.Background()
		userService := mock.NewMockService(ctrl)
		jwt := faker.Jwt()

		userService.
			EXPECT().
			GetUserByAccessToken(ctx, jwt).
			Return(user.User{}, mongo.ErrNilCursor).
			Times(1)

		authSrv := auth.NewServer(config.Config{}, userService)
		resp, err := authSrv.GetUserByAccessToken(ctx, &proto.GetUserByAccessTokenRequest{
			AccessToken: jwt,
		})
		assert.Nil(t, resp)
		assert.NotNil(t, err)
		assert.ErrorIs(t, err, expectedErr)
	})

	t.Run("success get user", func(t *testing.T) {
		t.Parallel()

		ctx := context.Background()
		userService := mock.NewMockService(ctrl)
		jwt := faker.Jwt()
		id := primitive.NewObjectID()
		username := faker.Username()
		email := faker.Email()
		password := faker.Password()

		userService.
			EXPECT().
			GetUserByAccessToken(ctx, jwt).
			Return(user.User{
				ID:       id,
				Username: username,
				Email:    email,
				Password: password,
			}, nil).
			Times(1)

		authSrv := auth.NewServer(config.Config{}, userService)
		resp, err := authSrv.GetUserByAccessToken(ctx, &proto.GetUserByAccessTokenRequest{
			AccessToken: jwt,
		})
		assert.Nil(t, err)
		assert.NotNil(t, resp)
		assert.Equal(t, username, resp.User.GetUsername())
		assert.Equal(t, email, resp.User.GetEmail())
		assert.Equal(t, id.Hex(), resp.User.GetId())
	})
}
