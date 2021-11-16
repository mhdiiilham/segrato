package auth

import (
	"context"
	"errors"

	"github.com/mhdiiilham/segrato/internal/auth/model/user"
	"github.com/mhdiiilham/segrato/internal/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) RegisterUser(ctx context.Context, request *proto.RegisterUserRequest) (*proto.RegisterUserResponse, error) {
	username := request.GetUsername()
	email := request.GetEmail()
	password := request.GetPassword()

	if email == "" || password == "" || len(password) < 8 {
		return nil, status.Errorf(codes.InvalidArgument, "invalid email or/and password")
	}

	newUser, accessToken, err := s.userService.RegisterUser(ctx, username, email, password)
	if err != nil {
		if errors.Is(err, user.ErrUsernameEmailRegistered) {
			return nil, status.Errorf(codes.AlreadyExists, err.Error())
		}

		return nil, status.Error(codes.Internal, "internal server error")
	}

	return &proto.RegisterUserResponse{
		Id: newUser.ID.Hex(),
		User: &proto.User{
			Id:       newUser.ID.Hex(),
			Username: newUser.Username,
			Email:    newUser.Email,
		},
		AccessToken: accessToken,
	}, nil
}
