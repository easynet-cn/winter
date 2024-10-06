package zentao

import (
	"fmt"
	"net/url"
)

type Project struct {
	Id         int          `json:"id"`         // 项目ID
	Parent     int          `json:"parent"`     // 所属项目集
	Name       string       `json:"name"`       // 项目名称
	Code       string       `json:"code"`       // 项目编号
	Model      string       `json:"model"`      // 项目模型(scrum敏捷 | waterfall 瀑布)
	Budget     string       `json:"budget"`     // 项目预算
	BudgetUnit string       `json:"budgetUnit"` // 预算币种(CNY | USD)
	Begin      *string      `json:"begin"`      // 预计开始日期
	End        *string      `json:"end"`        // 预计结束日期
	RealBegan  *string      `json:"realBegan"`  // 实际开始日期
	RealEnd    *string      `json:"realEnd"`    // 实际结束日期
	Status     string       `json:"status"`     // 项目状态(wait 未开始 | doing 进行中 | suspend 已挂起 | closed 已关闭)
	Desc       string       `json:"desc"`       // 项目描述
	OpenedBy   *SimpleUser  `json:"openedBy"`   // 创建人
	OpenedDate string       `json:"openedDate"` // 创建时间
	PM         *SimpleUser  `json:"pm"`         // 项目经理
	Acl        string       `json:"acl"`        // 访问控制(open 公开 | private 私有)
	Whitelist  []SimpleUser `json:"whitelist"`  // 白名单
	Progress   string       `json:"progress"`   // 进度
}

type ProjectPageResult struct {
	PageResult
	Projects []Project `json:"projects"`
}

// 获取项目列表
func (s *ZentaoClient) GetProjects(token string, pageParam PageParam, urlValues url.Values) (int, []byte, *ProjectPageResult, error) {
	if urlValues == nil {
		urlValues = make(url.Values)
	}

	urlValues.Set("page", pageParam.Page)
	urlValues.Set("limit", pageParam.Limit)

	if status, bytes, err := s.webClient.Get(s.url, getProjectsPath, urlValues, nil, SetTokenHeaderFunc(token)); err != nil {
		return status, bytes, nil, err
	} else {
		pageResult := &ProjectPageResult{}

		status, bytes, err := ParseResult(status, bytes, pageResult)

		return status, bytes, pageResult, err
	}
}

// 获取项目详情
func (s *ZentaoClient) GetProject(token string, id int) (int, []byte, *Project, error) {
	if status, bytes, err := s.webClient.Get(s.url, fmt.Sprintf(getProjectPath, id), nil, nil, SetTokenHeaderFunc(token)); err != nil {
		return status, bytes, nil, err
	} else {
		project := &Project{}

		status, bytes, err := ParseResult(status, bytes, project)

		return status, bytes, project, err
	}
}
