package winter

import "math"

type PageResult struct {
	Total      int64 `json:"total"`
	TotalPages int64 `json:"totalPages"`
	Data       []any `json:"data"`
}

func NewPageResult() *PageResult {
	return &PageResult{Total: 0, TotalPages: 0, Data: make([]any, 0)}
}

func NewPatgeResultWithTotal(total int64) *PageResult {
	return &PageResult{Total: total, TotalPages: 0, Data: make([]any, 0)}
}

func NewPageResultWithTotalAndDataLen(total int64, l int) *PageResult {
	return &PageResult{Total: total, TotalPages: 0, Data: make([]any, 0, l)}
}

func (m *PageResult) GetTotalPages(pageSize int) int64 {
	if m.Total == 0 || pageSize == 0 {
		return 0
	}

	return int64(math.Ceil(float64(m.Total) / float64(pageSize)))
}
