package wx

type CheckImageTask struct {
	Appid string `json:"appid" name:"appid" title:"appid"`
	Url   string `json:"url" name:"url" title:"消息"`
}

func (T *CheckImageTask) GetName() string {
	return "check/image.json"
}

func (T *CheckImageTask) GetTitle() string {
	return "检查图片"
}
