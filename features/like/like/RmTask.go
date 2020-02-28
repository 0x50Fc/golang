package like

type RmTask struct {
	Tid	int64	`json:"tid" name:"tid" title:"目标ID"`
	Iid	interface{}	`json:"iid,omitempty" name:"iid" title:"项ID 默认 0"`
	Uid	int64	`json:"uid" name:"uid" title:"用户ID"`
}

func (T *RmTask) GetName() string {
	return "rm.json"
}

func (T *RmTask) GetTitle() string {
	return "取消赞"
}

