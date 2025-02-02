package server

import (
	"context"
	"fmt"

	echo "github.com/labstack/echo/v4"
	log "github.com/labstack/gommon/log"

	"github.com/soumik1987/asset_price/config"
)

type Server struct {
	Echo   *echo.Echo
	Config *config.Config
}

// Create a new server
func New(ctx context.Context, config *config.Config) *Server {
	return &Server{
		Echo:   echo.New(),
		Config: config,
	}
}

func (server *Server) Start() {
	server.Echo.Logger.SetLevel(log.DEBUG)
	server.Echo.HideBanner = true

	server.configureRoutes()
	// We can add swagger for documentation

	addr := fmt.Sprintf("%v:%v", server.Config.Http.Host, server.Config.Http.Port)
	err := server.Echo.Start(addr)
	if err != nil {
		log.Errorf("web server failed to start, Echo.start error: %v", err)
	}
}
