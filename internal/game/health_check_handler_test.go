package game

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/mhdiiilham/segrato/config"
	"github.com/mhdiiilham/segrato/internal/apiresponse"
	"github.com/stretchr/testify/assert"
)

func Test_server_HealtCheck(t *testing.T) {
	var resp apiresponse.HealtCheck
	ctx := context.Background()
	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/games/health-check", nil)
	srv, srvErr := NewServer(config.Config{})
	assert.NoError(t, srvErr)
	routes := srv.CORS(srv.HandlerLogging(srv.Routes(ctx)))
	routes.ServeHTTP(w, r)

	respBody := w.Body.Bytes()
	unMarshallErr := json.Unmarshal(respBody, &resp)
	assert.NoError(t, unMarshallErr)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, http.StatusOK, resp.Code)
	assert.Equal(t, "game server is running", resp.Message)
}
