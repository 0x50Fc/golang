package comment

type TrashQueryTask struct {
	Id	interface{}	`json:"id,omitempty" name:"id" title:"评论ID"`
	Pid	interface{}	`json:"pid,omitempty" name:"pid" title:"父级别"`
	Eid	int64	`json:"eid" name:"eid" title:"评论目标ID"`
	Uid	interface{}	`json:"uid,omitempty" name:"uid" title:"用户ID,0不验证"`
	Q	interface{}	`json:"q,omitempty" name:"q" title:"内容模糊查询"`
	P	interface{}	`json:"p,omitempty" name:"p" title:"分页位置, 从1开始, 0 不处理分页"`
	N	interface{}	`json:"n,omitempty" name:"n" title:"分页大小，默认 20"`
}

func (T *TrashQueryTask) GetName() string {
	return "trash/query.json"
}

func (T *TrashQueryTask) GetTitle() string {
	return "查询回收站中的评论"
}

