package hhttp

import (
	"github.com/gintokos/vpagrSite/internal/config"
	"github.com/gintokos/vpagrSite/internal/transport/http/auth"

	"github.com/gin-gonic/gin"
	"github.com/gintokos/serverdb/pkg/logger"
)

type Server struct {
	*gin.Engine
	Logger *logger.CustomLogger
	Config *config.ServerConfig
}

func NewServer(lg *logger.CustomLogger, cfg *config.ServerConfig) *Server {
	server := Server{
		Config: cfg,
		Logger: lg,
	}

	gsrv := gin.Default()
	server.Engine = gsrv

	return &server
}

func (s *Server) MustStartServer() {
	s.initHandlers()
	s.Logger.Info("Handlers was inited")

	err := s.Run(s.Config.Port)
	if err != nil {
		panic(err)
	}
}

func (s *Server) initHandlers() {
	auth.InitHanlers(s.Engine, *s.Logger)
}
