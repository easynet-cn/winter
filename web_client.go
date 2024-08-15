package winter

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"strings"
)

type WebClient struct {
	Client *http.Client
}

func NewWebClient(client *http.Client) *WebClient {
	return &WebClient{Client: client}
}

func (m *WebClient) Get(baseUrl string, path string, urlValues url.Values) (int, []byte, error) {
	_, code, bytees, err := m.Do("GET", "application/json;charset=UTF-8", baseUrl, path, urlValues, nil)

	return code, bytees, err
}

func (m *WebClient) Post(baseUrl string, path string, urlValues url.Values, bodyValue any) (int, []byte, error) {
	_, code, bytees, err := m.Do("POST", "application/json;charset=UTF-8", baseUrl, path, urlValues, bodyValue)

	return code, bytees, err
}

func (m *WebClient) Put(baseUrl string, path string, urlValues url.Values, bodyValue any) (int, []byte, error) {
	_, code, bytees, err := m.Do("PUT", "application/json;charset=UTF-8", baseUrl, path, urlValues, bodyValue)

	return code, bytees, err
}

func (m *WebClient) Delete(baseUrl string, path string, urlValues url.Values, bodyValue any) (int, []byte, error) {
	_, code, bytees, err := m.Do("DELETE", "application/json;charset=UTF-8", baseUrl, path, urlValues, bodyValue)

	return code, bytees, err
}

func (m *WebClient) Do(method string, contentType string, baseUrl string, path string, urlValues url.Values, bodyValue any) (*http.Response, int, []byte, error) {
	bodyValueBytes, err1 := json.Marshal(bodyValue)

	if err1 != nil {
		return nil, 0, nil, err1
	}

	if req, err := http.NewRequest(method, m.url(baseUrl, path, urlValues), bytes.NewReader(bodyValueBytes)); err != nil {
		return nil, 0, nil, err
	} else {
		return m.doRequest(req, contentType)
	}
}

func (m *WebClient) url(baseUrl string, path string, urlValues url.Values) string {
	result, _ := url.JoinPath(baseUrl, path)

	sb := new(strings.Builder)

	sb.WriteString(result)

	if len(urlValues) > 0 {
		sb.WriteString("?")
		sb.WriteString(urlValues.Encode())
	}

	return sb.String()

}

func (m *WebClient) doRequest(req *http.Request, contentType string) (*http.Response, int, []byte, error) {
	req.Header.Set("Content-Type", contentType)

	if res, err := m.Client.Do(req); err != nil {
		return res, res.StatusCode, nil, err
	} else {
		if resBytes, err := io.ReadAll(res.Body); err != nil {
			return res, res.StatusCode, nil, err
		} else {
			return res, res.StatusCode, resBytes, nil
		}
	}
}
