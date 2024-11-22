package utils

import (
	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"net/http"
	"orderbook/internal/pkg/response"
)

func Contains(slice []string, item string) bool {
	for _, value := range slice {
		if value == item {
			return true
		}
	}
	return false
}

type Validator interface{}

func ValidateStruct[T Validator](ctx echo.Context, validator *validator.Validate, obj *T) (*T, error) {
	if err := ctx.Bind(obj); err != nil {
		return nil, response.FailureResponse(http.StatusBadRequest, err.Error())
	}
	if err := validator.Struct(obj); err != nil {
		return nil, response.FailureResponse(http.StatusBadRequest, echo.Map{
			"Message": "Invalid input",
			"Error":   err.Error(),
		})
	}
	return obj, nil
}
