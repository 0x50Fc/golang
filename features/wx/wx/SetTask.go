package wx

type SetTask struct {
	Type	interface{}	`json:"type,omitempty" name:"type" title:"类型"`
	Appid	interface{}	`json:"appid,omitempty" name:"appid" title:"appid"`
	Openid	interface{}	`json:"openid,omitempty" name:"openid" title:"openid"`
	Unionid	interface{}	`json:"unionid,omitempty" name:"unionid" title:"unionid"`
	State	int32	`json:"state" name:"state" title:"关注状态"`
}

func (T *SetTask) GetName() string {
	return "set.json"
}

func (T *SetTask) GetTitle() string {
	return "修改用户关注状态"
}

