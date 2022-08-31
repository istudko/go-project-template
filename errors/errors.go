package errors

import "fmt"

type InvalidPayloadError struct {
	message string
}

func (e *InvalidPayloadError) Error() string {
	return e.message
}

func NewInvalidPayloadError(message string) *InvalidPayloadError {
	return &InvalidPayloadError{message: message}
}

func NewInvalidPayloadErrorf(format string, args ...interface{}) *InvalidPayloadError {
	message := fmt.Sprintf(format, args...)
	return &InvalidPayloadError{message: message}
}

type UnauthorizedError struct {
	message string
}

func (e *UnauthorizedError) Error() string {
	return e.message
}

func NewUnauthorizedError(message string) *UnauthorizedError {
	return &UnauthorizedError{message: message}
}

func NewUnauthorizedErrorf(format string, args ...interface{}) *UnauthorizedError {
	message := fmt.Sprintf(format, args...)
	return &UnauthorizedError{message: message}
}

type DBError struct {
	message string
}

func (e *DBError) Error() string {
	return e.message
}

func NewDBError(message string) *DBError {
	return &DBError{message: message}
}

func NewDBErrorf(format string, args ...interface{}) *DBError {
	message := fmt.Sprintf(format, args...)
	return &DBError{message: message}
}

type GenericError struct {
	message string
}

func (e *GenericError) Error() string {
	return e.message
}

func NewGenericError(message string) *GenericError {
	return &GenericError{message: message}
}

func NewGenericErrorf(format string, args ...interface{}) *GenericError {
	message := fmt.Sprintf(format, args...)
	return &GenericError{message: message}
}
