package winter

import (
	"fmt"
	"net/http"
	"runtime"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

func SystemStats(ctx *gin.Context) {
	memStats := &runtime.MemStats{}

	runtime.ReadMemStats(memStats)

	ctx.JSON(http.StatusOK, gin.H{
		"GOARCH":       runtime.GOARCH,
		"GOOS":         runtime.GOOS,
		"NumCPU":       runtime.NumCPU(),
		"NumGoroutine": runtime.NumGoroutine(),
		"MemStats":     memStats,
		"Version":      runtime.Version(),
	})
}

func Recovery(logger *zap.Logger) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		defer func() {
			if r := recover(); r != nil {
				logger.Error("middleware", zap.Error(fmt.Errorf("%v", r)))

				ctx.JSON(http.StatusOK, RestResult{Status: 500, Code: "500", Message: "系统内部错误", Error: fmt.Sprintf("%v", r)})

			}
		}()

		ctx.Next()
	}
}

func Version(config *viper.Viper, version string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, NewSystemVersion(version, config))
	}
}

func SyncDB(f func() error) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if err := f(); err != nil {
			ctx.JSON(http.StatusInternalServerError, RestResult{Status: 500, Code: "500", Message: err.Error()})
		} else {
			ctx.JSON(http.StatusOK, RestResult{Status: 200, Code: "200", Message: "同步成功"})
		}

	}
}
