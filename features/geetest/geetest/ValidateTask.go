package geetest

type ValidateTask struct {
	CaptchaId	string	`json:"captchaId" name:"captchaid" title:"极验ID"`
	Key	string	`json:"key" name:"key" title:"验证唯一键"`
	Challenge	string	`json:"challenge" name:"challenge" title:""`
	Validate	string	`json:"validate" name:"validate" title:""`
	Seccode	string	`json:"seccode" name:"seccode" title:""`
}

func (T *ValidateTask) GetName() string {
	return "validate.json"
}

func (T *ValidateTask) GetTitle() string {
	return "验证"
}

