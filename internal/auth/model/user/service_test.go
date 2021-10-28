package user_test

import (
	"context"
	"errors"
	"fmt"
	"testing"

	"github.com/bxcodec/faker/v3"
	"github.com/golang/mock/gomock"
	"github.com/mhdiiilham/segrato/internal/auth/model/user"
	mockUser "github.com/mhdiiilham/segrato/internal/auth/model/user/mock"
	mockPassword "github.com/mhdiiilham/segrato/pkg/password/mock"
	"github.com/mhdiiilham/segrato/pkg/token"
	mockToken "github.com/mhdiiilham/segrato/pkg/token/mock"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func Test_service_RegisterUser(t *testing.T) {
	ctx := context.Background()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	t.Run("register user success", func(t *testing.T) {
		userRepositoryMock := mockUser.NewMockRepository(ctrl)
		tokenServiceMock := mockToken.NewMockService(ctrl)
		passwordServiceMock := mockPassword.NewMockService(ctrl)

		username := faker.Username()
		password := faker.Password()
		email := faker.Email()
		jwt := faker.Jwt()
		newUser := user.User{
			Email:    email,
			Username: username,
			Password: password,
		}

		userRepositoryMock.
			EXPECT().
			CheckUniqueness(ctx, username, email).
			Return(true).
			Times(1)

		passwordServiceMock.
			EXPECT().
			HashPassword(newUser.Password).
			Return(newUser.Password, nil).
			Times(1)

		userRepositoryMock.
			EXPECT().
			Create(ctx, newUser).
			Return(newUser, nil).
			Times(1)

		tokenServiceMock.
			EXPECT().
			SignPayload(token.TokenPayload{
				ID:       newUser.ID.Hex(),
				Username: newUser.Username,
			}).
			Return(jwt, nil).
			Times(1)

		service := user.NewService(userRepositoryMock, tokenServiceMock, passwordServiceMock)
		_, _, err := service.RegisterUser(ctx, newUser.Username, newUser.Email, newUser.Password)
		assert.NoError(t, err)
	})

	t.Run("register user failed - email/username duplicate", func(t *testing.T) {
		userRepositoryMock := mockUser.NewMockRepository(ctrl)
		tokenServiceMock := mockToken.NewMockService(ctrl)
		passwordServiceMock := mockPassword.NewMockService(ctrl)

		email := faker.Email()
		username := faker.Username()
		passwd := faker.Password()

		userRepositoryMock.
			EXPECT().
			CheckUniqueness(ctx, username, email).
			Return(false).
			Times(1)

		expectedErr := user.ErrUsernameEmailRegistered
		service := user.NewService(userRepositoryMock, tokenServiceMock, passwordServiceMock)
		_, _, err := service.RegisterUser(ctx, username, email, passwd)
		assert.True(t, errors.Is(err, expectedErr), "expecting error to be: '%s' but got: '%s' instead", expectedErr.Error(), err.Error())
	})

	t.Run("register user failed - hashing password failed", func(t *testing.T) {
		userRepositoryMock := mockUser.NewMockRepository(ctrl)
		tokenServiceMock := mockToken.NewMockService(ctrl)
		passwordServiceMock := mockPassword.NewMockService(ctrl)

		email := faker.Email()
		username := faker.Username()
		passwd := faker.Password()
		expectedErr := errors.New("hashing password failed")

		userRepositoryMock.
			EXPECT().
			CheckUniqueness(context.Background(), username, email).
			Return(true).
			Times(1)

		passwordServiceMock.
			EXPECT().
			HashPassword(passwd).
			Return(gomock.Any().String(), expectedErr).
			Times(1)

		service := user.NewService(userRepositoryMock, tokenServiceMock, passwordServiceMock)
		_, _, err := service.RegisterUser(ctx, username, email, passwd)
		assert.True(t, errors.Is(err, expectedErr), "expecting error to be: '%s' but got: '%s' instead", expectedErr.Error(), err.Error())
	})

	t.Run("register user failed - repository failed creating document", func(t *testing.T) {
		userRepositoryMock := mockUser.NewMockRepository(ctrl)
		tokenServiceMock := mockToken.NewMockService(ctrl)
		passwordServiceMock := mockPassword.NewMockService(ctrl)

		email := faker.Email()
		username := faker.Username()
		passwd := faker.Password()
		expectedErr := errors.New("repository failed creating document")

		userRepositoryMock.
			EXPECT().
			CheckUniqueness(context.Background(), username, email).
			Return(true).
			Times(1)

		passwordServiceMock.
			EXPECT().
			HashPassword(passwd).
			Return(passwd, nil).
			Times(1)

		userRepositoryMock.
			EXPECT().
			Create(context.Background(), user.User{Email: email, Username: username, Password: passwd}).
			Return(user.User{}, expectedErr).
			Times(1)

		service := user.NewService(userRepositoryMock, tokenServiceMock, passwordServiceMock)
		_, _, err := service.RegisterUser(ctx, username, email, passwd)
		assert.True(t, errors.Is(err, expectedErr), "expecting error to be: '%s' but got: '%s' instead", expectedErr.Error(), err.Error())
	})

	t.Run("register user failed - failed generating access token", func(t *testing.T) {
		userRepositoryMock := mockUser.NewMockRepository(ctrl)
		tokenServiceMock := mockToken.NewMockService(ctrl)
		passwordServiceMock := mockPassword.NewMockService(ctrl)

		email := faker.Email()
		username := faker.Username()
		passwd := faker.Password()
		expectedErr := errors.New("failed generating tokena")
		objID := primitive.NewObjectID()

		userRepositoryMock.
			EXPECT().
			CheckUniqueness(context.Background(), username, email).
			Return(true).
			Times(1)

		passwordServiceMock.
			EXPECT().
			HashPassword(passwd).
			Return(passwd, nil).
			Times(1)

		userRepositoryMock.
			EXPECT().
			Create(context.Background(), user.User{Email: email, Username: username, Password: passwd}).
			Return(user.User{ID: objID, Email: email, Username: username, Password: passwd}, nil).
			Times(1)

		tokenServiceMock.
			EXPECT().
			SignPayload(token.TokenPayload{ID: objID.Hex(), Username: username}).
			Return(gomock.Any().String(), expectedErr).
			Times(1)

		service := user.NewService(userRepositoryMock, tokenServiceMock, passwordServiceMock)
		_, _, err := service.RegisterUser(ctx, username, email, passwd)
		assert.True(t, errors.Is(err, expectedErr), "expecting error to be: '%s' but got: '%s' instead", expectedErr.Error(), err.Error())
	})

}

func Test_service_GetUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	t.Run("failed get user", func(t *testing.T) {
		userRepositoryMock := mockUser.NewMockRepository(ctrl)
		tokenServiceMock := mockToken.NewMockService(ctrl)
		passwordServiceMock := mockPassword.NewMockService(ctrl)

		ctx := context.Background()
		objID := primitive.NewObjectID()
		expectedErr := mongo.ErrNoDocuments

		userRepositoryMock.
			EXPECT().
			FindByID(ctx, objID.Hex()).
			Return(user.User{}, mongo.ErrNoDocuments)

		s := user.NewService(userRepositoryMock, tokenServiceMock, passwordServiceMock)
		_, err := s.GetUser(ctx, objID.Hex())
		assert.True(t, errors.Is(err, expectedErr), fmt.Sprintf(`expecting error to be: "%s", but got: "%s"`, expectedErr.Error(), err.Error()))

	})

	t.Run("success get user", func(t *testing.T) {
		userRepositoryMock := mockUser.NewMockRepository(ctrl)
		tokenServiceMock := mockToken.NewMockService(ctrl)
		passwordServiceMock := mockPassword.NewMockService(ctrl)

		ctx := context.Background()
		objID := primitive.NewObjectID()
		uname := faker.Username()
		passwd := faker.Password()
		resultUser := user.User{ID: objID, Username: uname, Password: passwd}

		userRepositoryMock.
			EXPECT().
			FindByID(ctx, objID.Hex()).
			Return(resultUser, nil).
			Times(1)

		s := user.NewService(userRepositoryMock, tokenServiceMock, passwordServiceMock)
		eUser, err := s.GetUser(ctx, objID.Hex())
		assert.Nil(t, err, "expecting err to be nill")
		assert.Equal(t, resultUser.ID, eUser.ID)
		assert.Equal(t, resultUser.Email, eUser.Email)
		assert.Equal(t, resultUser.Password, eUser.Password)
	})
}

