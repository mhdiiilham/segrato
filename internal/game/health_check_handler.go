package game

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/mhdiiilham/segrato/internal/apiresponse"
	"github.com/sirupsen/logrus"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *server) HealtCheck(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	server := true
	mongoDB := true

	respAuth, err := s.authAPI.HealthCheck(ctx, new(emptypb.Empty))
	logrus.Infof("response from auth health check: %+v, err: %v", respAuth, err)
	if err != nil {
		server = false
		mongoDB = false
	}

	server = respAuth.GetServerRunning()
	mongoDB = respAuth.GetMongoDBConnection()

	resp := apiresponse.HealtCheck{
		Code:    http.StatusOK,
		Message: "game server is running",
		AuthHealth: apiresponse.AuthHealth{
			Server:          server,
			MongoConnection: mongoDB,
		},
	}

	jsonResponse, _ := json.Marshal(resp)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}
