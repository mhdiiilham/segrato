package auth

import (
	"github.com/mhdiiilham/segrato/config"
	"github.com/mhdiiilham/segrato/internal/auth/model/user"
	"github.com/sirupsen/logrus"
)

type Server struct {
	cfg         config.Config
	userService user.Service
}

func NewServer(cfg config.Config, userService user.Service) *Server {
	logrus.Info("creating new API Server")
	return &Server{
		cfg:         cfg,
		userService: userService,
	}
}
