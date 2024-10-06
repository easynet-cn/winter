package zentao

import (
	"encoding/json"
	"net/http"

	"github.com/easynet-cn/winter"
)

const (
	tokenHeader = "Token"

	getTokenPath             = "/api.php/v1/tokens"
	getDepartmentsPath       = "/api.php/v1/departments"
	getDepartmentPath        = "/api.php/v1/departments/%d"
	getCurrentUserPath       = "/api.php/v1/user"
	getProjectsPath          = "/api.php/v1/projects"
	getProjectPath           = "/api.php/v1/projects/%d"
	getProjectExecutionsPath = "/api.php/v1/projects/%d/executions"
	getProjectStoriesPath    = "/api.php/v1/projects/%d/stories"
	getStoryPath             = "/api.php/v1/stories/%d"
	getExecutionPath         = "/api.php/v1/executions/%d"
	getExecutionStoriesPath  = "/api.php/v1/executions/%d/stories"
	getExecutionTasksPath    = "/api.php/v1/executions/%d/tasks"
	getTaskPath              = "/api.php/v1/tasks/%d"
)

type PageParam struct {
	Page  string `json:"page"`  // 第几页，默认为1
	Limit string `json:"limit"` // 每页项目数量，默认20
}

type PageResult struct {
	Page  int `json:"page"`  // 页数
	Total int `json:"total"` // 总数
	Limit int `json:"limit"` // 每页数量
}

type ZentaoClient struct {
	url       string
	webClient *winter.WebClient
}

func NewZentaoClient(url string, webClient *winter.WebClient) *ZentaoClient {
	return &ZentaoClient{
		url:       url,
		webClient: webClient,
	}
}

func ParseError(bytes []byte) string {
	if len(bytes) == 0 {
		return ""
	}

	errMap := make(map[string]any)

	if err := json.Unmarshal(bytes, &errMap); err == nil {
		if msg, ok := errMap["error"].(string); ok {
			return msg
		}
	}

	return string(bytes)
}

func parseResult[T any](status int, bytes []byte, result *T) (int, []byte, error) {
	if statusIsOk(status) && len(bytes) > 0 {
		if err := json.Unmarshal(bytes, &result); err != nil {
			return status, bytes, err
		}

		return status, bytes, nil
	} else {
		return status, bytes, nil
	}
}

func statusIsOk(status int) bool {
	return status >= http.StatusOK && status < http.StatusMultipleChoices
}

func setTokenHeaderFunc(token string) winter.RequestHeaderFunc {
	return func(header http.Header) {
		header.Set(tokenHeader, token)
	}
}
