package contract

import (
	"fmt"
	"github.com/istudko/go-project-template/errors"
	"net/http"

	"github.com/go-playground/validator/v10"
)

type ErrorResponse struct {
	Code            string `json:"code"`
	Message         string `json:"message"`
	MessageTitle    string `json:"message_title"`
	MessageSeverity string `json:"message_severity"`
}

type Response struct {
	StatusCode int             `json:"-"`
	Success    bool            `json:"success"`
	Errors     []ErrorResponse `json:"errors"`
	Data       interface{}     `json:"data"`
}

const (
	InvalidPayloadErrorCode = "InvalidPayload"
	ValidationErrorCode     = "InvalidRequest"
	UnauthorizedErrorCode   = "Unauthorized"
	InternalServerErrorCode = "InternalServerError"
)

func NewResponse(status int, data interface{}) Response {
	return Response{
		StatusCode: status,
		Success:    true,
		Data:       data,
	}
}

func NewErrorResponse(err error) Response {
	var status int
	var errs []ErrorResponse
	switch err.(type) {
	case validator.ValidationErrors:
		status = http.StatusBadRequest
		for _, e := range err.(validator.ValidationErrors) {
			errMessage := fmt.Sprintf("failure to validate field '%s', require type '%s' with tag '%s'", e.Field(), e.Kind(), e.Tag())
			errResponse := ErrorResponse{
				Code:            ValidationErrorCode,
				Message:         errMessage,
				MessageTitle:    "Validation Error",
				MessageSeverity: "error",
			}
			errs = append(errs, errResponse)
		}
	case *errors.InvalidPayloadError:
		status = http.StatusBadRequest
		errResponse := ErrorResponse{
			Code:            InvalidPayloadErrorCode,
			Message:         err.Error(),
			MessageTitle:    "Invalid Payload",
			MessageSeverity: "error",
		}
		errs = append(errs, errResponse)
	case *errors.UnauthorizedError:
		status = http.StatusUnauthorized
		errResponse := ErrorResponse{
			Code:            UnauthorizedErrorCode,
			Message:         err.Error(),
			MessageTitle:    "Unauthorized",
			MessageSeverity: "error",
		}
		errs = append(errs, errResponse)
	default:
		status = http.StatusInternalServerError
		errResponse := ErrorResponse{
			Code:            InternalServerErrorCode,
			Message:         err.Error(),
			MessageTitle:    "Internal Server Error",
			MessageSeverity: "error",
		}
		errs = append(errs, errResponse)
	}
	return Response{
		StatusCode: status,
		Success:    false,
		Errors:     errs,
	}
}
