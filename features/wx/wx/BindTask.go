package wx

type BindTask struct {
	Type	interface{}	`json:"type,omitempty" name:"type" title:"类型"`
	Appid	interface{}	`json:"appid,omitempty" name:"appid" title:"appid"`
	Openid	interface{}	`json:"openid,omitempty" name:"openid" title:"openid"`
	Unionid	interface{}	`json:"unionid,omitempty" name:"unionid" title:"unionid"`
	Uid	int64	`json:"uid" name:"uid" title:"用户ID"`
}

func (T *BindTask) GetName() string {
	return "bind.json"
}

func (T *BindTask) GetTitle() string {
	return "绑定用户"
}

