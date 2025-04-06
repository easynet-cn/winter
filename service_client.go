package winter

import (
	"net/http"
	"net/url"
)

type Banlancer interface {
	GetUri(serviceName string) (string, error)
}

type ServiceClient struct {
	webClient *WebClient
	banlancer Banlancer
}

func NewServiceClient(
	webClient *WebClient,
	banlancer Banlancer) *ServiceClient {

	return &ServiceClient{
		webClient: webClient,
		banlancer: banlancer,
	}
}

func (m *ServiceClient) Get(
	serviceName string,
	path string,
	urlValues url.Values,
	encodingFunc EncodingFunc,
	funcs ...RequestHeaderFunc,
) (int, []byte, error) {
	_, code, bytees, err := m.Do(serviceName, "POST", path, urlValues, nil, encodingFunc, funcs...)

	return code, bytees, err
}

func (m *ServiceClient) Post(
	serviceName string,
	path string,
	urlValues url.Values,
	bodyValue any,
	encodingFunc EncodingFunc,
	funcs ...RequestHeaderFunc,
) (int, []byte, error) {
	_, code, bytees, err := m.Do(serviceName, "POST", path, urlValues, bodyValue, encodingFunc, funcs...)

	return code, bytees, err
}

func (m *ServiceClient) Put(
	serviceName string,
	path string,
	urlValues url.Values,
	bodyValue any,
	encodingFunc EncodingFunc,
	funcs ...RequestHeaderFunc,
) (int, []byte, error) {
	_, code, bytees, err := m.Do(serviceName, "PUT", path, urlValues, bodyValue, encodingFunc, funcs...)

	return code, bytees, err
}

func (m *ServiceClient) Delete(
	serviceName string,
	path string,
	urlValues url.Values,
	bodyValue any,
	encodingFunc EncodingFunc,
	funcs ...RequestHeaderFunc,
) (int, []byte, error) {
	_, code, bytees, err := m.Do(serviceName, "DELETE", path, urlValues, bodyValue, encodingFunc, funcs...)

	return code, bytees, err
}

func (m *ServiceClient) Do(
	serviceName string,
	method string,
	path string,
	urlValues url.Values,
	bodyValue any,
	encodingFunc EncodingFunc,
	funcs ...RequestHeaderFunc,
) (*http.Response, int, []byte, error) {
	serviceUrl, err := m.banlancer.GetUri(serviceName)

	if err != nil {
		return nil, 0, nil, err
	}

	return m.webClient.Do(method, serviceUrl, path, urlValues, bodyValue, encodingFunc, funcs...)
}
