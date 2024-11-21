package response

import (
	"github.com/labstack/echo/v4"
)

type response struct {
	Success bool `json:"success" example:"true"`
	Data    any  `json:"data,omitempty"`
}

func SuccessResponse(ctx echo.Context, code int, data any) error {
	rsp := &response{
		Success: true,
		Data:    data,
	}
	return ctx.JSON(code, rsp)
}

type errorResponse struct {
	Success bool  `json:"success" example:"false"`
	Errors  []any `json:"data" example:"Error message 1, Error message 2"`
}

func FailureResponse(code int, data ...any) error {
	rsp := &errorResponse{
		Success: false,
		Errors:  data,
	}

	return echo.NewHTTPError(code, rsp)
}
