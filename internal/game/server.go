package game

import (
	"context"
	"net/http"

	"github.com/gofiber/adaptor/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/mhdiiilham/segrato/config"
	"github.com/sirupsen/logrus"
)

type server struct {
	cfg config.Config
}

func NewServer(cfg config.Config) (*server, error) {
	return &server{
		cfg: cfg,
	}, nil
}

func (s *server) Routes(ctx context.Context) http.Handler {
	logrus.Info("initializing game API routes'")

	logrus.Info("initializing fiber app and logger")
	app := fiber.New()
	app.Use(logger.New(logger.Config{
		Format:     "${pid} ${status} - ${method} ${path}\n",
		TimeFormat: "02-Jan-2006",
		TimeZone:   "Asia/Jakarta",
	}))

	logrus.Info("setting up fiber cors config")
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "*",
	}))

	logrus.Info("game api is ready")
	return adaptor.FiberApp(app)
}
