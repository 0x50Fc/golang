package auth

type GetTask struct {
	Key	string	`json:"key,omitempty" title:"键值"`
	Type	interface{}	`json:"type,omitempty" title:"类型"`
}

func (T *GetTask) GetName() string {
	return "get.json"
}

func (T *GetTask) GetTitle() string {
	return "获取"
}

