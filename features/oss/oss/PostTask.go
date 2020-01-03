package oss

type PostTask struct {
	Name	interface{}	`json:"name,omitempty" name:"name" title:"配置名称"`
	Key	string	`json:"key" name:"key" title:"Key"`
	Expires	interface{}	`json:"expires,omitempty" name:"expires" title:"超时时间(秒) type == url 使用"`
}

func (T *PostTask) GetName() string {
	return "post.json"
}

func (T *PostTask) GetTitle() string {
	return "上传"
}

