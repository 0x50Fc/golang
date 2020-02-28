package comment

type TrashAddTask struct {
	Id	int64	`json:"id" name:"id" title:"评论ID"`
	Eid	int64	`json:"eid" name:"eid" title:"评论目标ID"`
}

func (T *TrashAddTask) GetName() string {
	return "trash/add.json"
}

func (T *TrashAddTask) GetTitle() string {
	return "添加评论到回收站"
}

