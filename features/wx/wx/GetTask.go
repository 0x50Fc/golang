package wx

type GetTask struct {
	Type	int32	`json:"type" name:"type" title:"类型,多个逗号分割"`
	Appid	string	`json:"appid" name:"appid" title:"appid"`
	Openid	string	`json:"openid" name:"openid" title:"openid"`
	Update	interface{}	`json:"update,omitempty" name:"update" title:"是否更新用户信息"`
}

func (T *GetTask) GetName() string {
	return "get.json"
}

func (T *GetTask) GetTitle() string {
	return "获取用户"
}

