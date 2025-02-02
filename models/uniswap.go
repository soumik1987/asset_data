package models

import (
	h "github.com/soumik1987/asset_price/handlers"
	"github.com/soumik1987/asset_price/response"
)

type IUniswapService interface {
	FetchSpotPrices(tokenAddress string) ([]*response.Price, error)
}

type Uniswap struct {
	uniswapService IUniswapService
}

func NewUniswap(u IUniswapService) h.IDefiProtocol {
	return &Uniswap{
		uniswapService: u,
	}
}

func (u *Uniswap) FetchSpotPrices(tokenAddress string) ([]*response.Price, error) {

	return u.uniswapService.FetchSpotPrices(tokenAddress)
}
