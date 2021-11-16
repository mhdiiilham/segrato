package main

import (
	"context"
	"os/signal"
	"syscall"

	"github.com/mhdiiilham/segrato/config"
	"github.com/mhdiiilham/segrato/internal/game"
	"github.com/mhdiiilham/segrato/internal/proto"
	"github.com/mhdiiilham/segrato/pkg/server"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
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

	gRPCConnection, err := realMain(ctx)
	done()

	if err != nil {
		panic(err)
	}

	logrus.Info("disconnection from auth gRPC Server")
	gRPCConnection.Close()
	logrus.Info("successfully shutdown")
}

func realMain(ctx context.Context) (*grpc.ClientConn, error) {
	cfg, cgfErr := config.ReadConfig()
	if cgfErr != nil {
		return nil, cgfErr
	}

	conn, err := grpc.Dial("localhost:8088", grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	authRPC := proto.NewAuthServiceClient(conn)
	gameAPI, err := game.NewServer(cfg, authRPC)
	if err != nil {
		return nil, err
	}

	srv, err := server.New(cfg.Port.Game)
	if err != nil {
		return nil, err
	}

	logrus.Infof("listening game API on: %v", cfg.Port.Game)
	return conn, srv.ServeHTTPHandler(ctx, gameAPI.CORS(gameAPI.HandlerLogging(gameAPI.Routes(ctx))))
}
