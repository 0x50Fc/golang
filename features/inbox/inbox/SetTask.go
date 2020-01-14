package inbox

type SetTask struct {
	Id	int64	`json:"id" name:"id" title:"ID"`
	Uid	int64	`json:"uid" name:"uid" title:"用户ID"`
	Fuid	interface{}	`json:"fuid,omitempty" name:"fuid" title:"发布者ID"`
	Options	interface{}	`json:"options,omitempty" name:"options" title:"其他数据 JSON 叠加数据"`
}

func (T *SetTask) GetName() string {
	return "set.json"
}

func (T *SetTask) GetTitle() string {
	return "修改"
}

