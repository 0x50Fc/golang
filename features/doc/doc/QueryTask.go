package doc

type QueryTask struct {
	Uid	int64	`json:"uid" name:"uid" title:"用户ID"`
	Pid	interface{}	`json:"pid,omitempty" name:"pid" title:"父级ID"`
	Type	interface{}	`json:"type,omitempty" name:"type" title:"类型"`
	Ext	interface{}	`json:"ext,omitempty" name:"ext" title:"扩展名"`
	Prefix	interface{}	`json:"prefix,omitempty" name:"prefix" title:"路径前缀"`
	Q	interface{}	`json:"q,omitempty" name:"q" title:"搜索关键字"`
	P	interface{}	`json:"p,omitempty" name:"p" title:"分页位置, 从1开始, 0 不处理分页"`
	N	interface{}	`json:"n,omitempty" name:"n" title:"分页大小，默认 20"`
}

func (T *QueryTask) GetName() string {
	return "query.json"
}

func (T *QueryTask) GetTitle() string {
	return "查询"
}

