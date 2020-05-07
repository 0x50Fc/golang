package article

type OutboxGetTask struct {
	Id	int64	`json:"id" name:"id" title:"草稿ID"`
	Uid	int64	`json:"uid" name:"uid" title:"用户ID"`
}

func (T *OutboxGetTask) GetName() string {
	return "outbox/get.json"
}

func (T *OutboxGetTask) GetTitle() string {
	return "获取草稿"
}

