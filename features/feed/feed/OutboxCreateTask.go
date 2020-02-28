package feed

type OutboxCreateTask struct {
	Uid	int64	`json:"uid" name:"uid" title:"用户ID"`
	Body	string	`json:"body" name:"body" title:"内容"`
	Options	interface{}	`json:"options,omitempty" name:"options" title:"其他数据 JSON 叠加数据"`
}

func (T *OutboxCreateTask) GetName() string {
	return "outbox/create.json"
}

func (T *OutboxCreateTask) GetTitle() string {
	return "创建草稿"
}

