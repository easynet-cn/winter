package winter

import (
	"bytes"
	"fmt"
	"path"
	"path/filepath"
	"strings"

	"github.com/hashicorp/consul/api"
	"github.com/nacos-group/nacos-sdk-go/v2/common/file"
	"github.com/spf13/viper"
)

const (
	HealthCheckUrl = "/actuator/health"
)

type Consul struct {
	config         *viper.Viper
	defaultClient  *api.Client
	serviceClients map[string]*api.Client
}

func NewConsul() *Consul {
	config := viper.New()

	config.SetConfigType("yml")

	return &Consul{
		config:         viper.New(),
		serviceClients: make(map[string]*api.Client),
	}
}

func (m *Consul) GetConfig() *viper.Viper {
	return m.config
}

func (m *Consul) GetDefaultClient() *api.Client {
	return m.defaultClient
}

func (m *Consul) GetServiceClients() map[string]*api.Client {
	return m.serviceClients
}

func (m *Consul) Init() {
	if flagSet, err := getFlatSet(); err != nil {
		panic(err)
	} else if err := m.config.BindPFlags(flagSet); err != nil {
		panic(err)
	}

	m.setLocalConfigPathAndName()

	if err := m.config.ReadInConfig(); err != nil {
		panic(err)
	}

	if m.config.InConfig("consul") {
		config := api.DefaultConfig()

		config.Address = fmt.Sprintf("http://%s", m.config.GetString("consul.host"))

		if client, err := api.NewClient(config); err != nil {
			panic(err)
		} else {
			m.defaultClient = client
		}

		configKey := fmt.Sprintf("%s/%s,%s/data", m.config.GetString("consul.config.prefix"), m.config.GetString("spring.application.name"), m.config.GetString("spring.application.profile.active"))

		if kv, _, err := m.defaultClient.KV().Get(configKey, nil); err != nil {
			panic(err)
		} else {
			remoteConfig := viper.New()

			remoteConfig.SetConfigType("yaml")

			if err := remoteConfig.ReadConfig(bytes.NewBuffer(kv.Value)); err != nil {
				panic(err)
			} else {
				keys := remoteConfig.AllKeys()

				for _, key := range keys {
					m.config.Set(key, remoteConfig.Get(key))
				}
			}

			m.registerServices()
		}
	}
}

func (m *Consul) setLocalConfigPathAndName() {
	configPath := "./"
	configName := "application"

	configLocation := m.config.GetString("spring.config.location")

	if configLocation != "" {
		if file.IsExistFile(configLocation) {
			configPath = filepath.Dir(configLocation)
			configName = strings.TrimSuffix(filepath.Base(configLocation), filepath.Ext(configLocation))
		}
	} else {
		activeProfile := m.config.GetString("spring.profiles.active")

		if activeProfile != "" && file.IsExistFile(path.Join("./", fmt.Sprintf("application-%s.yml", activeProfile))) {
			configName = fmt.Sprintf("application-%s", activeProfile)
		}
	}

	m.config.AddConfigPath(configPath)
	m.config.SetConfigName(configName)
	m.config.SetConfigType("yml")
}

func (m *Consul) registerServices() error {
	ip := m.config.GetString("consul.discovery.ip-address")

	if ip == "" {
		ip = ExternalIP().String()
	}

	registration := new(api.AgentServiceRegistration)

	registration.ID = fmt.Sprintf("%s-%s-%d", m.config.GetString("spring.application.name"), ip, m.config.GetInt("server.port"))
	registration.Name = m.config.GetString("spring.application.name")
	registration.Port = m.config.GetInt("server.port")
	registration.Tags = []string{}
	registration.Address = ip
	registration.Check = &api.AgentServiceCheck{
		HTTP:                           fmt.Sprintf("http://%s:%d%s", registration.Address, m.config.GetInt("server.port"), HealthCheckUrl),
		Timeout:                        "60s",
		Interval:                       "3s",
		DeregisterCriticalServiceAfter: "30s",
	}

	return m.defaultClient.Agent().ServiceRegister(registration)
}
