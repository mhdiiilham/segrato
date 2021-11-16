package auth

import (
	"context"
	"errors"

	"github.com/mhdiiilham/segrato/internal/auth/model/user"
	"github.com/mhdiiilham/segrato/internal/proto"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) Login(ctx context.Context, request *proto.LoginRequest) (*proto.LoginResponse, error) {
	username := request.GetUsername()
	password := request.GetPassword()

	u, accessToken, err := s.userService.Login(ctx, username, password)
	if err != nil {
		logrus.Errorf("error: %v", err)
		if errors.Is(err, mongo.ErrNoDocuments) || errors.Is(err, user.ErrInvalidUsernamePassword) {
			return nil, status.Errorf(codes.InvalidArgument, "invalid username or/and password")
		}

		logrus.Errorf("error calling userService.Login: %v", err)
		return nil, status.Error(codes.Internal, "internal server error")
	}

	return &proto.LoginResponse{
		Id:          u.ID.Hex(),
		AccessToken: accessToken,
	}, nil
}
