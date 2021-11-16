package game

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/mhdiiilham/segrato/config"
	"github.com/mhdiiilham/segrato/internal/apiresponse"
	authGRPCMock "github.com/mhdiiilham/segrato/internal/game/mock"
	"github.com/mhdiiilham/segrato/internal/proto"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func Test_server_HealtCheck(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	ctx := context.Background()

	t.Run("server is running well", func(t *testing.T) {
		var resp apiresponse.HealtCheck
		expectedAuthHealthResp := proto.HealthCheckResponse{
			ServerRunning:     true,
			MongoDBConnection: true,
		}

		authClient := authGRPCMock.NewMockAuthServiceClient(ctrl)
		authClient.
			EXPECT().
			HealthCheck(ctx, new(emptypb.Empty)).
			Return(&expectedAuthHealthResp, nil).
			Times(1)

		httpCodeExpected := http.StatusOK
		msgExpected := "game server is running"

		w := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodGet, "/games/health-check", nil)
		defer r.Body.Close()

		srv, srvErr := NewServer(config.Config{}, authClient)
		assert.NoError(t, srvErr)
		routes := srv.CORS(srv.HandlerLogging(srv.Routes(ctx)))
		routes.ServeHTTP(w, r)

		assert.NoError(t, json.NewDecoder(w.Body).Decode(&resp), "should be no error when decoding response body")

		assert.Equal(t, http.StatusOK, w.Code, "expecting http code %d, but got %d", httpCodeExpected, w.Code)
		assert.Equal(t, httpCodeExpected, resp.Code, "expecting http code %d, but got %d", httpCodeExpected, resp.Code)
		assert.Equal(t, msgExpected, resp.Message, "apiresponse.HealtCheck.Message = %s, want = %s", resp.Message, msgExpected)
		assert.True(t, resp.AuthHealth.Server)
		assert.True(t, resp.AuthHealth.MongoConnection)
	})

	t.Run("auth client in compromise", func(t *testing.T) {
		var resp apiresponse.HealtCheck

		authClient := authGRPCMock.NewMockAuthServiceClient(ctrl)
		authClient.
			EXPECT().
			HealthCheck(ctx, new(emptypb.Empty)).
			Return(nil, status.Error(codes.Internal, "internal server error")).
			Times(1)

		httpCodeExpected := http.StatusOK
		msgExpected := "game server is running"

		w := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodGet, "/games/health-check", nil)
		defer r.Body.Close()

		srv, srvErr := NewServer(config.Config{}, authClient)
		assert.NoError(t, srvErr)
		routes := srv.CORS(srv.HandlerLogging(srv.Routes(ctx)))
		routes.ServeHTTP(w, r)

		assert.NoError(t, json.NewDecoder(w.Body).Decode(&resp), "should be no error when decoding response body")

		assert.Equal(t, http.StatusOK, w.Code, "expecting http code %d, but got %d", httpCodeExpected, w.Code)
		assert.Equal(t, httpCodeExpected, resp.Code, "expecting http code %d, but got %d", httpCodeExpected, resp.Code)
		assert.Equal(t, msgExpected, resp.Message, "apiresponse.HealtCheck.Message = %s, want = %s", resp.Message, msgExpected)
		assert.False(t, resp.AuthHealth.Server)
		assert.False(t, resp.AuthHealth.MongoConnection)
	})

}
