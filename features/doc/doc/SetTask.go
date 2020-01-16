package doc

type SetTask struct {
	Id	int64	`json:"id" name:"id" title:"ID"`
	Uid	int64	`json:"uid" name:"uid" title:"用户ID"`
	Pid	interface{}	`json:"pid,omitempty" name:"pid" title:"父级ID"`
	Title	interface{}	`json:"title,omitempty" name:"title" title:"标题"`
	Ext	interface{}	`json:"ext,omitempty" name:"ext" title:"扩展名"`
	Keyword	interface{}	`json:"keyword,omitempty" name:"keyword" title:"搜索关键字"`
	Options	interface{}	`json:"options,omitempty" name:"options" title:"其他数据 JSON"`
	Mtime	interface{}	`json:"mtime,omitempty" name:"mtime" title:"更新最近修改时间"`
	Atime	interface{}	`json:"atime,omitempty" name:"atime" title:"更新最近访问时间"`
}

func (T *SetTask) GetName() string {
	return "set.json"
}

func (T *SetTask) GetTitle() string {
	return "修改"
}

