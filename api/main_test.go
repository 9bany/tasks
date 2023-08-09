package api

import (
	"os"
	"testing"

	db "github.com/9bany/task/db/sqlc"
	httpclient "github.com/9bany/task/http_client"
	"github.com/9bany/task/util"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
)

func newTestServer(t *testing.T,
	store db.Store,
	iframelyClient httpclient.IframelyClient) *Server {

	config := util.Config{}
	server, err := NewServer(config, store, iframelyClient)
	require.NoError(t, err)
	require.NotEmpty(t, server)
	return server
}

func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode)
	os.Exit(m.Run())
}
