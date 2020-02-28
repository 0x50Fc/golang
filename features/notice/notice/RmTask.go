package notice

type RmTask struct {
	Id	int64	`json:"id" name:"id" title:"ID"`
	Uid	int64	`json:"uid" name:"uid" title:"用户ID"`
}

func (T *RmTask) GetName() string {
	return "rm.json"
}

func (T *RmTask) GetTitle() string {
	return "删除"
}

