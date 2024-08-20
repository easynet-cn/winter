package winter

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type ApplicationFunc func()

type Application struct {
	engine *gin.Engine
	nacos  *Nacos
}

func NewApplication(
	nacos *Nacos,
) *Application {
	gin.SetMode(gin.ReleaseMode)

	application := &Application{
		engine: gin.Default(),
		nacos:  nacos,
	}

	return application
}

func (m *Application) GetEngine() *gin.Engine {
	return m.engine
}

func (m *Application) Run(funcs ...ApplicationFunc) {
	for _, f := range funcs {
		if f != nil {
			f()
		}
	}

	m.engine.Run(fmt.Sprintf(":%d", m.nacos.GetConfig().GetInt("server.port")))
}
