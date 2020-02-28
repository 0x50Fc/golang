package comment

type CreateTask struct {
	Pid	int64	`json:"pid" name:"pid" title:"父级ID"`
	Eid	int64	`json:"eid" name:"eid" title:"评论目标ID"`
	Uid	int64	`json:"uid" name:"uid" title:"用户ID"`
	Body	string	`json:"body" name:"body" title:"内容"`
	Options	interface{}	`json:"options,omitempty" name:"options" title:"其他选项 JSON 叠加"`
}

func (T *CreateTask) GetName() string {
	return "create.json"
}

func (T *CreateTask) GetTitle() string {
	return "创建评论"
}

