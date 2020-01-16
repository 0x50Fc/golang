package auth

type CreateTask struct {
	Key	string	`json:"key" name:"key" title:"键值"`
	Type	interface{}	`json:"type,omitempty" name:"type" title:"类型"`
	Value	string	`json:"value" name:"value" title:"值"`
	Expires	int32	`json:"expires" name:"expires" title:"超时时间(秒)"`
}

func (T *CreateTask) GetName() string {
	return "create.json"
}

func (T *CreateTask) GetTitle() string {
	return "创建"
}

