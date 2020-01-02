package user

type InfoSetTask struct {
	Uid	int64	`json:"uid" name:"uid" title:"用户ID"`
	Key	string	`json:"key" name:"key" title:"key"`
	Type	interface{}	`json:"type,omitempty" name:"type" title:"类型"`
	Value	interface{}	`json:"value,omitempty" name:"value" title:"内容"`
}

func (T *InfoSetTask) GetName() string {
	return "info/set.json"
}

func (T *InfoSetTask) GetTitle() string {
	return "修改用户信息"
}

