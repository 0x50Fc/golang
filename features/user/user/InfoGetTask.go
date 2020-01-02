package user

type InfoGetTask struct {
	Uid	int64	`json:"uid" name:"uid" title:"用户ID"`
	Key	string	`json:"key" name:"key" title:"key"`
	Type	interface{}	`json:"type,omitempty" name:"type" title:"类型"`
	Value	interface{}	`json:"value,omitempty" name:"value" title:"内容"`
}

func (T *InfoGetTask) GetName() string {
	return "info/get.json"
}

func (T *InfoGetTask) GetTitle() string {
	return "获取用户信息"
}

