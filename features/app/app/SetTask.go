package app

type SetTask struct {
	Id	int64	`json:"id" name:"id" title:"ID"`
	Uid	interface{}	`json:"uid,omitempty" name:"uid" title:"用户ID"`
	Ver	interface{}	`json:"ver,omitempty" name:"ver" title:"默认版本号"`
	Title	interface{}	`json:"title,omitempty" name:"title" title:"标题"`
	Options	interface{}	`json:"options,omitempty" name:"options" title:"其他数据"`
}

func (T *SetTask) GetName() string {
	return "set.json"
}

func (T *SetTask) GetTitle() string {
	return "修改"
}

