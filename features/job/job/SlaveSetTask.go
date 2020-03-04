package job

type SlaveSetTask struct {
	Id	int64	`json:"id" name:"id" title:"主机ID"`
	Prefix	interface{}	`json:"prefix,omitempty" name:"prefix" title:"别名前缀 默认不限制"`
	Token	interface{}	`json:"token,omitempty" name:"token" title:"生产授权token"`
	Options	interface{}	`json:"options,omitempty" name:"options" title:"其他数据 JSON 叠加数据"`
	State	interface{}	`json:"state,omitempty" name:"state" title:"状态"`
	Etime	interface{}	`json:"etime,omitempty" name:"etime" title:"超时时间"`
}

func (T *SlaveSetTask) GetName() string {
	return "slave/set.json"
}

func (T *SlaveSetTask) GetTitle() string {
	return "修改主机"
}

