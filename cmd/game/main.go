package main

import (
	"context"
	"os/signal"
	"syscall"

	"github.com/mhdiiilham/segrato/config"
	"github.com/mhdiiilham/segrato/internal/game"
	"github.com/mhdiiilham/segrato/pkg/server"
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
	cfg, cgfErr := config.ReadConfig()
	if cgfErr != nil {
		return cgfErr
	}

	gameAPI, err := game.NewServer(cfg)
	if err != nil {
		return err
	}

	srv, err := server.New(cfg.Port.Game)
	if err != nil {
		return err
	}

	logrus.Infof("listening game API on: %v", cfg.Port.Game)
	return srv.ServeHTTPHandler(ctx, gameAPI.Routes(ctx))
}
