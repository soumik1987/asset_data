package handlers

import (
	"net/http"

	echo "github.com/labstack/echo/v4"

	log "github.com/labstack/gommon/log"

	"github.com/soumik1987/asset_price/helpers"
	"github.com/soumik1987/asset_price/requests"
	"github.com/soumik1987/asset_price/response"
)

type IDefiProtocol interface {
	FetchSpotPrices(tokenAddress string) ([]*response.Price, error)
}

type UniswapHandler struct {
	Protocol IDefiProtocol
}

func NewUniswapHandler(protocol IDefiProtocol) *UniswapHandler {
	return &UniswapHandler{
		Protocol: protocol,
	}
}

func (p *UniswapHandler) GetPrice(ctx echo.Context) error {
	reqParams := &requests.PriceRequest{}
	if err := helpers.BindQueryParams(ctx, reqParams); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"message": "missing address"})
	}
	// Other address validations has to be done As well
	if reqParams.TokenAddress == "" {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"message": "empty address"})
	}
	result, err := p.Protocol.FetchSpotPrices(reqParams.TokenAddress)

	// Error wrapping cam be done with fmt
	// Other errors also should be handled
	switch {
	case err != nil:
		log.Errorf("API returning Error", err)
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"message": "Server Error"})

	default:
		return ctx.JSON(http.StatusOK, result)
	}
}
