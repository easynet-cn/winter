package winter

import (
	"fmt"

	"github.com/spf13/viper"
	"xorm.io/xorm"
)

type Database struct {
	dbs map[string]*xorm.Engine
}

func NewDatabase() *Database {
	return &Database{
		dbs: make(map[string]*xorm.Engine),
	}
}

func (m *Database) Init(config *viper.Viper) {
	dbConfigs := config.GetStringMap("spring.datasources")

	for k := range dbConfigs {
		dbType := config.GetString(fmt.Sprintf("spring.datasources.%s.type", k))

		if dbType == "" {
			dbType = "mysql"
		}

		if engine, err := xorm.NewEngine(dbType, config.GetString(fmt.Sprintf("spring.datasources.%s.url", k))); err != nil {
			panic(fmt.Errorf("连接数据库失败：%w", err))
		} else {
			engine.SetMaxOpenConns(config.GetInt(fmt.Sprintf("spring.datasources.%s.maxOpenConns", k)))
			engine.SetMaxIdleConns(config.GetInt(fmt.Sprintf("spring.datasources.%s.maxIdleConns", k)))

			m.dbs[k] = engine
		}
	}

}

func (m *Database) GetDatabases() map[string]*xorm.Engine {
	return m.dbs
}
