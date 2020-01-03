package vcode

type GetTask struct {
	Key	string	`json:"key" name:"key" title:"Key"`
}

func (T *GetTask) GetName() string {
	return "get.json"
}

func (T *GetTask) GetTitle() string {
	return "获取验证码"
}

