package zentao

import (
	"fmt"
	"net/url"
)

type Task struct {
	Id           int          `json:"id"`           // 任务ID
	Project      int          `json:"project"`      // 所属项目
	Parent       int          `json:"parent"`       // 父任务
	Execution    int          `json:"execution"`    // 所属执行
	Module       int          `json:"module"`       // 所属模块
	Story        int          `json:"story"`        // 所属需求
	FromBug      int          `json:"fromBug"`      // 来源于Bug
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
	EstStarted   string       `json:"estStarted"`   // 预计开始时间
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

// 获取任务详情
func (s *ZentaoClient) GetTask(token string, id int) (int, []byte, *Task, error) {
	if status, bytes, err := s.webClient.Get(s.url, fmt.Sprintf(getTaskPath, id), nil, nil, setTokenHeaderFunc(token)); err != nil {
		return status, bytes, nil, err
	} else {
		task := &Task{}

		status, bytes, err := parseResult(status, bytes, task)

		return status, bytes, task, err
	}
}

// 获取执行任务列表
func (s *ZentaoClient) GetExecutionTasks(token string, executionId int, pageParam PageParam, urlValues url.Values) (int, []byte, *TaskPageResult, error) {
	if urlValues == nil {
		urlValues = make(url.Values)
	}

	urlValues.Set("page", pageParam.Page)
	urlValues.Set("limit", pageParam.Limit)

	if status, bytes, err := s.webClient.Get(s.url, fmt.Sprintf(getExecutionTasksPath, executionId), urlValues, nil, setTokenHeaderFunc(token)); err != nil {
		return status, bytes, nil, err
	} else {
		pageResult := &TaskPageResult{}

		status, bytes, err := parseResult(status, bytes, pageResult)

		return status, bytes, pageResult, err
	}
}