func Test_service_Login(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	t.Run("success login an user", func(t *testing.T) {
		ctx := context.Background()
		objID := primitive.NewObjectID()
		username := faker.Username()
		passwd := faker.Password()
		hashedPassword := faker.Password()
		jwtToken := faker.Jwt()

		userRepositoryMock := mockUser.NewMockRepository(ctrl)
		tokenServiceMock := mockToken.NewMockService(ctrl)
		passwordServiceMock := mockPassword.NewMockService(ctrl)

		userRepositoryMock.
			EXPECT().
			FindOne(ctx, username).
			Return(user.User{Password: hashedPassword, ID: objID, Username: username, Email: faker.Email()}, nil).
			Times(1)

		passwordServiceMock.
			EXPECT().
			ComparePassword(hashedPassword, passwd).
			Return(nil).
			Times(1)

		tokenServiceMock.
			EXPECT().
			SignPayload(token.TokenPayload{
				ID:       objID.Hex(),
				Username: username,
			}).Return(jwtToken, nil)

		s := user.NewService(userRepositoryMock, tokenServiceMock, passwordServiceMock)
		rUser, at, err := s.Login(ctx, username, passwd)
		assert.Nil(t, err)
		assert.Equal(t, username, rUser.Username)
		assert.Equal(t, jwtToken, at)
	})

	t.Run("login an user failed - user not foud", func(t *testing.T) {
		ctx := context.Background()
		username := faker.Username()
		passwd := faker.Password()
		expectedErr := user.ErrInvalidUsernamePassword

		userRepositoryMock := mockUser.NewMockRepository(ctrl)
		tokenServiceMock := mockToken.NewMockService(ctrl)
		passwordServiceMock := mockPassword.NewMockService(ctrl)

		userRepositoryMock.
			EXPECT().
			FindOne(ctx, username).
			Return(user.User{}, expectedErr).
			Times(1)

		s := user.NewService(userRepositoryMock, tokenServiceMock, passwordServiceMock)
		_, _, err := s.Login(ctx, username, passwd)
		assert.True(t, errors.Is(err, expectedErr))
	})

	t.Run("login an user failed - password wrong", func(t *testing.T) {
		ctx := context.Background()
		objID := primitive.NewObjectID()
		username := faker.Username()
		passwd := faker.Password()
		hashedPassword := faker.Password()
		expectedErr := user.ErrInvalidUsernamePassword

		userRepositoryMock := mockUser.NewMockRepository(ctrl)
		tokenServiceMock := mockToken.NewMockService(ctrl)
		passwordServiceMock := mockPassword.NewMockService(ctrl)

		userRepositoryMock.
			EXPECT().
			FindOne(ctx, username).
			Return(user.User{
				ID:       objID,
				Password: hashedPassword,
				Email:    username,
			}, nil).
			Times(1)

		passwordServiceMock.
			EXPECT().
			ComparePassword(hashedPassword, passwd).
			Return(expectedErr).
			Times(1)

		s := user.NewService(userRepositoryMock, tokenServiceMock, passwordServiceMock)
		_, _, err := s.Login(ctx, username, passwd)
		assert.True(t, errors.Is(err, expectedErr))
	})

	t.Run("login an user - failed generating access token", func(t *testing.T) {
		ctx := context.Background()
		objID := primitive.NewObjectID()
		username := faker.Username()
		passwd := faker.Password()
		hashedPassword := faker.Password()
		expectedErr := errors.New("something")

		userRepositoryMock := mockUser.NewMockRepository(ctrl)
		tokenServiceMock := mockToken.NewMockService(ctrl)
		passwordServiceMock := mockPassword.NewMockService(ctrl)

		userRepositoryMock.
			EXPECT().
			FindOne(ctx, username).
			Return(user.User{Password: hashedPassword, ID: objID, Username: username, Email: faker.Email()}, nil).
			Times(1)

		passwordServiceMock.
			EXPECT().
			ComparePassword(hashedPassword, passwd).
			Return(nil).
			Times(1)

		tokenServiceMock.
			EXPECT().
			SignPayload(token.TokenPayload{
				ID:       objID.Hex(),
				Username: username,
			}).Return(gomock.Any().String(), expectedErr)

		s := user.NewService(userRepositoryMock, tokenServiceMock, passwordServiceMock)
		_, _, err := s.Login(ctx, username, passwd)
		assert.NotNil(t, err)
		assert.True(t, errors.Is(err, expectedErr))
	})
}
