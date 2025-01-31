package handlers

import (
	"net/http"

	echo "github.com/labstack/echo/v4"

	"github.com/soumik1987/asset_price/helpers"
	"github.com/soumik1987/asset_price/requests"
	"github.com/soumik1987/asset_price/services"
)

type IPriceHandler interface {
	GetPrice(ctx echo.Context) error
}

type PriceHandler struct {
	priceService services.IPriceService
}

func NewPriceHandler(priceService services.IPriceService) IPriceHandler {
	return &PriceHandler{
		priceService: priceService,
	}
}

func (p *PriceHandler) GetPrice(ctx echo.Context) error {
	reqParams := &requests.PriceRequest{}
	if err := helpers.BindQueryParams(ctx, reqParams); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"message": "missing address"})
	}
	// Other address validations has to be done As well
	if reqParams.TokenAddress == "" {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"message": "empty address"})
	}
	result, err := p.priceService.FetchSpotPrices(reqParams.TokenAddress)

	// Error wrapping cam be done with fmt
	// Other errors also should be handled
	switch {
	case err != nil:
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"message": "Server Error"})

	default:
		return ctx.JSON(http.StatusOK, result)
	}

}
