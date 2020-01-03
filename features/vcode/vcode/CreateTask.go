package vcode

type CreateTask struct {
	Key	string	`json:"key" name:"key" title:"Key"`
	Expires	float64	`json:"expires" name:"expires" title:"超时时间(秒)"`
	Length	interface{}	`json:"length,omitempty" name:"length" title:"验证码长度 默认 4"`
}

func (T *CreateTask) GetName() string {
	return "create.json"
}

func (T *CreateTask) GetTitle() string {
	return "创建验证码"
}

