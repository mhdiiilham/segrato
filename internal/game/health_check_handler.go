package game

import (
	"encoding/json"
	"net/http"

	"github.com/mhdiiilham/segrato/internal/apiresponse"
)

func (s *server) HealtCheck(w http.ResponseWriter, r *http.Request) {
	resp := apiresponse.HealtCheck{
		Code:    http.StatusOK,
		Message: "game server is running",
	}

	jsonResponse, _ := json.Marshal(resp)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}
