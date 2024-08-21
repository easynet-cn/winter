package winter

import (
	"net/http"
	"strconv"
)

type BusinessError struct {
	Status  int    `json:"status"`
	Code    string `json:"code"`
	Message string `json:"message"`
}

func NewBusinessError(status int, code string, message string) *BusinessError {
	return &BusinessError{
		Status:  status,
		Code:    code,
		Message: message,
	}
}

func NewBadRequestBusinessError(message string) *BusinessError {
	return NewBusinessError(http.StatusBadRequest, strconv.Itoa(http.StatusBadRequest), message)
}

func NewNotFoundBusinessError(message string) *BusinessError {
	return NewBusinessError(http.StatusNotFound, strconv.Itoa(http.StatusNotFound), message)
}

func NewUnauthorizedBusinessError(message string) *BusinessError {
	return NewBusinessError(http.StatusUnauthorized, strconv.Itoa(http.StatusUnauthorized), message)
}

func NewForbiddenBusinessError(message string) *BusinessError {
	return NewBusinessError(http.StatusForbidden, strconv.Itoa(http.StatusForbidden), message)
}

func NewInternalServerErrorBusinessError(message string) *BusinessError {
	return NewBusinessError(http.StatusInternalServerError, strconv.Itoa(http.StatusInternalServerError), message)
}

func NewServiceUnavailableBusinessError(message string) *BusinessError {
	return NewBusinessError(http.StatusServiceUnavailable, strconv.Itoa(http.StatusServiceUnavailable), message)
}

func (m *BusinessError) Error() string {
	return m.Message
}
