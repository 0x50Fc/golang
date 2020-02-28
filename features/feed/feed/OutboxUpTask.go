package feed

type OutboxUpTask struct {
	Id	int64	`json:"id,omitempty" title:"草稿ID"`
	Uid	int64	`json:"uid,omitempty" title:"用户ID"`
	Body	interface{}	`json:"body,omitempty" title:"内容"`
	Options	interface{}	`json:"options,omitempty" title:"其他数据 JSON 叠加数据"`
}

func (T *OutboxUpTask) GetName() string {
	return "outbox/up.json"
}

func (T *OutboxUpTask) GetTitle() string {
	return "发布动态"
}

