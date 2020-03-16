package geetest

type RegTask struct {
	CaptchaId	string	`json:"captchaId" name:"captchaid" title:"极验ID"`
	Key	string	`json:"key" name:"key" title:"验证唯一键"`
	Expires	int32	`json:"expires" name:"expires" title:"超时时间"`
}

func (T *RegTask) GetName() string {
	return "reg.json"
}

func (T *RegTask) GetTitle() string {
	return "验证初始化"
}

