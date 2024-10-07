package winter

type IdName struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}

func NewIdName(id int64, name string) *IdName {
	return &IdName{
		Id:   id,
		Name: name,
	}
}
