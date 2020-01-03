package auth

type CreateTask struct {
	Key     string      `json:"key,omitempty" title:"键值"`
	Type    interface{} `json:"type,omitempty" title:"类型"`
	Value   string      `json:"value,omitempty" title:"值"`
	Expires int32       `json:"expires,omitempty" title:"超时时间(秒)"`
}

func (T *CreateTask) GetName() string {
	return "create.json"
}

func (T *CreateTask) GetTitle() string {
	return "创建"
}
