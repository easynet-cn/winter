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
	return &BusinessError{
		Status:  http.StatusBadRequest,
		Code:    strconv.Itoa(http.StatusBadRequest),
		Message: message,
	}
}

func NewNotFoundBusinessError(message string) *BusinessError {
	return &BusinessError{
		Status:  http.StatusNotFound,
		Code:    strconv.Itoa(http.StatusNotFound),
		Message: message,
	}
}

func NewInternalServerErrorBusinessError(message string) *BusinessError {
	return &BusinessError{
		Status:  http.StatusInternalServerError,
		Code:    strconv.Itoa(http.StatusInternalServerError),
		Message: message,
	}
}

func (m *BusinessError) Error() string {
	return m.Message
}
