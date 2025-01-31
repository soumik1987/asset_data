package main

import (
	"context"

	log "github.com/labstack/gommon/log"
	"github.com/soumik1987/asset_price/config"
	"github.com/soumik1987/asset_price/server"
)

func main() {
	config := config.Load()
	log.Info("Started Web Api")

	server := server.New(context.Background(), config)
	server.Start()

	log.Info("Stopping Web Api")
}
