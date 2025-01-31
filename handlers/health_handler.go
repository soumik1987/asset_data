package handlers

import (
	"net/http"

	echo "github.com/labstack/echo/v4"
)

// Deployment will be made to Kubernetese. Hence we need a livness and readyness probe
// for health check. Bith are basic but configurable with any advanced criteria
// like if redis is Up or DB connection is up
type HealthHandler struct {
}

func NewHealthHandler() *HealthHandler {
	return &HealthHandler{}
}

// Liveness
func (h *HealthHandler) LiveProbe(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, map[string]string{"message": "Ok"})
}

// Readuyness
func (h *HealthHandler) ReadyProbe(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, map[string]string{"message": "Ok"})
}
