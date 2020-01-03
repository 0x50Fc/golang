package oss

type RmTask struct {
	Name	interface{}	`json:"name,omitempty" name:"name" title:"配置名称"`
	Key	string	`json:"key" name:"key" title:"Key"`
}

func (T *RmTask) GetName() string {
	return "rm.json"
}

func (T *RmTask) GetTitle() string {
	return "删除"
}

