package auth

import (
	"context"
	"errors"
	"time"

	"github.com/mhdiiilham/segrato/internal/proto"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) GetUser(ctx context.Context, request *proto.GetUserRequest) (*proto.GetUserResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	userID := request.GetUserId()

	u, err := s.userService.GetUser(ctx, userID)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, status.Errorf(codes.NotFound, "user with id %s not found", userID)
		}
		return nil, status.Error(codes.Internal, "internal server error")
	}

	return &proto.GetUserResponse{
		Code: "00",
		User: &proto.User{
			Id:       u.ID.Hex(),
			Username: u.Username,
			Email:    u.Email,
		},
	}, nil
}
