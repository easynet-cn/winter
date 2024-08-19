package winter

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/nacos-group/nacos-sdk-go/v2/clients/naming_client"
	"github.com/nacos-group/nacos-sdk-go/v2/vo"
)

type Banlancer interface {
	GetUri(serviceName string) (string, error)
}

type NacosBanlancer struct {
	namingClient naming_client.INamingClient
}

func (m *NacosBanlancer) GetUri(serviceName string) (string, error) {
	instance, err := m.namingClient.SelectOneHealthyInstance(vo.SelectOneHealthInstanceParam{ServiceName: serviceName})

	if err != nil {
		return "", err
	}

	if instance == nil {
		return "", fmt.Errorf("服务（%s）没有可用的实例", serviceName)
	}

	return fmt.Sprintf("http://%s:%d", instance.Ip, instance.Port), nil
}

type ServiceClient struct {
	namingClient naming_client.INamingClient
	webClient    *WebClient
	banlancer    Banlancer
}

func NewServiceClient(
	namingClient naming_client.INamingClient,
	webClient *WebClient,
	banlancer Banlancer) *ServiceClient {

	if banlancer == nil {
		banlancer = &NacosBanlancer{namingClient: namingClient}
	}

	return &ServiceClient{
		namingClient: namingClient,
		webClient:    webClient,
		banlancer:    banlancer,
	}
}

func (m *ServiceClient) Get(serviceName string, path string, urlValues url.Values) (int, []byte, error) {
	_, code, bytees, err := m.Do(serviceName, "POST", "application/json;charset=UTF-8", path, urlValues, nil)

	return code, bytees, err
}

func (m *ServiceClient) Post(serviceName string, path string, urlValues url.Values, bodyValue any) (int, []byte, error) {
	_, code, bytees, err := m.Do(serviceName, "POST", "application/json;charset=UTF-8", path, urlValues, bodyValue)

	return code, bytees, err
}

func (m *ServiceClient) Put(serviceName string, path string, urlValues url.Values, bodyValue any) (int, []byte, error) {
	_, code, bytees, err := m.Do(serviceName, "PUT", "application/json;charset=UTF-8", path, urlValues, bodyValue)

	return code, bytees, err
}

func (m *ServiceClient) Delete(serviceName string, path string, urlValues url.Values, bodyValue any) (int, []byte, error) {
	_, code, bytees, err := m.Do(serviceName, "DELETE", "application/json;charset=UTF-8", path, urlValues, bodyValue)

	return code, bytees, err
}

func (m *ServiceClient) Do(
	serviceName string,
	method string,
	contentType string,
	path string, urlValues url.Values,
	bodyValue any) (*http.Response, int, []byte, error) {
	serviceUrl, err := m.banlancer.GetUri(serviceName)

	if err != nil {
		return nil, 0, nil, err
	}

	return m.webClient.Do(method, contentType, serviceUrl, path, urlValues, bodyValue)
}
