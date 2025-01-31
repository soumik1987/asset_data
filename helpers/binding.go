package helpers

import (
	echo "github.com/labstack/echo/v4"
)

// Cretaing a seperate module for Parameters binding
// we can add PAth or Body param binding here as separate functions and use all across the handler
func BindQueryParams(ctx echo.Context, payload interface{}) (err error) {
	return (&echo.DefaultBinder{}).BindQueryParams(ctx, payload)
}
