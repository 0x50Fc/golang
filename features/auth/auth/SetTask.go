package auth

type SetTask struct {
	Key	string	`json:"key" name:"key" title:"键值"`
	Type	interface{}	`json:"type,omitempty" name:"type" title:"类型"`
	Value	interface{}	`json:"value,omitempty" name:"value" title:"值"`
	Expires	interface{}	`json:"expires,omitempty" name:"expires" title:"超时时间(秒)"`
}

func (T *SetTask) GetName() string {
	return "set.json"
}

func (T *SetTask) GetTitle() string {
	return "修改"
}

