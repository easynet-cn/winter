package winter

import (
	"net/http"
	"strconv"
)

type RestResult struct {
	Status  int    `json:"status"`
	Code    string `json:"code"`
	Data    any    `json:"data"`
	Message string `json:"message"`
}

func NewRestResult(
	status int,
	code string,
	data any,
	message string,
) *RestResult {
	return &RestResult{
		Status:  status,
		Code:    code,
		Data:    data,
		Message: message,
	}
}

func NewSuccessRestResult(data any, message string) *RestResult {
	return NewRestResult(http.StatusOK, strconv.Itoa(http.StatusOK), data, message)
}

func NewBadRequestRestResult(data any, message string) *RestResult {
	return NewRestResult(http.StatusBadRequest, strconv.Itoa(http.StatusBadRequest), data, message)
}

func NewNotFoundRestResult(data any, message string) *RestResult {
	return NewRestResult(http.StatusNotFound, strconv.Itoa(http.StatusNotFound), data, message)
}

func NewUnauthorizedRestResult(data any, message string) *RestResult {
	return NewRestResult(http.StatusUnauthorized, strconv.Itoa(http.StatusUnauthorized), data, message)
}

func NewForbiddenRestResult(data any, message string) *RestResult {
	return NewRestResult(http.StatusForbidden, strconv.Itoa(http.StatusForbidden), data, message)
}

func NewInternalServerErrorRestResult(data any, message string) *RestResult {
	return NewRestResult(http.StatusInternalServerError, strconv.Itoa(http.StatusInternalServerError), data, message)
}

func NewServiceUnavailableRestResult(data any, message string) *RestResult {
	return NewRestResult(http.StatusServiceUnavailable, strconv.Itoa(http.StatusServiceUnavailable), data, message)
}
