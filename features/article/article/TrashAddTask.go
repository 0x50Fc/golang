package article

type TrashAddTask struct {
	Id	int64	`json:"id" name:"id" title:"动态ID"`
}

func (T *TrashAddTask) GetName() string {
	return "trash/add.json"
}

func (T *TrashAddTask) GetTitle() string {
	return "添加动态到回收站"
}

