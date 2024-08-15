package winter

import (
	"fmt"

	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
)

type Redis struct {
	redisClients map[string]*redis.Client
}

func NewRedis() *Redis {
	return &Redis{
		redisClients: make(map[string]*redis.Client),
	}
}

func (m *Redis) Init(viper *viper.Viper) {
	for k := range viper.GetStringMap("spring.redis") {
		redisClient := redis.NewClient(&redis.Options{
			Addr:     fmt.Sprintf("%s:%d", viper.GetString("spring.redis.host"), viper.GetInt("spring.redis.port")),
			Password: viper.GetString("spring.redis.password"),
		})

		m.redisClients[k] = redisClient
	}
}

func (m *Redis) GetRedisClients() map[string]*redis.Client {
	return m.redisClients
}
