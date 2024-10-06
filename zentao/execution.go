package zentao

import (
	"fmt"
	"net/url"
)

type Execution struct {
	Id          int          `json:"id"`          // 执行ID
	Project     int          `json:"project"`     // 所属项目
	Name        string       `json:"name"`        // 执行名称
	Code        string       `json:"code"`        // 执行代号
	Begin       *string      `json:"begin"`       // 预计开始日期
	End         *string      `json:"end"`         // 预计结束日期
	Days        int          `json:"days"`        // 可用工作日
	Lifetime    string       `json:"lifetime"`    // 类型(short 短期 | long 长期 | ops 运维)
	PO          *SimpleUser  `json:"po"`          // 产品负责人
	PM          *SimpleUser  `json:"pm"`          // 迭代负责人
	QD          *SimpleUser  `json:"qd"`          // 测试负责人
	RD          *SimpleUser  `json:"rd"`          // 发布负责人
	TeamMembers []SimpleUser `json:"teamMembers"` // 团队成员
	Desc        string       `json:"desc"`        // 迭代描述
	Acl         string       `json:"acl"`         // 访问控制(private 私有 | open 继承项目权限)
	Whitelist   []SimpleUser `json:"whitelist"`   // 白名单
	Status      string       `json:"status"`      // 项目状态(wait 未开始 | doing 进行中 | suspend 已挂起 | closed 已关闭)
	OpenedBy    *SimpleUser  `json:"openedBy"`    // 创建人
	OpenedDate  string       `json:"openedDate"`  // 创建时间
	Progress    any          `json:"progress"`    // 进度
}

type ExecutionPageResult struct {
	PageResult
	Executions []Execution `json:"executions"`
}

// 获取执行详情
func (s *ZentaoClient) GetExecution(token string, id int) (int, []byte, *Execution, error) {
	if status, bytes, err := s.webClient.Get(s.url, fmt.Sprintf(getExecutionPath, id), nil, nil, SetTokenHeaderFunc(token)); err != nil {
		return status, bytes, nil, err
	} else {
		execution := &Execution{}

		status, bytes, err := ParseResult(status, bytes, execution)

		return status, bytes, execution, err
	}
}

// 获取项目执行列表
func (s *ZentaoClient) GetProjectExecutions(token string, projectId int, pageParam PageParam, urlValues url.Values) (int, []byte, *ExecutionPageResult, error) {
	if urlValues == nil {
		urlValues = make(url.Values)
	}

	urlValues.Set("page", pageParam.Page)
	urlValues.Set("limit", pageParam.Limit)

	if status, bytes, err := s.webClient.Get(s.url, fmt.Sprintf(getProjectExecutionsPath, projectId), urlValues, nil, SetTokenHeaderFunc(token)); err != nil {
		return status, bytes, nil, err
	} else {
		pageResult := &ExecutionPageResult{}

		status, bytes, err := ParseResult(status, bytes, pageResult)

		return status, bytes, pageResult, err
	}
}
