package user

import (
	"context"
	"errors"

	"github.com/mhdiiilham/segrato/pkg/token"
	"golang.org/x/crypto/bcrypt"
)

type service struct {
	userRepository Repository
	token          token.Service
}

func NewService(userRepository Repository, token token.Service) Service {
	return &service{
		userRepository: userRepository,
		token:          token,
	}
}

func (s *service) RegisterUser(ctx context.Context, username, plainPassword string, blockWords []string) (u User, accessToken string, err error) {
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

func (s *service) GetUser(ctx context.Context, userID string) (u User, err error) {
	u, err = s.userRepository.FindByID(ctx, userID)
	if err != nil {
		return
	}
	return
}