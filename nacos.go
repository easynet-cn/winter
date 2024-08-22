package winter

import (
	"bytes"
	"fmt"
	"log"
	"maps"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/nacos-group/nacos-sdk-go/v2/clients"
	"github.com/nacos-group/nacos-sdk-go/v2/clients/config_client"
	"github.com/nacos-group/nacos-sdk-go/v2/clients/naming_client"
	"github.com/nacos-group/nacos-sdk-go/v2/common/constant"
	"github.com/nacos-group/nacos-sdk-go/v2/common/file"
	"github.com/nacos-group/nacos-sdk-go/v2/vo"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

const (
	defaultTimeoutMs = uint64(10000)
	defaultLogLevel  = "warn"
	defaultWeight    = 1
)

type Nacos struct {
	config               *viper.Viper
	defaultNameingClient naming_client.INamingClient
	serviceNamingClients map[string]naming_client.INamingClient
	metadata             map[string]string
}

func NewNacos(
	metadata map[string]string,
) *Nacos {
	config := viper.New()

	config.SetConfigType("yml")

	if metadata == nil {
		metadata = make(map[string]string)
	}

	return &Nacos{
		config:               viper.New(),
		metadata:             metadata,
		serviceNamingClients: make(map[string]naming_client.INamingClient),
	}
}

func (m *Nacos) GetConfig() *viper.Viper {
	return m.config
}
func (m *Nacos) GetDefaultNameingClient() naming_client.INamingClient {
	return m.defaultNameingClient
}

func (m *Nacos) GetServiceNamingClient(configKey string) naming_client.INamingClient {
	return m.serviceNamingClients[configKey]
}

func (m *Nacos) GetServiceNamingClients() map[string]naming_client.INamingClient {
	return m.serviceNamingClients
}

func (m *Nacos) GetMetadata() map[string]string {
	return m.metadata
}

func (m *Nacos) BuildDefaultServiceClient(webClient *WebClient) *ServiceClient {
	return NewServiceClient(m.defaultNameingClient, webClient, nil)
}

func (m *Nacos) BuildServiceClientWithConfigKey(webClient *WebClient, configKey string) *ServiceClient {
	return NewServiceClient(m.serviceNamingClients[configKey], webClient, nil)
}

func (m *Nacos) Init() {
	if flagSet, err := getFlatSet(); err != nil {
		panic(err)
	} else if err := m.config.BindPFlags(flagSet); err != nil {
		panic(err)
	}

	m.setLocalConfigPathAndName()

	if err := m.config.ReadInConfig(); err != nil {
		panic(err)
	}

	serverConfig, clientConfig := getNacosConfig(m.config)

	if configClient, err := clients.NewConfigClient(
		vo.NacosClientParam{
			ClientConfig:  &clientConfig,
			ServerConfigs: serverConfig,
		},
	); err != nil {
		panic(err)
	} else if remoteConfig, err := m.getRemoteConfig(configClient); err != nil {
		panic(err)
	} else {
		keys := remoteConfig.AllKeys()

		for _, key := range keys {
			m.config.Set(key, remoteConfig.Get(key))
		}
	}

	if namingClient, err := clients.NewNamingClient(
		vo.NacosClientParam{
			ClientConfig:  &clientConfig,
			ServerConfigs: serverConfig,
		},
	); err != nil {
		panic(err)
	} else if success, err := m.registerService(namingClient); err != nil {
		panic(err)
	} else if !success {
		log.Fatalf("Failed to register service")
	} else {
		m.defaultNameingClient = namingClient
	}

	m.registerNacoseServices()
}

func getFlatSet() (*pflag.FlagSet, error) {
	flagSet := pflag.NewFlagSet("system", pflag.ContinueOnError)

	flagSet.ParseErrorsWhitelist = pflag.ParseErrorsWhitelist{UnknownFlags: true}
	flagSet.SortFlags = false

	flagSet.String("spring.config.location", "", "config location")
	flagSet.String("spring.profiles.active", "dev", "active profile")
	flagSet.String("nacos.namespace", "", "nacos namespace")
	flagSet.String("nacos.host", "127.0.0.1", "nacos host")
	flagSet.Int("nacos.port", 8848, "nacos port")
	flagSet.String("nacos.context-path", "nacos", "nacos context path")
	flagSet.String("nacos.group", "group", "nacos group")
	flagSet.String("nacos.username", "", "nacos username")
	flagSet.String("nacos.password", "", "nacos password")
	flagSet.Int("server.port", 6103, "server port")

	if err := flagSet.Parse(os.Args[1:]); err != nil {
		return nil, err
	}

	return flagSet, nil
}

func getNacosConfig(config *viper.Viper) ([]constant.ServerConfig, constant.ClientConfig) {
	serverConfig := []constant.ServerConfig{
		{
			IpAddr:      config.GetString("nacos.host"),
			Port:        config.GetUint64("nacos.port"),
			ContextPath: config.GetString("nacos.context-path"),
		},
	}

	timeoutMs := config.GetUint64("nacos.timeout-ms")

	if timeoutMs <= 0 {
		timeoutMs = defaultTimeoutMs
	}

	notLoadCacheAtStart := config.GetBool("nacos.not-load-cache-at-start")
	logLevel := config.GetString("nacos.log-level")

	if logLevel == "" {
		logLevel = defaultLogLevel
	}

	clientConfig := constant.ClientConfig{
		NamespaceId:         config.GetString("nacos.namespace"),
		Username:            config.GetString("nacos.username"),
		Password:            config.GetString("nacos.password"),
		TimeoutMs:           timeoutMs,
		NotLoadCacheAtStart: notLoadCacheAtStart,
		LogLevel:            logLevel,
	}

	return serverConfig, clientConfig
}

func (m *Nacos) setLocalConfigPathAndName() {
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

func (m *Nacos) getRemoteConfig(configClient config_client.IConfigClient) (*viper.Viper, error) {
	remoteConfig := viper.New()

	remoteConfig.SetConfigType("yml")

	if content, err := configClient.GetConfig(vo.ConfigParam{
		DataId: fmt.Sprintf("%s-%s.yml", m.config.GetString("spring.application.name"), m.config.GetString("spring.profiles.active")),
		Group:  m.config.GetString("nacos.group")}); err != nil {

		return nil, err
	} else if err := remoteConfig.ReadConfig(bytes.NewBuffer([]byte(content))); err != nil {
		return nil, err
	}

	return remoteConfig, nil
}

func (m *Nacos) registerService(namingClient naming_client.INamingClient) (bool, error) {
	serviceName := m.config.GetString("nacos.service-name")

	if serviceName == "" {
		serviceName = m.config.GetString("spring.application.name")
	}

	weight := m.config.GetFloat64("nacos.weight")

	if weight <= 0 {
		weight = defaultWeight
	}

	metadata := m.config.GetStringMapString("nacos.metadata")

	maps.Copy(m.metadata, metadata)

	return namingClient.RegisterInstance(vo.RegisterInstanceParam{
		Ip:          ExternalIP().String(),
		Port:        m.config.GetUint64("server.port"),
		ServiceName: serviceName,
		Weight:      weight,
		Enable:      true,
		Healthy:     true,
		Ephemeral:   true,
		Metadata:    m.metadata,
	})
}

func (m *Nacos) registerNacoseServices() {
	for k := range m.config.GetStringMap("nacos.services") {
		serverConfig := []constant.ServerConfig{
			{
				IpAddr:      m.config.GetString(fmt.Sprintf("nacos.services.%s.host", k)),
				Port:        m.config.GetUint64(fmt.Sprintf("nacos.services.%s.port", k)),
				ContextPath: m.config.GetString(fmt.Sprintf("nacos.services.%s.context-path", k)),
			},
		}

		timeoutMs := m.config.GetUint64(fmt.Sprintf("nacos.services.%s.timeout-ms", k))

		if timeoutMs <= 0 {
			timeoutMs = defaultTimeoutMs
		}

		notLoadCacheAtStart := m.config.GetBool(fmt.Sprintf("nacos.services.%s.not-load-cache-at-start", k))
		logLevel := m.config.GetString(fmt.Sprintf("nacos.services.%s.log-level", k))

		if logLevel == "" {
			logLevel = defaultLogLevel
		}

		clientConfig := constant.ClientConfig{
			NamespaceId:         m.config.GetString(fmt.Sprintf("nacos.services.%s.namespace", k)),
			Username:            m.config.GetString(fmt.Sprintf("nacos.services.%s.username", k)),
			Password:            m.config.GetString(fmt.Sprintf("nacos.services.%s.password", k)),
			TimeoutMs:           timeoutMs,
			NotLoadCacheAtStart: notLoadCacheAtStart,
			LogLevel:            logLevel,
		}

		namingClient, err := clients.NewNamingClient(
			vo.NacosClientParam{
				ClientConfig:  &clientConfig,
				ServerConfigs: serverConfig,
			},
		)

		if err != nil {
			panic(fmt.Errorf("初始化nacos服务发现客户端失败: %w", err))
		}

		serviceName := m.config.GetString(fmt.Sprintf("nacos.services.%s.service-name", k))

		if serviceName == "" {
			serviceName = m.config.GetString("spring.application.name")
		}

		weight := m.config.GetFloat64(fmt.Sprintf("nacos.services.%s.weight", k))

		if weight <= 0 {
			weight = defaultWeight
		}

		metadata := m.config.GetStringMapString(fmt.Sprintf("nacos.services.%s.metadata", k))

		maps.Copy(m.metadata, metadata)

		success, err := namingClient.RegisterInstance(vo.RegisterInstanceParam{
			Ip:          ExternalIP().String(),
			Port:        m.config.GetUint64("server.port"),
			ServiceName: serviceName,
			Weight:      weight,
			Enable:      true,
			Healthy:     true,
			Ephemeral:   true,
			Metadata:    m.metadata,
		})

		if !success || err != nil {
			panic(fmt.Errorf("初始化nacos服务注册失败: %w", err))
		}

		m.serviceNamingClients[k] = namingClient
	}
}
