package server

import (
	"github.com/soumik1987/asset_price/handlers"
	"github.com/soumik1987/asset_price/services"
)

func (server *Server) configureRoutes() {
	priceService := services.NewPriceService(server.Config.ApiKey)

	// price handler - PRIVATE
	// Authentication should be added here
	priceHandler := handlers.NewPriceHandler(priceService)
	server.Echo.GET("/v1/asset/price", priceHandler.GetPrice)

	// Healthchek
	healthHandler := handlers.NewHealthHandler()
	server.Echo.GET("/live", healthHandler.LiveProbe)
	server.Echo.GET("/ready", healthHandler.ReadyProbe)
}
