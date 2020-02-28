package feed

type OutboxSetTask struct {
	Id	int64	`json:"id" name:"id" title:"草稿ID"`
	Uid	int64	`json:"uid" name:"uid" title:"用户ID"`
	Body	interface{}	`json:"body,omitempty" name:"body" title:"内容"`
	Options	interface{}	`json:"options,omitempty" name:"options" title:"其他数据 JSON 叠加数据"`
}

func (T *OutboxSetTask) GetName() string {
	return "outbox/set.json"
}

func (T *OutboxSetTask) GetTitle() string {
	return "修改草稿"
}

