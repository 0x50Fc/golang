package feed

type OutboxSendTask struct {
	Id	int64	`json:"id" name:"id" title:"草稿ID"`
	Uid	int64	`json:"uid" name:"uid" title:"用户ID"`
	Body	interface{}	`json:"body,omitempty" name:"body" title:"内容"`
	Options	interface{}	`json:"options,omitempty" name:"options" title:"其他数据 JSON 叠加数据"`
}

func (T *OutboxSendTask) GetName() string {
	return "outbox/send.json"
}

func (T *OutboxSendTask) GetTitle() string {
	return "发布"
}

