package article

type TrashRmTask struct {
	Id	int64	`json:"id" name:"id" title:"动态ID"`
}

func (T *TrashRmTask) GetName() string {
	return "trash/rm.json"
}

func (T *TrashRmTask) GetTitle() string {
	return "恢复动态"
}

