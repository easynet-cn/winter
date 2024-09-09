package zentao

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/easynet-cn/winter"
)

const (
	tokenHeader = "Token"

	getTokenPath             = "/api.php/v1/tokens"
	getDepartmentsPath       = "/api.php/v1/departments"
	getCurrentUserPath       = "/api.php/v1/user"
	getProjectsPath          = "/api.php/v1/projects"
	getProjectExecutionsPath = "/api.php/v1/projects/%d/executions"
	getProjectStoriesPath    = "/api.php/v1/projects/%d/stories"
	getExecutionStoriesPath  = "/api.php/v1/executions/%d/stories"
	getExecutionTasksPath    = "/api.php/v1/executions/%d/tasks"
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

type GetTokenRequest struct {
	Account  string `json:"account"`
	Password string `json:"password"`
}

type GetTokenResponse struct {
	Token string `json:"token"`
}

type Department struct {
	Id       int    `json:"id"`       // 部门编号
	Name     string `json:"name"`     // 部门名称
	Parent   int    `json:"parent"`   // 上级部门
	Path     string `json:"path"`     // 路径
	Grade    int    `json:"grade"`    // 部门级别
	Order    int    `json:"order"`    // 排序
	Position string `json:"position"` // 职位
	Function string `json:"function"` // 职能
	Manager  string `json:"manager"`  // 负责人
}

type UserProfile struct {
	Profile User `json:"profile"`
}

type User struct {
	Id       int     `json:"id"`       // 用户编号
	Type     string  `json:"type"`     // 类型(inside 内部用户 | outside 外部用户)
	Dept     int     `json:"dept"`     // 所属部门
	Account  string  `json:"account"`  // 用户名
	RealName string  `json:"realname"` // 真实姓名
	Nickname string  `json:"nickname"` // 昵称
	Avatar   string  `json:"avatar"`   // 头像
	Birthday *string `json:"birthday"` // 生日
	Gender   string  `json:"gender"`   // 性别(f 女性 | m 男性)
	Mobile   string  `json:"mobile"`   // 手机号
	Phone    string  `json:"phone"`    // 电话号码
	Weixin   string  `json:"weixin"`   // 微信号码
	Address  string  `json:"address"`  // 住址
	Join     string  `json:"join"`     // 加入日期
	Admin    bool    `json:"admin"`    // 是否管理员
}

type SimpleUser struct {
	Id       int    `json:"id"`       // 用户编号
	Account  string `json:"account"`  // 用户名
	Avatar   string `json:"avatar"`   // 头像
	RealName string `json:"realname"` // 真实姓名
}

type Project struct {
	Id         int     `json:"id"`         // 项目ID
	Name       string  `json:"name"`       // 项目名称
	Code       string  `json:"code"`       // 项目编号
	Model      string  `json:"model"`      // 项目模型(scrum敏捷 | waterfall 瀑布)
	Budget     int     `json:"budget"`     // 项目预算
	BudgetUnit string  `json:"budgetUnit"` // 预算币种(CNY | USD)
	Parent     int     `json:"parent"`     // 所属项目集
	Begin      *string `json:"begin"`      // 预计开始日期
	End        *string `json:"end"`        // 预计结束日期
	Status     string  `json:"status"`     // 项目状态(wait 未开始 | doing 进行中 | suspend 已挂起 | closed 已关闭)
	OpenedBy   string  `json:"openedBy"`   // 创建人
	OpenedDate string  `json:"openedDate"` // 创建时间
	PM         string  `json:"pm"`         // 项目经理
	Progress   int     `json:"progress"`   // 进度
}

type ProjectPageResult struct {
	PageResult
	Projects []Project `json:"projects"`
}

type Execution struct {
	Id         int     `json:"id"`         // 执行ID
	Name       string  `json:"name"`       // 执行名称
	Code       string  `json:"code"`       // 执行代号
	Begin      *string `json:"begin"`      // 预计开始日期
	End        *string `json:"end"`        // 预计结束日期
	Status     string  `json:"status"`     // 项目状态(wait 未开始 | doing 进行中 | suspend 已挂起 | closed 已关闭)
	OpenedBy   string  `json:"openedBy"`   // 创建人
	OpenedDate string  `json:"openedDate"` // 创建时间
	Progress   int     `json:"progress"`   // 进度
}

type ExecutionPageResult struct {
	PageResult
	Executions []Execution `json:"executions"`
}

type Story struct {
	Id         int    `json:"id"`         // 需求ID
	Product    int    `json:"product"`    // 所属产品
	Branch     int    `json:"branch"`     // 所属分支
	Module     int    `json:"module"`     // 所属产品模块
	FromBug    int    `json:"fromBug"`    // 来自于Bug
	Source     string `json:"source"`     // 需求来源(customer 客户 | user 用户 | po 产品经理 | market 市场)
	SourceNo   string `json:"sourceNo"`   // 来源备注
	Title      string `json:"title"`      // 需求标题
	Category   string `json:"category"`   // 类型(feature 功能 | interface 接口 | performance 性能 | safe 安全 | experience 体验 | improve 改进 | other 其他)
	Stage      string `json:"stage"`      // 阶段(wait 未开始 | planned 已计划 | projected 已立项 | developing 研发中 | developed 研发完毕 | testing 测试中 | tested 测试完毕 | verified 已验收 | released 已发布 | closed 已关闭)
	Pri        int    `json:"pri"`        // 优先级
	Estimate   int    `json:"estimate"`   // 预计工时
	Status     string `json:"status"`     // 状态(draft 草稿 | active 激活 | closed 已关闭 | changed 已变更)
	OpenedBy   string `json:"openedBy"`   // 创建人
	OpenedDate string `json:"openedDate"` // 创建时间
	ToBug      int    `json:"toBug"`      // 转为Bug
}

type StoryPageResult struct {
	PageResult
	Stories []Story `json:"stories"`
}

type Task struct {
	Id           int          `json:"id"`           // 任务ID
	Project      int          `json:"project"`      // 所属项目
	Execution    int          `json:"execution"`    // 所属执行
	Module       int          `json:"module"`       // 所属模块
	Story        int          `json:"story"`        // 所属需求
	Name         string       `json:"name"`         // 任务名称
	Type         string       `json:"type"`         // 任务类型(design 设计 | devel 开发 | request 需求 | test 测试 | study 研究 | discuss 讨论 | ui 界面 | affair 事务 | misc 其他)
	Pri          int          `json:"pri"`          // 优先级
	Estimate     float64      `json:"estimate"`     // 预计工时
	Left         float64      `json:"left"`         // 剩余工时
	Deadline     *string      `json:"deadline"`     // 截止日期
	Consumed     float64      `json:"consumed"`     // 消耗工时
	Status       string       `json:"status"`       // 状态(wait 未开始 | doing 进行中 | done 已完成 | closed 已关闭 | cancel 已取消)
	Desc         string       `json:"desc"`         // 任务描述
	OpenedBy     *SimpleUser  `json:"openedBy"`     // 创建人
	OpenedDate   string       `json:"openedDate"`   // 创建时间
	AssignedTo   *SimpleUser  `json:"assignedTo"`   // 指派给
	EsStarted    string       `json:"esStarted"`    // 预计开始时间
	RealStarted  string       `json:"realStarted"`  // 实际开始时间
	FinishedBy   *SimpleUser  `json:"finishedBy"`   // 由谁完成
	FinishedDate string       `json:"finishedDate"` // 完成时间
	ClosedBy     *SimpleUser  `json:"closedBy"`     // 由谁关闭
	ClosedDate   string       `json:"closedDate"`   // 关闭时间
	MailTo       []SimpleUser `json:"mailTo"`       // 抄送给
}

type TaskPageResult struct {
	PageResult
	Tasks []Task `json:"tasks"`
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

// 获取Token
func (s *ZentaoClient) GetToken(request *GetTokenRequest) (int, []byte, *GetTokenResponse, error) {
	if status, bytes, err := s.webClient.Post(s.url, getTokenPath, nil, request, nil); err != nil {
		return status, bytes, nil, err
	} else if s.statusIsOk(status) && len(bytes) > 0 {
		response := &GetTokenResponse{}

		if err := json.Unmarshal(bytes, &response); err != nil {
			return status, bytes, nil, err
		}

		return status, bytes, response, nil
	} else {
		return status, bytes, nil, nil
	}
}

// 获取部门列表
func (s *ZentaoClient) GetDepartments(token string) (int, []byte, []Department, error) {
	if status, bytes, err := s.webClient.Get(s.url, getDepartmentsPath, nil, nil, nil, s.setTokenHeaderfunc(token)); err != nil {
		return status, bytes, nil, err
	} else if s.statusIsOk(status) && len(bytes) > 0 {
		departments := make([]Department, 0)

		if err := json.Unmarshal(bytes, &departments); err != nil {
			return status, bytes, nil, err
		}

		return status, bytes, departments, nil
	} else {
		return status, bytes, nil, nil
	}
}

// 获取当前用户信息
func (s *ZentaoClient) GetCurrentUser(token string) (int, []byte, *UserProfile, error) {
	if status, bytes, err := s.webClient.Get(s.url, getCurrentUserPath, nil, nil, nil, s.setTokenHeaderfunc(token)); err != nil {
		return status, bytes, nil, err
	} else if s.statusIsOk(status) && len(bytes) > 0 {
		profile := &UserProfile{}

		if err := json.Unmarshal(bytes, &profile); err != nil {
			return status, bytes, nil, err
		}

		return status, bytes, profile, nil
	} else {
		return status, bytes, nil, nil
	}
}

// 获取项目列表
func (s *ZentaoClient) GetProjects(token string, pageParam PageParam, urlValues url.Values) (int, []byte, []Project, error) {
	if urlValues == nil {
		urlValues = make(url.Values)
	}

	urlValues.Set("page", pageParam.Page)
	urlValues.Set("limit", pageParam.Limit)

	if status, bytes, err := s.webClient.Get(s.url, getProjectsPath, urlValues, nil, s.setTokenHeaderfunc(token)); err != nil {
		return status, bytes, nil, err
	} else if s.statusIsOk(status) && len(bytes) > 0 {
		projects := make([]Project, 0)

		if err := json.Unmarshal(bytes, &projects); err != nil {
			return status, bytes, nil, err
		}

		return status, bytes, projects, nil
	} else {
		return status, bytes, nil, nil
	}
}

// 获取项目执行列表
func (s *ZentaoClient) GetProjectExecutions(token string, projectId int, pageParam PageParam, urlValues url.Values) (int, []byte, *ExecutionPageResult, error) {
	if urlValues == nil {
		urlValues = make(url.Values)
	}

	urlValues.Set("page", pageParam.Page)
	urlValues.Set("limit", pageParam.Limit)

	if status, bytes, err := s.webClient.Get(s.url, fmt.Sprintf(getProjectExecutionsPath, projectId), urlValues, nil, s.setTokenHeaderfunc(token)); err != nil {
		return status, bytes, nil, err
	} else if s.statusIsOk(status) && len(bytes) > 0 {
		pageResult := &ExecutionPageResult{}

		if err := json.Unmarshal(bytes, &pageResult); err != nil {
			return status, bytes, nil, err
		}

		return status, bytes, pageResult, nil
	} else {
		return status, bytes, nil, nil
	}
}

// 获取项目需求列表
func (s *ZentaoClient) GetProjectStories(token string, projectId int, pageParam PageParam, urlValues url.Values) (int, []byte, *StoryPageResult, error) {
	if urlValues == nil {
		urlValues = make(url.Values)
	}

	urlValues.Set("page", pageParam.Page)
	urlValues.Set("limit", pageParam.Limit)

	if status, bytes, err := s.webClient.Get(s.url, fmt.Sprintf(getProjectStoriesPath, projectId), urlValues, nil, s.setTokenHeaderfunc(token)); err != nil {
		return status, bytes, nil, err
	} else if s.statusIsOk(status) && len(bytes) > 0 {
		pageResult := &StoryPageResult{}

		if err := json.Unmarshal(bytes, &pageResult); err != nil {
			return status, bytes, nil, err
		}

		return status, bytes, pageResult, nil
	} else {
		return status, bytes, nil, nil
	}
}

// 获取执行需求列表
func (s *ZentaoClient) GetExecutionStories(token string, executionId int, pageParam PageParam, urlValues url.Values) (int, []byte, *StoryPageResult, error) {
	if urlValues == nil {
		urlValues = make(url.Values)
	}

	urlValues.Set("page", pageParam.Page)
	urlValues.Set("limit", pageParam.Limit)

	if status, bytes, err := s.webClient.Get(s.url, fmt.Sprintf(getExecutionStoriesPath, executionId), urlValues, nil, s.setTokenHeaderfunc(token)); err != nil {
		return status, bytes, nil, err
	} else if s.statusIsOk(status) && len(bytes) > 0 {
		pageResult := &StoryPageResult{}

		if err := json.Unmarshal(bytes, &pageResult); err != nil {
			return status, bytes, nil, err
		}

		return status, bytes, pageResult, nil
	} else {
		return status, bytes, nil, nil
	}
}

// 获取执行任务列表
func (s *ZentaoClient) GetExecutionTasks(token string, executionId int, pageParam PageParam, urlValues url.Values) (int, []byte, *TaskPageResult, error) {
	if urlValues == nil {
		urlValues = make(url.Values)
	}

	urlValues.Set("page", pageParam.Page)
	urlValues.Set("limit", pageParam.Limit)

	if status, bytes, err := s.webClient.Get(s.url, fmt.Sprintf(getExecutionTasksPath, executionId), urlValues, nil, s.setTokenHeaderfunc(token)); err != nil {
		return status, bytes, nil, err
	} else if s.statusIsOk(status) && len(bytes) > 0 {
		pageResult := &TaskPageResult{}

		if err := json.Unmarshal(bytes, &pageResult); err != nil {
			return status, bytes, nil, err
		}

		return status, bytes, pageResult, nil
	} else {
		return status, bytes, nil, nil
	}
}

func (s *ZentaoClient) statusIsOk(status int) bool {
	return status >= http.StatusOK && status < http.StatusMultipleChoices
}

func (s *ZentaoClient) setTokenHeaderfunc(token string) winter.RequestHeaderFunc {
	return func(header http.Header) {
		header.Set(tokenHeader, token)
	}
}
