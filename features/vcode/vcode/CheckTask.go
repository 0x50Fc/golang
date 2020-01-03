package vcode

type CheckTask struct {
	Key	string	`json:"key" name:"key" title:"Key"`
	Code	interface{}	`json:"code,omitempty" name:"code" title:"数字验证码" length:"12"`
	Hash	interface{}	`json:"hash,omitempty" name:"hash" title:"32位 HASH" length:"32"`
}

func (T *CheckTask) GetName() string {
	return "check.json"
}

func (T *CheckTask) GetTitle() string {
	return "校验验证码"
}

