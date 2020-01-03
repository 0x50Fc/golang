package app

type RmTask struct {
	Id	int64	`json:"id" name:"id" title:"ID"`
	Uid	interface{}	`json:"uid,omitempty" name:"uid" title:"用户ID"`
}

func (T *RmTask) GetName() string {
	return "rm.json"
}

func (T *RmTask) GetTitle() string {
	return "删除"
}

