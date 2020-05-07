package article

type OutboxRmTask struct {
	Id	int64	`json:"id" name:"id" title:"草稿ID"`
	Uid	int64	`json:"uid" name:"uid" title:"用户ID"`
}

func (T *OutboxRmTask) GetName() string {
	return "outbox/rm.json"
}

func (T *OutboxRmTask) GetTitle() string {
	return "删除草稿/动态"
}

