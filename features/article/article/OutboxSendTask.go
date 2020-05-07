package article

type OutboxSendTask struct {
	Id	int64	`json:"id" name:"id" title:"草稿ID"`
	Uid	int64	`json:"uid" name:"uid" title:"用户ID"`
}

func (T *OutboxSendTask) GetName() string {
	return "outbox/send.json"
}

func (T *OutboxSendTask) GetTitle() string {
	return "发布"
}

