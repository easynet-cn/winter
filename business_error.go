package winter

type BusinessError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

func (m *BusinessError) Error() string {
	return m.Message
}
