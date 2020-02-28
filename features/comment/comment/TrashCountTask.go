package comment

type TrashCountTask struct {
	Id	interface{}	`json:"id,omitempty" name:"id" title:"评论ID"`
	Pid	interface{}	`json:"pid,omitempty" name:"pid" title:"父级别"`
	Eid	int64	`json:"eid" name:"eid" title:"评论目标ID"`
	Uid	interface{}	`json:"uid,omitempty" name:"uid" title:"用户ID,0不验证"`
	Q	interface{}	`json:"q,omitempty" name:"q" title:"内容模糊查询"`
}

func (T *TrashCountTask) GetName() string {
	return "trash/count.json"
}

func (T *TrashCountTask) GetTitle() string {
	return "获取回收站中评论数量"
}

