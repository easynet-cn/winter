package winter

import (
	"net/http"
	"runtime"

	"github.com/gin-gonic/gin"
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
