package winter

import (
	"runtime"

	"github.com/dromara/carbon/v2"
	"github.com/spf13/viper"
)

type SystemVersion struct {
	ApplicationName string `json:"applicationName"`
	Version         string `json:"version"`
	Profile         string `json:"profile"`
	GoVersion       string `json:"goVersion"`
	CurrentTime     string `json:"currentTime"`
}

func NewSystemVersion(version string, config *viper.Viper) *SystemVersion {
	return &SystemVersion{
		ApplicationName: config.GetString("spring.application.name"),
		Version:         version,
		Profile:         config.GetString("spring.profiles.active"),
		GoVersion:       runtime.Version(),
		CurrentTime:     carbon.Now().ToDateTimeString(),
	}
}
