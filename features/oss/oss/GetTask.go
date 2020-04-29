package oss

type GetTask struct {
	Name	interface{}	`json:"name,omitempty" name:"name" title:"配置名称"`
	Key	string	`json:"key" name:"key" title:"Key"`
	Type	interface{}	`json:"type,omitempty" name:"type" title:"类型 默认 Type.Url"`
	Expires	interface{}	`json:"expires,omitempty" name:"expires" title:"超时时间(秒) 公开读不设置"`
	Header	interface{}	`json:"header,omitempty" name:"header" title:"头 JSON 格式"`
}

func (T *GetTask) GetName() string {
	return "get.json"
}

func (T *GetTask) GetTitle() string {
	return "获取"
}

