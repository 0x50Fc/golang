package doc

type CountTask struct {
	Uid	int64	`json:"uid" name:"uid" title:"用户ID"`
	Pid	interface{}	`json:"pid,omitempty" name:"pid" title:"父级ID"`
	Type	interface{}	`json:"type,omitempty" name:"type" title:"类型"`
	Ext	interface{}	`json:"ext,omitempty" name:"ext" title:"扩展名"`
	Prefix	interface{}	`json:"prefix,omitempty" name:"prefix" title:"路径前缀"`
	Q	interface{}	`json:"q,omitempty" name:"q" title:"搜索关键字"`
}

func (T *CountTask) GetName() string {
	return "count.json"
}

func (T *CountTask) GetTitle() string {
	return "数量"
}

