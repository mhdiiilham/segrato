package game

import (
	"context"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/mhdiiilham/segrato/config"
	"github.com/rs/cors"
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
	mux := mux.NewRouter()

	mux.HandleFunc("/games/health-check", s.HealtCheck).
		Methods(http.MethodGet)

	logrus.Info("game api is ready")
	return mux
}

func (s *server) CORS(mux http.Handler) http.Handler {
	logrus.Info("initilizing CORS")

	c := cors.New(cors.Options{
		AllowedHeaders: []string{"*"},
	})

	return c.Handler(mux)
}

func (s *server) HandlerLogging(mux http.Handler) http.Handler {
	logrus.Info("initilizing handler logging")
	return handlers.LoggingHandler(os.Stdout, mux)
}
