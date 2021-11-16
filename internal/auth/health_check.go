package auth

import (
	"context"

	"github.com/mhdiiilham/segrato/internal/auth/model/proto"
	"github.com/sirupsen/logrus"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *Server) HealthCheck(ctx context.Context, req *emptypb.Empty) (*proto.HealthCheckResponse, error) {

	mongoStatus := true
	if err := s.userService.PingMongoDB(ctx); err != nil {
		logrus.Errorf("mongodb connection is compromise %v", err)
		mongoStatus = false
	}

	return &proto.HealthCheckResponse{
		MongoDBConnection: mongoStatus,
		ServerRunning:     true,
	}, nil
}
