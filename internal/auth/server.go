package auth

import (
	"context"
	"net/http"

	"github.com/gofiber/adaptor/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/mhdiiilham/segrato/config"
	"github.com/mhdiiilham/segrato/user"
	"github.com/sirupsen/logrus"
)

type Server struct {
	cfg         config.Config
	userService user.Service
}

func NewServer(cfg config.Config, userService user.Service) (*Server, error) {
	logrus.Info("creating new API Server")
	return &Server{
		cfg:         cfg,
		userService: userService,
	}, nil
}

func (s *Server) Routes(ctx context.Context) http.Handler {
	logrus.Info("initializing API routes'")

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

	api := app.Group("/api")
	v1 := api.Group("/v1")

	v1.Post("/users", s.RegisterUser)
	v1.Get("/users/:userid", s.GetUser)
	v1.Post("users/login", s.Login)

	logrus.Info("API routes ready")
	return adaptor.FiberApp(app)
}
