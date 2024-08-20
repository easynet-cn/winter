package winter

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
