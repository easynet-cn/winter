package winter

type PageParam struct {
	PageIndex int `json:"pageIndex" form:"pageIndex"`
	PageSize  int `json:"pageSize" form:"pageSize"`
}

func (pageParam *PageParam) Start() int {
	return (pageParam.PageIndex - 1) * pageParam.PageSize
}

func (pageParam *PageParam) ParamLength() int {
	return 2
}
