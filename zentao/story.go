package zentao

import (
	"fmt"
	"net/url"
)

type Story struct {
	Id         int         `json:"id"`         // 需求ID
	Product    int         `json:"product"`    // 所属产品
	Branch     int         `json:"branch"`     // 所属分支
	Module     int         `json:"module"`     // 所属产品模块
	FromBug    int         `json:"fromBug"`    // 来自于Bug
	Source     string      `json:"source"`     // 需求来源(customer 客户 | user 用户 | po 产品经理 | market 市场)
	SourceNote string      `json:"sourceNote"` // 来源备注
	Title      string      `json:"title"`      // 需求标题
	Category   string      `json:"category"`   // 类型(feature 功能 | interface 接口 | performance 性能 | safe 安全 | experience 体验 | improve 改进 | other 其他)
	Stage      string      `json:"stage"`      // 阶段(wait 未开始 | planned 已计划 | projected 已立项 | developing 研发中 | developed 研发完毕 | testing 测试中 | tested 测试完毕 | verified 已验收 | released 已发布 | closed 已关闭)
	Pri        int         `json:"pri"`        // 优先级
	Estimate   int         `json:"estimate"`   // 预计工时
	Verify     string      `json:"verify"`     // 验收标准
	Status     string      `json:"status"`     // 状态(draft 草稿 | active 激活 | closed 已关闭 | changed 已变更)
	OpenedBy   *SimpleUser `json:"openedBy"`   // 创建人
	OpenedDate string      `json:"openedDate"` // 创建时间
	ToBug      int         `json:"toBug"`      // 转为Bug
}

type StoryPageResult struct {
	PageResult
	Stories []Story `json:"stories"`
}

// 获取需求详情
func (s *ZentaoClient) GetStory(token string, id int) (int, []byte, *Story, error) {
	if status, bytes, err := s.webClient.Get(s.url, fmt.Sprintf(getStoryPath, id), nil, nil, setTokenHeaderFunc(token)); err != nil {
		return status, bytes, nil, err
	} else {
		story := &Story{}

		status, bytes, err := parseResult(status, bytes, story)

		return status, bytes, story, err
	}
}

// 获取项目需求列表
func (s *ZentaoClient) GetProjectStories(token string, projectId int, pageParam PageParam, urlValues url.Values) (int, []byte, *StoryPageResult, error) {
	if urlValues == nil {
		urlValues = make(url.Values)
	}

	urlValues.Set("page", pageParam.Page)
	urlValues.Set("limit", pageParam.Limit)

	if status, bytes, err := s.webClient.Get(s.url, fmt.Sprintf(getProjectStoriesPath, projectId), urlValues, nil, setTokenHeaderFunc(token)); err != nil {
		return status, bytes, nil, err
	} else {
		pageResult := &StoryPageResult{}

		status, bytes, err := parseResult(status, bytes, pageResult)

		return status, bytes, pageResult, err
	}
}

// 获取执行需求列表
func (s *ZentaoClient) GetExecutionStories(token string, executionId int, pageParam PageParam, urlValues url.Values) (int, []byte, *StoryPageResult, error) {
	if urlValues == nil {
		urlValues = make(url.Values)
	}

	urlValues.Set("page", pageParam.Page)
	urlValues.Set("limit", pageParam.Limit)

	if status, bytes, err := s.webClient.Get(s.url, fmt.Sprintf(getExecutionStoriesPath, executionId), urlValues, nil, setTokenHeaderFunc(token)); err != nil {
		return status, bytes, nil, err
	} else {
		pageResult := &StoryPageResult{}

		status, bytes, err := parseResult(status, bytes, pageResult)

		return status, bytes, pageResult, err
	}
}
