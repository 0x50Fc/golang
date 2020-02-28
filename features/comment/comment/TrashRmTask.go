package comment

type TrashRmTask struct {
	Id	int64	`json:"id" name:"id" title:"评论ID"`
	Eid	int64	`json:"eid" name:"eid" title:"评论目标ID"`
}

func (T *TrashRmTask) GetName() string {
	return "trash/rm.json"
}

func (T *TrashRmTask) GetTitle() string {
	return "恢复评论"
}

