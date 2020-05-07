package article

type OutboxQueryTask struct {
	Uid	int64	`json:"uid" name:"uid" title:"用户ID"`
	IsPublished	interface{}	`json:"isPublished,omitempty" name:"ispublished" title:"是否发布"`
	Q	interface{}	`json:"q,omitempty" name:"q" title:"模糊匹配关键字"`
	P	interface{}	`json:"p,omitempty" name:"p" title:"分页位置, 从1开始, 0 不处理分页"`
	N	interface{}	`json:"n,omitempty" name:"n" title:"分页大小，默认 20"`
}

func (T *OutboxQueryTask) GetName() string {
	return "outbox/query.json"
}

func (T *OutboxQueryTask) GetTitle() string {
	return "查询发件箱"
}

