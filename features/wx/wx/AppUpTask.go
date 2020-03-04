package wx

type AppUpTask struct {
	Appid	string	`json:"appid" name:"appid" title:"appid"`
	Type	interface{}	`json:"type,omitempty" name:"type" title:"文件类型, 默认 image"`
	Name	string	`json:"name" name:"name" title:"文件名"`
	Content	string	`json:"content" name:"content" title:"文件内容 base64"`
}

func (T *AppUpTask) GetName() string {
	return "app/up.json"
}

func (T *AppUpTask) GetTitle() string {
	return "上传媒体文件"
}

