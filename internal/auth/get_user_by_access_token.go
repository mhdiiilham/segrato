package auth

import (
	"context"
	"errors"

	"github.com/mhdiiilham/segrato/internal/auth/model/user"
	"github.com/mhdiiilham/segrato/internal/proto"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) GetUserByAccessToken(ctx context.Context, request *proto.GetUserByAccessTokenRequest) (*proto.GetUserByAccessTokenResponse, error) {
	accessToken := request.GetAccessToken()
	if accessToken == "" {
		return nil, status.Error(codes.InvalidArgument, "accessToken couldn't be empty")
	}

	u, err := s.userService.GetUserByAccessToken(ctx, accessToken)
	if err != nil {
		if errors.Is(err, user.ErrAccessTokenInvalid) || errors.Is(err, mongo.ErrNoDocuments) {
			return nil, status.Error(codes.InvalidArgument, "invalid access token")
		}

		return nil, status.Error(codes.Internal, "internal server error")
	}
	return &proto.GetUserByAccessTokenResponse{
		User: &proto.User{
			Id:       u.ID.Hex(),
			Email:    u.Email,
			Username: u.Username,
		},
	}, nil
}
