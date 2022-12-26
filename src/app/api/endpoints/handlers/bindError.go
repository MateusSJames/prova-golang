package handlers

import (
	"Golang/src/app/api/endpoints/dto/response"
	"Golang/src/core/errors"
	"Golang/src/core/messages"
	"net/http"

	"github.com/labstack/echo/v4"
)

func bindError(err errors.Error) response.ErrorMessage {
	dto := response.ErrorMessage{}

	dto.Msg = err.FriendlyMsg()

	switch v := err.(type) {
	case *errors.ValidationErr:
		dto.StatusCode = http.StatusUnprocessableEntity
		fields := v.InvalidFields()
		for _, f := range fields {
			dto.InvalidFields = append(dto.InvalidFields, response.InvalidField{
				FieldName:   f.Name,
				Description: f.Description,
			})
		}
	}

	return dto
}

func getDefaultBadRequestResponse(c echo.Context, fields ...response.InvalidField) error {
	return c.JSON(http.StatusBadRequest, response.ErrorMessage{
		StatusCode:    http.StatusBadRequest,
		Msg:           messages.BadRequestMsg,
		InvalidFields: fields,
	})
}
