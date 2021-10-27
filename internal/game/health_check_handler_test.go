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
	ctx := context.Background()
	var resp apiresponse.HealtCheck

	httpCodeExpected := http.StatusOK
	msgExpected := "game server is running"

	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/games/health-check", nil)
	defer r.Body.Close()

	srv, srvErr := NewServer(config.Config{})
	assert.NoError(t, srvErr)
	routes := srv.CORS(srv.HandlerLogging(srv.Routes(ctx)))
	routes.ServeHTTP(w, r)

	assert.NoError(t, json.NewDecoder(w.Body).Decode(&resp), "should be no error when decoding response body")

	assert.Equal(t, http.StatusOK, w.Code, "expecting http code %d, but got %d", httpCodeExpected, w.Code)
	assert.Equal(t, httpCodeExpected, resp.Code, "expecting http code %d, but got %d", httpCodeExpected, resp.Code)
	assert.Equal(t, msgExpected, resp.Message, "apiresponse.HealtCheck.Message = %s, want = %s", resp.Message, msgExpected)
}
