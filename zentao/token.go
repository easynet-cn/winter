package zentao

type GetTokenRequest struct {
	Account  string `json:"account"`
	Password string `json:"password"`
}

type GetTokenResponse struct {
	Token string `json:"token"`
}

// 获取Token
func (s *ZentaoClient) GetToken(request *GetTokenRequest) (int, []byte, *GetTokenResponse, error) {
	if status, bytes, err := s.webClient.Post(s.url, getTokenPath, nil, request, nil); err != nil {
		return status, bytes, nil, err
	} else {
		response := &GetTokenResponse{}

		status, bytes, err := ParseResult(status, bytes, response)

		return status, bytes, response, err
	}
}
