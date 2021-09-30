package service

import (
	"context"
	"errors"

	"github.com/mhdiiilham/segrato/pkg/token"
	"github.com/mhdiiilham/segrato/user"
	"golang.org/x/crypto/bcrypt"
)

type userService struct {
	userRepository user.Repository
	token          token.Service
}

func NewUserService(userRepository user.Repository, tokenService token.Service) user.Service {
	return &userService{
		userRepository: userRepository,
		token:          tokenService,
	}
}

func (s userService) RegisterUser(ctx context.Context, username, plainPassword string, blockWords []string) (u user.User, accessToken string, err error) {
	var bytePassword []byte

	if !s.userRepository.CheckUniqueness(ctx, username) {
		err = errors.New("username already taken")
		return
	}

	bytePassword, err = bcrypt.GenerateFromPassword([]byte(plainPassword), bcrypt.MinCost)
	if err != nil {
		return
	}

	password := string(bytePassword)
	u.Password = password
	u.Username = username
	u.BlockedWords = blockWords

	u, err = s.userRepository.Create(ctx, u)
	if err != nil {
		return
	}

	accessToken, err = s.token.SignPayload(token.TokenPayload{
		ID:        u.ID.Hex(),
		Username:  u.Username,
		IsPremium: u.IsPremium,
	})
	if err != nil {
		return
	}

	return
}

func (s userService) GetUser(ctx context.Context, username string) (u user.User, err error) {
	return
}
