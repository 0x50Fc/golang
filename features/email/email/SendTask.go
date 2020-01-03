package email

type SendTask struct {
	To          string `json:"to,omitempty" title:"收件人"`
	Subject     string `json:"subject,omitempty" title:"标题"`
	Body        string `json:"body,omitempty" title:"内容"`
	ContentType string `json:"contentType,omitempty" title:"内容类型"`
}

func (T *SendTask) GetName() string {
	return "send.json"
}

func (T *SendTask) GetTitle() string {
	return "发送邮件"
}
