package response

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"time"
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
	Data    []any `json:"data" example:"Error message 1, Error message 2"`
}

func FailureResponse(code int, data ...any) error {
	rsp := &errorResponse{
		Success: false,
		Data:    data,
	}

	return echo.NewHTTPError(code, rsp)
}

func SetSecureCookies(ctx echo.Context, keyName string, value string, expires time.Time, maxAge int) {
	ctx.SetCookie(&http.Cookie{
		Name:     keyName,
		Value:    value,
		Expires:  expires,
		MaxAge:   maxAge / 1000,
		Secure:   true,
		HttpOnly: true,
		SameSite: http.SameSiteStrictMode,
	})
}
