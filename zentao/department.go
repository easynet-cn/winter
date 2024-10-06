package zentao

import (
	"fmt"
)

type Department struct {
	Id       int          `json:"id"`       // 部门编号
	Name     string       `json:"name"`     // 部门名称
	Parent   int          `json:"parent"`   // 上级部门
	Path     string       `json:"path"`     // 路径
	Grade    int          `json:"grade"`    // 部门级别
	Order    int          `json:"order"`    // 排序
	Position string       `json:"position"` // 职位
	Function string       `json:"function"` // 职能
	Manager  string       `json:"manager"`  // 负责人
	Children []Department `json:"children"` // 子部门
}

// 获取部门列表
func (s *ZentaoClient) GetDepartments(token string) (int, []byte, []Department, error) {
	if status, bytes, err := s.webClient.Get(s.url, getDepartmentsPath, nil, nil, nil, SetTokenHeaderFunc(token)); err != nil {
		return status, bytes, nil, err
	} else {
		departments := make([]Department, 0)

		status, bytes, err := ParseResult(status, bytes, &departments)

		return status, bytes, departments, err
	}
}

// 获取部门详情
func (s *ZentaoClient) GetDepartment(token string, id int) (int, []byte, *Department, error) {
	if status, bytes, err := s.webClient.Get(s.url, fmt.Sprintf(getDepartmentPath, id), nil, nil, nil, SetTokenHeaderFunc(token)); err != nil {
		return status, bytes, nil, err
	} else {
		department := &Department{}

		status, bytes, err := ParseResult(status, bytes, department)

		return status, bytes, department, err
	}
}
