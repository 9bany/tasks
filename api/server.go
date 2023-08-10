package api

import (
	db "github.com/9bany/task/db/sqlc"
	httpclient "github.com/9bany/task/http_client"
	"github.com/9bany/task/middleware"
	"github.com/9bany/task/util"
	"github.com/gin-gonic/gin"
)

func NewServer(config util.Config,
	store db.Store,
	iframelyClient httpclient.IframelyClient) (*Server, error) {

	server := &Server{
		config:         config,
		store:          store,
		iframelyClient: iframelyClient,
	}

	server.setupRouters()

	return server, nil
}

type Server struct {
	config         util.Config
	store          db.Store
	router         *gin.Engine
	iframelyClient httpclient.IframelyClient
}

func (server *Server) setupRouters() {
	router := gin.Default()
	
	router.Use(middleware.TimeoutMiddleware())

	router.GET("/api/meta", server.GetSite)

	server.router = router
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}
