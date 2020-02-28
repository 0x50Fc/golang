package like

type GetTask struct {
	Tid	int64	`json:"tid" name:"tid" title:"目标ID"`
	Iid	interface{}	`json:"iid,omitempty" name:"iid" title:"项ID 默认 0"`
	Uid	int64	`json:"uid" name:"uid" title:"用户ID"`
}

func (T *GetTask) GetName() string {
	return "get.json"
}

func (T *GetTask) GetTitle() string {
	return "获取赞"
}

