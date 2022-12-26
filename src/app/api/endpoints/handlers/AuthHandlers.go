package handlers

import (
	"Golang/src/app/api/endpoints/dto/request"
	"Golang/src/core/interfaces/iservices"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

type AuthHandlers struct {
	service iservices.AuthManager
}

func (handler AuthHandlers) SignUp(c echo.Context) error {
	var dto request.Password

	err := c.Bind(&dto)

	if err != nil {
		getDefaultBadRequestResponse(c)
	}

	

	password, conversionErr := dto.AsDomain()
	if conversionErr != nil {
		handledErr := bindError(conversionErr)
		return c.JSON(handledErr.StatusCode, handledErr)
	}
vfmt.Pr intln(dto)
	return c.JSON(http.StatusOK, dto)
}

func NewAuthHandlers() *AuthHandlers {
	return &AuthHandlers{}
}
