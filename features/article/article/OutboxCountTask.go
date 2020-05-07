package article

type OutboxCountTask struct {
	Uid	int64	`json:"uid" name:"uid" title:"用户ID"`
	IsPublished	interface{}	`json:"isPublished,omitempty" name:"ispublished" title:"是否发布"`
	Q	interface{}	`json:"q,omitempty" name:"q" title:"模糊匹配关键字"`
}

func (T *OutboxCountTask) GetName() string {
	return "outbox/count.json"
}

func (T *OutboxCountTask) GetTitle() string {
	return "查询发件箱数量"
}

