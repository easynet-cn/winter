package winter

import (
	"fmt"

	"github.com/spf13/viper"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	ess "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/ess/v20201111"
)

type TencentEssClient struct {
	essClient     *ess.Client
	essFileClient *ess.Client
}

func (m *TencentEssClient) GetEssClient() *ess.Client {
	return m.essClient
}

func (m *TencentEssClient) GetEssFileClient() *ess.Client {
	return m.essFileClient
}

type TencentEss struct {
	config     *viper.Viper
	essClients map[string]*TencentEssClient
}

func NewTencentEss(config *viper.Viper) *TencentEss {
	return &TencentEss{
		config:     config,
		essClients: make(map[string]*TencentEssClient),
	}
}

func (m *TencentEss) Init() {
	for k := range m.config.GetStringMap("tencent.ess") {
		region := m.config.GetString(fmt.Sprintf("tencent.ess.%s.region", k))
		secretId := m.config.GetString(fmt.Sprintf("tencent.ess.%s.secret-id", k))
		secretKey := m.config.GetString(fmt.Sprintf("tencent.ess.%s.secret-key", k))
		endpoint := m.config.GetString(fmt.Sprintf("tencent.ess.%s.endpoint", k))
		fileEndpoint := m.config.GetString(fmt.Sprintf("tencent.ess.%s.file-endpoint", k))

		tencentEssClient := &TencentEssClient{}

		m.essClients[k] = tencentEssClient

		credential := common.NewCredential(secretId, secretKey)

		clientProfile := profile.NewClientProfile()

		if endpoint != "" {
			clientProfile.HttpProfile.Endpoint = endpoint
		}

		if essClient, err := ess.NewClient(credential, region, clientProfile); err != nil {
			panic(fmt.Errorf("连接tencent ess失败：%w", err))
		} else {
			m.essClients[k].essClient = essClient
		}

		fileClientProfile := profile.NewClientProfile()

		if fileEndpoint != "" {
			fileClientProfile.HttpProfile.Endpoint = fileEndpoint
		}

		if essClient, err := ess.NewClient(credential, region, fileClientProfile); err != nil {
			panic(fmt.Errorf("连接tencent ess失败：%w", err))
		} else {
			m.essClients[k].essFileClient = essClient
		}
	}
}

func (m *TencentEss) GetEssClients() map[string]*TencentEssClient {
	return m.essClients
}

func (m *TencentEss) GetEssClient() *ess.Client {
	return m.essClients[defaultMapKey].GetEssClient()
}

func (m *TencentEss) GetEssFileClient() *ess.Client {
	return m.essClients[defaultMapKey].GetEssFileClient()
}
