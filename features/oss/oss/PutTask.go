package oss

type PutTask struct {
	Name	interface{}	`json:"name,omitempty" name:"name" title:"配置名称"`
	Key	string	`json:"key" name:"key" title:"Key"`
	Type	interface{}	`json:"type,omitempty" name:"type" title:"类型 默认 url"`
	Content	interface{}	`json:"content,omitempty" name:"content" title:"内容\n当 Type.Text || Type.Base64 时使用"`
	Expires	interface{}	`json:"expires,omitempty" name:"expires" title:"超时时间(秒) type == url 使用"`
}

func (T *PutTask) GetName() string {
	return "put.json"
}

func (T *PutTask) GetTitle() string {
	return "上传"
}

