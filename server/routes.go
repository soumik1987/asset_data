package server

import (
	"github.com/soumik1987/asset_price/handlers"
	"github.com/soumik1987/asset_price/models"
	"github.com/soumik1987/asset_price/services"
)

func (server *Server) configureRoutes() {
	uniswapService := services.NewUniswapService(server.Config.GraphApiKey)
	// Model here is redundant. But it is kept for future uses
	uniswapModel := models.NewUniswap(uniswapService)

	// price handler - PRIVATE
	// Authentication should be added here
	priceHandler := handlers.NewUniswapHandler(uniswapModel)
	server.Echo.GET("/v1/asset/price", priceHandler.GetPrice)

	// Healthchek
	healthHandler := handlers.NewHealthHandler()
	server.Echo.GET("/live", healthHandler.LiveProbe)
	server.Echo.GET("/ready", healthHandler.ReadyProbe)
}
