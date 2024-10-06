package zentao

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

type UserProfile struct {
	Profile User `json:"profile"`
}

// 获取当前用户信息
func (s *ZentaoClient) GetCurrentUser(token string) (int, []byte, *UserProfile, error) {
	if status, bytes, err := s.webClient.Get(s.url, getCurrentUserPath, nil, nil, nil, SetTokenHeaderFunc(token)); err != nil {
		return status, bytes, nil, err
	} else {
		profile := &UserProfile{}

		status, bytes, err := ParseResult(status, bytes, profile)

		return status, bytes, profile, err
	}
}
