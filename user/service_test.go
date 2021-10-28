package user_test

import (
	"context"
	"testing"

	"github.com/bxcodec/faker/v3"
	"github.com/golang/mock/gomock"
	mockPassword "github.com/mhdiiilham/segrato/pkg/password/mock"
	"github.com/mhdiiilham/segrato/pkg/token"
	mockToken "github.com/mhdiiilham/segrato/pkg/token/mock"
	"github.com/mhdiiilham/segrato/user"
	mockUser "github.com/mhdiiilham/segrato/user/mock"
	"github.com/stretchr/testify/assert"
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

}
