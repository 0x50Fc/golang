package vcode

type DelTask struct {
	Key	string	`json:"key" name:"key" title:"Key"`
}

func (T *DelTask) GetName() string {
	return "del.json"
}

func (T *DelTask) GetTitle() string {
	return "删除验证码"
}

