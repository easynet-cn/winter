package winter

import (
	"bytes"
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

type RequestHeaderFunc func(http.Header)
type EncodingFunc func(v any) ([]byte, error)

func (m *WebClient) Get(
	baseUrl string,
	path string,
	urlValues url.Values,
	encodingFunc EncodingFunc,
	funcs ...RequestHeaderFunc,
) (int, []byte, error) {
	_, code, bytes, err := m.Do("GET", baseUrl, path, urlValues, nil, encodingFunc, funcs...)

	return code, bytes, err
}

func (m *WebClient) Post(
	baseUrl string,
	path string,
	urlValues url.Values,
	bodyValue any,
	encodingFunc EncodingFunc,
	funcs ...RequestHeaderFunc,
) (int, []byte, error) {
	_, code, bytes, err := m.Do("POST", baseUrl, path, urlValues, bodyValue, encodingFunc, funcs...)

	return code, bytes, err
}

func (m *WebClient) Put(
	baseUrl string,
	path string,
	urlValues url.Values,
	bodyValue any,
	encodingFunc EncodingFunc,
	funcs ...RequestHeaderFunc,
) (int, []byte, error) {
	_, code, bytes, err := m.Do("PUT", baseUrl, path, urlValues, bodyValue, encodingFunc, funcs...)

	return code, bytes, err
}

func (m *WebClient) Delete(
	baseUrl string,
	path string,
	urlValues url.Values,
	bodyValue any,
	encodingFunc EncodingFunc,
	funcs ...RequestHeaderFunc,
) (int, []byte, error) {
	_, code, bytes, err := m.Do("DELETE", baseUrl, path, urlValues, bodyValue, encodingFunc, funcs...)

	return code, bytes, err
}

func (m *WebClient) Do(
	method string,
	baseUrl string,
	path string,
	urlValues url.Values,
	bodyValue any,
	encodingFunc EncodingFunc,
	funcs ...RequestHeaderFunc,
) (*http.Response, int, []byte, error) {
	var reader io.Reader

	if bodyValue != nil && encodingFunc != nil {
		if bodyBytes, err := encodingFunc(bodyValue); err != nil {
			return nil, 0, nil, err
		} else {
			reader = bytes.NewReader(bodyBytes)
		}
	}

	if req, err := http.NewRequest(method, m.url(baseUrl, path, urlValues), reader); err != nil {
		return nil, 0, nil, err
	} else {
		for _, f := range funcs {
			f(req.Header)
		}

		return m.DoRequest(req)
	}
}

func (m *WebClient) DoRequest(req *http.Request) (*http.Response, int, []byte, error) {
	res, err := m.Client.Do(req)

	statusCode := 0

	if res != nil {
		statusCode = res.StatusCode
	}

	if err != nil {
		return res, statusCode, nil, err
	}

	if res == nil || res.Body == nil {
		return res, statusCode, nil, nil
	} else {
		defer res.Body.Close()
	}

	if bytes, err := io.ReadAll(res.Body); err != nil {
		return res, statusCode, nil, err
	} else {
		return res, statusCode, bytes, nil
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
