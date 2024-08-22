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
	systemMiddleware *SystemMiddleware,
) *Application {
	gin.SetMode(gin.ReleaseMode)

	application := &Application{
		engine: gin.Default(),
		nacos:  nacos,
	}

	RegisterDefaultMiddleware(application.engine, systemMiddleware)

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

	if err := m.engine.Run(fmt.Sprintf(":%d", m.nacos.GetConfig().GetInt("server.port"))); err != nil {
		panic(err)
	}
}
