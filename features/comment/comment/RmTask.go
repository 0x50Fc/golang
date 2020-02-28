package comment

type RmTask struct {
	Id	int64	`json:"id" name:"id" title:"评论ID"`
	Eid	int64	`json:"eid" name:"eid" title:"评论目标ID"`
	Uid	interface{}	`json:"uid,omitempty" name:"uid" title:"用户ID,0不验证"`
}

func (T *RmTask) GetName() string {
	return "rm.json"
}

func (T *RmTask) GetTitle() string {
	return "删除评论"
}

