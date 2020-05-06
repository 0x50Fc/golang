package lookin

type CountTask struct {
	Tid	int64	`json:"tid" name:"tid" title:"目标"`
	Iid	interface{}	`json:"iid,omitempty" name:"iid" title:"项ID 默认 0"`
	Uid	interface{}	`json:"uid,omitempty" name:"uid" title:"用户ID"`
	Fuid	interface{}	`json:"fuid,omitempty" name:"fuid" title:"用户ID"`
	Flevel	interface{}	`json:"flevel,omitempty" name:"flevel" title:"好友级别，多个逗号分割"`
	GroupBy	interface{}	`json:"groupBy,omitempty" name:"groupby" title:"分组"`
}

func (T *CountTask) GetName() string {
	return "count.json"
}

func (T *CountTask) GetTitle() string {
	return "数量"
}

