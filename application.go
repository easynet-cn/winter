package winter

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type ApplicationFunc func()

type Application struct {
	engine   *gin.Engine
	nacos    *Nacos
	logger   *zap.Logger
	database *Database
	redis    *Redis
}

func NewApplication(
	metadata map[string]string,
	version string,
	syncDBFunc func() error,
) *Application {
	nacos := NewNacos(metadata)

	nacos.Init()

	logger := NewLogger(nacos.GetConfig())

	gin.SetMode(gin.ReleaseMode)

	engine := gin.Default()

	application := &Application{
		engine: engine,
		nacos:  nacos,
		logger: logger,
	}

	application.nacos.Init()
	application.logger = NewLogger(nacos.GetConfig())

	database := NewDatabase(nacos.GetConfig())

	database.Init()

	application.database = database

	redis := NewRedis(nacos.GetConfig())

	redis.Init()

	application.redis = redis

	RegisterDefaultMiddleware(
		application.engine,
		&SystemMiddleware{
			Logger:     application.logger,
			Config:     nacos.GetConfig(),
			Version:    version,
			SyncDBFunc: syncDBFunc,
		})

	return application
}

func (m *Application) GetEngine() *gin.Engine {
	return m.engine
}

func (m *Application) GetNacos() *Nacos {
	return m.nacos
}

func (m *Application) GetLogger() *zap.Logger {
	return m.logger
}

func (m *Application) GetDatabase() *Database {
	return m.database
}

func (m *Application) GetRedis() *Redis {
	return m.redis
}

func (m *Application) RegisterScheduler(scheduler *Scheduler) {
	m.engine.GET("/system/jobs", m.Jobs(scheduler))
	m.engine.POST("/system/jobs/start", m.Start(scheduler))
	m.engine.POST("/system/jobs/stop", m.StopJobs(scheduler))
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

func (m *Application) Jobs(scheduler *Scheduler) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		RenderSuccessResult(ctx, scheduler.GetJobs())
	}
}

func (m *Application) Start(scheduler *Scheduler) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		scheduler.InnerScheduler.Start()

		RenderSuccessResult(ctx, NewSuccessRestResult(true, "已启动调度器"))
	}
}

func (m *Application) StopJobs(scheduler *Scheduler) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		if err := scheduler.InnerScheduler.StopJobs(); err != nil {
			RenderInternalServerErrorResult(ctx, err)
		} else {
			RenderSuccessResult(ctx, NewSuccessRestResult(true, "已停止所有任务"))
		}
	}
}
