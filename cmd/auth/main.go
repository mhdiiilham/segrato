package main

import (
	"context"
	"os/signal"
	"syscall"

	"github.com/mhdiiilham/segrato/config"
	"github.com/mhdiiilham/segrato/internal/auth"
	"github.com/mhdiiilham/segrato/pkg/db"
	"github.com/mhdiiilham/segrato/pkg/server"
	"github.com/mhdiiilham/segrato/pkg/token"
	"github.com/mhdiiilham/segrato/user"
	"github.com/sirupsen/logrus"
)

func main() {
	ctx, done := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)

	logrus.SetFormatter(&logrus.TextFormatter{
		DisableColors: true,
		FullTimestamp: true,
	})

	defer func() {
		done()
		if r := recover(); r != nil {
			logrus.Fatalf("application panic: %w", r)
		}
	}()

	err := realMain(ctx)
	done()

	if err != nil {
		panic(err)
	}

	logrus.Info("successfully shutdown")
}

func realMain(ctx context.Context) error {
	cfg, cfgErr := config.ReadConfig()
	if cfgErr != nil {
		return cfgErr
	}

	mongoDB, err := db.NewMongoDBConnection(cfg.MongoDBURI)
	if err != nil {
		return err
	}

	database := mongoDB.Database(cfg.Database)
	userCollection := database.Collection("user")
	userRepository := user.NewRepository(userCollection)

	tokenService := token.TokenService{Config: &cfg}
	userService := user.NewService(userRepository, tokenService)

	segratoAPI, err := auth.NewServer(cfg, userService)
	if err != nil {
		return err
	}

	srv, err := server.New(cfg.Port.Auth)
	if err != nil {
		return err
	}

	logrus.Infof("listening on: %v", cfg.Port.Auth)
	return srv.ServeHTTPHandler(ctx, segratoAPI.Routes(ctx))

}
