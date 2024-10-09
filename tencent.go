package winter

import (
	"fmt"

	"github.com/spf13/viper"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	ess "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/ess/v20201111"
)

type TencentEss struct {
	config     *viper.Viper
	essClients map[string]*ess.Client
}

func NewTencentEss(config *viper.Viper) *TencentEss {
	return &TencentEss{
		config:     config,
		essClients: make(map[string]*ess.Client),
	}
}

func (m *TencentEss) Init() {
	for k := range m.config.GetStringMap("tencent.ess") {
		region := m.config.GetString(fmt.Sprintf("tencent.ess.%s.region", k))
		secretId := m.config.GetString(fmt.Sprintf("tencent.ess.%s.secret-id", k))
		secretKey := m.config.GetString(fmt.Sprintf("tencent.ess.%s.secret-key", k))

		credential := common.NewCredential(secretId, secretKey)
		clientProfile := profile.NewClientProfile()

		if essClient, err := ess.NewClient(credential, region, clientProfile); err != nil {
			panic(fmt.Errorf("连接tencent ess失败：%w", err))
		} else {
			m.essClients[k] = essClient
		}
	}
}

func (m *TencentEss) GetEssClients() map[string]*ess.Client {
	return m.essClients
}

func (m *TencentEss) GetEssClient() *ess.Client {
	return m.essClients[defaultRedisClientKey]
}
