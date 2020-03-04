package wx

type CheckTextTask struct {
	Appid   string `json:"appid" name:"appid" title:"appid"`
	Content string `json:"content" name:"content" title:"消息"`
}

func (T *CheckTextTask) GetName() string {
	return "check/text.json"
}

func (T *CheckTextTask) GetTitle() string {
	return "检查文本"
}
