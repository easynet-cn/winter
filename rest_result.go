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
	Error   string `json:"error"`
}

func NewRestResult(
	status int,
	code string,
	data any,
	message string,
	err string,
) *RestResult {
	return &RestResult{
		Status:  status,
		Code:    code,
		Data:    data,
		Message: message,
		Error:   err,
	}
}

func NewSuccessRestResult(data any, message string) *RestResult {
	return &RestResult{
		Status:  http.StatusOK,
		Code:    strconv.Itoa(http.StatusOK),
		Data:    data,
		Message: message,
	}
}

func NewBadRequestRestResult(data any, message string, err string) *RestResult {
	return &RestResult{
		Status:  http.StatusBadRequest,
		Code:    strconv.Itoa(http.StatusBadRequest),
		Data:    data,
		Message: message,
		Error:   err,
	}
}

func NewNotFoundRestResult(data any, message string, err string) *RestResult {
	return &RestResult{
		Status:  http.StatusNotFound,
		Code:    strconv.Itoa(http.StatusNotFound),
		Data:    data,
		Message: message,
		Error:   err,
	}
}

func NewInternalServerErrorRestResult(data any, message string, err string) *RestResult {
	return &RestResult{
		Status:  http.StatusInternalServerError,
		Code:    strconv.Itoa(http.StatusInternalServerError),
		Data:    data,
		Message: message,
		Error:   err,
	}
}
