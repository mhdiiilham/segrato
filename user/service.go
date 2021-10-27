package user

import (
	"context"
	"errors"

	"github.com/mhdiiilham/segrato/pkg/password"
	"github.com/mhdiiilham/segrato/pkg/token"
	"golang.org/x/crypto/bcrypt"
)

type service struct {
	userRepository Repository
	token          token.Service
	p              password.Service
}

func NewService(userRepository Repository, token token.Service, p password.Service) Service {
	return &service{
		userRepository: userRepository,
		token:          token,
		p:              p,
	}
}

func (s *service) RegisterUser(ctx context.Context, username, plainPassword string) (u User, accessToken string, err error) {
	if !s.userRepository.CheckUniqueness(ctx, username) {
		err = errors.New("username already taken")
		return
	}

	hashedPassword, hashedPasswordErr := s.p.HashPassword(plainPassword)
	if hashedPasswordErr != nil {
		err = hashedPasswordErr
		return
	}

	u.Password = hashedPassword
	u.Username = username

	u, err = s.userRepository.Create(ctx, u)
	if err != nil {
		return
	}

	accessToken, err = s.token.SignPayload(token.TokenPayload{
		ID:       u.ID.Hex(),
		Username: u.Username,
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

func (s *service) Login(ctx context.Context, username, password string) (user User, accessToken string, err error) {
	user, err = s.userRepository.FindOne(ctx, username)
	if err != nil {
		return
	}

	if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		err = ErrInvalidUsernamePassword
		return
	}

	accessToken, err = s.token.SignPayload(token.TokenPayload{ID: user.ID.Hex(), Username: user.Username})
	if err != nil {
		return
	}

	return
}
