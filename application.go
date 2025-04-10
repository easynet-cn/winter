package winter

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

type ApplicationFunc func()

type Application struct {
	config     *viper.Viper
	engine     *gin.Engine
	consul     *Consul
	nacos      *Nacos
	logger     *zap.Logger
	database   *Database
	redis      *Redis
	tencentEss *TencentEss
}

func NewApplication(
	metadata map[string]string,
	version string,
	syncDBFunc func() error,
) *Application {
	consul := NewConsul()

	consul.Init()

	config := consul.GetConfig()

	nacos := NewNacos(metadata)

	nacos.Init()

	keys := nacos.config.AllKeys()

	for _, key := range keys {
		config.Set(key, nacos.config.Get(key))
	}

	logger := NewLogger(nacos.GetConfig())

	gin.SetMode(gin.ReleaseMode)

	engine := gin.Default()

	application := &Application{
		config: config,
		consul: consul,
		engine: engine,
		nacos:  nacos,
		logger: logger,
	}

	application.logger = NewLogger(config)

	database := NewDatabase(config)

	database.Init()

	application.database = database

	redis := NewRedis(config)

	redis.Init()

	application.redis = redis

	tencentEss := NewTencentEss(config)

	tencentEss.Init()

	application.tencentEss = tencentEss

	RegisterDefaultMiddleware(
		application.engine,
		&SystemMiddleware{
			Logger:     application.logger,
			Config:     config,
			Version:    version,
			SyncDBFunc: syncDBFunc,
		})

	return application
}

func (m *Application) GetConfig() *viper.Viper {
	return m.config
}

func (m *Application) GetEngine() *gin.Engine {
	return m.engine
}

func (m *Application) GetConsul() *Consul {
	return m.consul
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

func (m *Application) GetTencentEss() *TencentEss {
	return m.tencentEss
}

func (m *Application) RegisterScheduler(scheduler *Scheduler) {
	m.engine.GET("/system/jobs", m.Jobs(scheduler))
	m.engine.POST("/system/jobs/start", m.Start(scheduler))
	m.engine.POST("/system/jobs/stop", m.StopJobs(scheduler))
	m.engine.DELETE("/system/jobs/:id", m.RemoveJobById(scheduler))
}

func (m *Application) Run(funcs ...ApplicationFunc) {
	for _, f := range funcs {
		if f != nil {
			f()
		}
	}

	if err := m.engine.Run(fmt.Sprintf(":%d", m.config.GetInt("server.port"))); err != nil {
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

func (m *Application) RemoveJobById(scheduler *Scheduler) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")

		if uuid, err := uuid.Parse(id); err != nil {
			RenderBadRequestResult(ctx, err)
		} else if err := scheduler.RemoveJob(uuid); err != nil {
			RenderInternalServerErrorResult(ctx, err)
		} else {
			RenderSuccessResult(ctx, NewSuccessRestResult(true, fmt.Sprintf("已删除任务（%s）", id)))
		}
	}
}
