package lookin

type AddTask struct {
	Tid	int64	`json:"tid" name:"tid" title:"目标ID"`
	Iid	interface{}	`json:"iid,omitempty" name:"iid" title:"项ID"`
	Uid	int64	`json:"uid" name:"uid" title:"用户ID"`
	Fcode	interface{}	`json:"fcode,omitempty" name:"fcode" title:"代码"`
	Fuid	interface{}	`json:"fuid,omitempty" name:"fuid" title:"好友ID"`
	Options	interface{}	`json:"options,omitempty" name:"options" title:"其他数据 JSON 叠加数据"`
}

func (T *AddTask) GetName() string {
	return "add.json"
}

func (T *AddTask) GetTitle() string {
	return "添加在看好友"
}

