package doc

type CreateTask struct {
	Pid	interface{}	`json:"pid,omitempty" name:"pid" title:"父级ID"`
	Title	string	`json:"title" name:"title" title:"标题"`
	Uid	int64	`json:"uid" name:"uid" title:"用户ID"`
	Type	interface{}	`json:"type,omitempty" name:"type" title:"类型"`
	Ext	interface{}	`json:"ext,omitempty" name:"ext" title:"扩展名"`
	Keyword	interface{}	`json:"keyword,omitempty" name:"keyword" title:"搜索关键字"`
	Options	interface{}	`json:"options,omitempty" name:"options" title:"其他数据  JSON"`
}

func (T *CreateTask) GetName() string {
	return "create.json"
}

func (T *CreateTask) GetTitle() string {
	return "创建"
}

