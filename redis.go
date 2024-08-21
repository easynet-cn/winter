package winter

import (
	"fmt"

	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
)

const (
	defaultRedisClientKey = "default"
)

type Redis struct {
	config       *viper.Viper
	redisClients map[string]*redis.Client
}

func NewRedis(config *viper.Viper) *Redis {
	return &Redis{
		config:       config,
		redisClients: make(map[string]*redis.Client),
	}
}

func (m *Redis) Init() {
	for k := range m.config.GetStringMap("spring.redis") {
		redisClient := redis.NewClient(&redis.Options{
			Addr:     fmt.Sprintf("%s:%d", m.config.GetString(fmt.Sprintf("spring.redis.%s.host", k)), m.config.GetInt(fmt.Sprintf("spring.redis.%s.port", k))),
			Password: m.config.GetString(fmt.Sprintf("spring.redis.%s.password", k)),
		})

		m.redisClients[k] = redisClient
	}
}

func (m *Redis) GetRedisClients() map[string]*redis.Client {
	return m.redisClients
}

func (m *Redis) GetRedisClient() *redis.Client {
	return m.redisClients[defaultRedisClientKey]
}
