package app

type CreateTask struct {
	Uid	int64	`json:"uid" name:"uid" title:"用户ID"`
	Title	interface{}	`json:"title,omitempty" name:"title" title:"标题"`
	Options	interface{}	`json:"options,omitempty" name:"options" title:"其他数据"`
}

func (T *CreateTask) GetName() string {
	return "create.json"
}

func (T *CreateTask) GetTitle() string {
	return "创建"
}

