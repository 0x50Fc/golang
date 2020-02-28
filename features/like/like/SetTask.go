package like

type SetTask struct {
	Tid	int64	`json:"tid" name:"tid" title:"目标ID"`
	Iid	interface{}	`json:"iid,omitempty" name:"iid" title:"项ID 默认 0"`
	Uid	int64	`json:"uid" name:"uid" title:"用户ID"`
	Options	interface{}	`json:"options,omitempty" name:"options" title:"其他数据 JSON 叠加数据"`
}

func (T *SetTask) GetName() string {
	return "set.json"
}

func (T *SetTask) GetTitle() string {
	return "点赞"
}

