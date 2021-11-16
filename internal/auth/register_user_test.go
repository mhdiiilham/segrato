package auth_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/bxcodec/faker/v3"
	"github.com/golang/mock/gomock"
	"github.com/mhdiiilham/segrato/config"
	"github.com/mhdiiilham/segrato/internal/auth"
	"github.com/mhdiiilham/segrato/internal/auth/model/user"
	"github.com/mhdiiilham/segrato/internal/auth/model/user/mock"
	"github.com/mhdiiilham/segrato/internal/proto"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestServer_RegisterUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	t.Run("payload is empty", func(t *testing.T) {
		var (
			username    string
			password    string
			email       string
			expectedErr error = status.Error(codes.InvalidArgument, "invalid email or/and password")
		)
		payload := &proto.RegisterUserRequest{
			Username: username,
			Password: password,
			Email:    email,
		}
		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()

		userService := mock.NewMockService(ctrl)

		authSrv := auth.NewServer(config.Config{}, userService)
		resp, err := authSrv.RegisterUser(ctx, payload)
		assert.Nil(t, resp, "when request is invalid, resp should be nil")
		assert.NotNil(t, err, "when reqiest is invalid, err should be not nil")
		assert.ErrorIs(t, err, expectedErr)
	})

	t.Run("username/email already registered", func(t *testing.T) {
		var (
			username    string = faker.Username()
			password    string = faker.Password()
			email       string = faker.Email()
			expectedErr error  = user.ErrUsernameEmailRegistered
		)
		payload := &proto.RegisterUserRequest{
			Username: username,
			Password: password,
			Email:    email,
		}
		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()

		userService := mock.NewMockService(ctrl)
		userService.
			EXPECT().
			RegisterUser(ctx, payload.Username, payload.Email, payload.Password).
			Return(user.User{}, "", expectedErr).
			Times(1)

		authSrv := auth.NewServer(config.Config{}, userService)
		resp, err := authSrv.RegisterUser(ctx, payload)
		assert.Nil(t, resp, "when request is invalid, resp should be nil")
		assert.NotNil(t, err, "when reqiest is invalid, err should be not nil")
		assert.ErrorIs(t, err, status.Errorf(codes.AlreadyExists, expectedErr.Error()))
	})

	t.Run("internal server error", func(t *testing.T) {
		var (
			username    string = faker.Username()
			password    string = faker.Password()
			email       string = faker.Email()
			expectedErr error  = status.Error(codes.Internal, "internal server error")
		)
		payload := &proto.RegisterUserRequest{
			Username: username,
			Password: password,
			Email:    email,
		}
		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()

		userService := mock.NewMockService(ctrl)
		userService.
			EXPECT().
			RegisterUser(ctx, payload.Username, payload.Email, payload.Password).
			Return(user.User{}, "", errors.New("something happened")).
			Times(1)

		authSrv := auth.NewServer(config.Config{}, userService)
		resp, err := authSrv.RegisterUser(ctx, payload)
		assert.Nil(t, resp, "when request is invalid, resp should be nil")
		assert.NotNil(t, err, "when reqiest is invalid, err should be not nil")
		assert.ErrorIs(t, err, expectedErr)
	})

	t.Run("register success", func(t *testing.T) {
		var (
			jwt      string = faker.Jwt()
			username string = faker.Username()
			password string = faker.Password()
			email    string = faker.Email()
		)
		payload := &proto.RegisterUserRequest{
			Username: username,
			Password: password,
			Email:    email,
		}
		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()

		userService := mock.NewMockService(ctrl)
		userService.
			EXPECT().
			RegisterUser(ctx, payload.Username, payload.Email, payload.Password).
			Return(user.User{
				Username: username,
				Email:    email,
				Password: password,
			}, jwt, nil).
			Times(1)

		authSrv := auth.NewServer(config.Config{}, userService)
		resp, err := authSrv.RegisterUser(ctx, payload)
		assert.NoError(t, err)
		assert.Equal(t, jwt, resp.AccessToken)
	})
}
