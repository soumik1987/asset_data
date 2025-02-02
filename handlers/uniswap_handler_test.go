package handlers_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"

	"github.com/soumik1987/asset_price/handlers"
	"github.com/soumik1987/asset_price/mocks"
	"github.com/soumik1987/asset_price/response"
)

func TestPriceHandler_GetPrice(t *testing.T) {
	e := echo.New()

	// Create a new request with query parameter
	req := httptest.NewRequest(http.MethodGet, "/v1/asset/price?token_address=0x123", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Setup mock controller
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockPriceService := mocks.NewMockIPriceService(ctrl)

	// Create expected response
	expectedPrices := []*response.Price{
		{
			Price:     "100.50",
			Timestamp: "2025-01-31T00:00:00Z",
		},
	}

	// Set expectations on mock
	mockPriceService.EXPECT().
		FetchSpotPrices("0x123").
		Return(expectedPrices, nil).
		Times(1)

	// Initialize handler with mock service
	handler := handlers.NewUniswapHandler(mockPriceService)

	err := handler.GetPrice(c)

	// Assertions
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)

	assert.Contains(t, rec.Body.String(), "100.50")
	assert.Contains(t, rec.Body.String(), "2025-01-31")
}
