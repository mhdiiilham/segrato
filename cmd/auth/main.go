package main

import (
	"context"
	"os/signal"
	"syscall"

	"github.com/mhdiiilham/segrato/config"
	"github.com/mhdiiilham/segrato/internal/auth"
	"github.com/mhdiiilham/segrato/pkg/db"
	"github.com/mhdiiilham/segrato/pkg/password"
	"github.com/mhdiiilham/segrato/pkg/server"
	"github.com/mhdiiilham/segrato/pkg/token"
	"github.com/mhdiiilham/segrato/user"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
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
			logrus.Fatalf("application panic: %v", r)
		}
	}()

	cfg, cfgErr := config.ReadConfig()
	if cfgErr != nil {
		panic(cfgErr)
	}

	mongoDB, err := db.NewMongoDBConnection(cfg.MongoDBURI)
	if err != nil {
		panic(err)
	}

	err = realMain(ctx, cfg, mongoDB.Database(cfg.Database).Collection("user"))
	done()

	if err != nil {
		panic(err)
	}

	logrus.Infof("disconneting mongoDB Client. Error: %v", mongoDB.Disconnect(context.Background()))
	logrus.Info("successfully shutdown")
}

func realMain(ctx context.Context, cfg config.Config, userCollection *mongo.Collection) error {
	userRepository := user.NewRepository(userCollection)
	passwordService := password.NewService()

	tokenService := token.TokenService{Config: &cfg}
	userService := user.NewService(userRepository, tokenService, passwordService)

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
