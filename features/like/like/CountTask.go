package like

type CountTask struct {
	Tid	int64	`json:"tid" name:"tid" title:"目标"`
	Iid	interface{}	`json:"iid,omitempty" name:"iid" title:"项ID 默认 0"`
	Uid	interface{}	`json:"uid,omitempty" name:"uid" title:"用户ID"`
}

func (T *CountTask) GetName() string {
	return "count.json"
}

func (T *CountTask) GetTitle() string {
	return "数量"
}

