package app

type VerAddTask struct {
	Appid	int64	`json:"appid" name:"appid" title:"appid"`
	Info	interface{}	`json:"info,omitempty" name:"info" title:"INFO"`
	Options	interface{}	`json:"options,omitempty" name:"options" title:"其他数据"`
}

func (T *VerAddTask) GetName() string {
	return "ver/add.json"
}

func (T *VerAddTask) GetTitle() string {
	return "创建版本"
}

