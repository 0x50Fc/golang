package oss

type HasTask struct {
	Name	interface{}	`json:"name,omitempty" name:"name" title:"配置名称"`
	Key	string	`json:"key" name:"key" title:"Key"`
}

func (T *HasTask) GetName() string {
	return "has.json"
}

func (T *HasTask) GetTitle() string {
	return "是否存在"
}

