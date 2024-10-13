package winter

import (
	"fmt"

	"github.com/spf13/viper"
	"xorm.io/xorm"
)

type Database struct {
	config *viper.Viper
	dbs    map[string]*xorm.Engine
}

func NewDatabase(config *viper.Viper) *Database {
	return &Database{
		config: config,
		dbs:    make(map[string]*xorm.Engine),
	}
}

func (m *Database) Init() {
	dbConfigs := m.config.GetStringMap("spring.datasources")

	for k := range dbConfigs {
		dbType := m.config.GetString(fmt.Sprintf("spring.datasources.%s.type", k))

		if dbType == "" {
			dbType = "mysql"
		}

		if engine, err := xorm.NewEngine(dbType, m.config.GetString(fmt.Sprintf("spring.datasources.%s.url", k))); err != nil {
			panic(fmt.Errorf("连接数据库失败：%w", err))
		} else {
			engine.SetMaxOpenConns(m.config.GetInt(fmt.Sprintf("spring.datasources.%s.maxOpenConns", k)))
			engine.SetMaxIdleConns(m.config.GetInt(fmt.Sprintf("spring.datasources.%s.maxIdleConns", k)))

			m.dbs[k] = engine
		}
	}

}

func (m *Database) GetDatabases() map[string]*xorm.Engine {
	return m.dbs
}

func (m *Database) GetDatabase(key string) *xorm.Engine {
	return m.dbs[key]
}
