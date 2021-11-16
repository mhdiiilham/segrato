package game

import (
	"context"

	"github.com/mhdiiilham/segrato/internal/proto"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

type AuthServiceClient interface {
	RegisterUser(ctx context.Context, in *proto.RegisterUserRequest, opts ...grpc.CallOption) (*proto.RegisterUserResponse, error)
	GetUser(ctx context.Context, in *proto.GetUserRequest, opts ...grpc.CallOption) (*proto.GetUserResponse, error)
	Login(ctx context.Context, in *proto.LoginRequest, opts ...grpc.CallOption) (*proto.LoginResponse, error)
	GetUserByAccessToken(ctx context.Context, in *proto.GetUserByAccessTokenRequest, opts ...grpc.CallOption) (*proto.GetUserByAccessTokenResponse, error)
	HealthCheck(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*proto.HealthCheckResponse, error)
}
