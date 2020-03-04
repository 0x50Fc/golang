package job

type SlaveCreateTask struct {
	Prefix	interface{}	`json:"prefix,omitempty" name:"prefix" title:"别名前缀 默认不限制"`
	Options	interface{}	`json:"options,omitempty" name:"options" title:"其他数据 JSON 叠加数据"`
}

func (T *SlaveCreateTask) GetName() string {
	return "slave/create.json"
}

func (T *SlaveCreateTask) GetTitle() string {
	return "创建主机"
}

