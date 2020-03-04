package job

type AppSetTask struct {
	Id	int64	`json:"id" name:"id" title:"应用ID"`
	Alias	interface{}	`json:"alias,omitempty" name:"alias" title:"别名"`
	Type	interface{}	`json:"type,omitempty" name:"type" title:"类型"`
	Content	interface{}	`json:"content,omitempty" name:"content" title:"内容"`
	Options	interface{}	`json:"options,omitempty" name:"options" title:"其他数据 JSON 叠加数据"`
}

func (T *AppSetTask) GetName() string {
	return "app/set.json"
}

func (T *AppSetTask) GetTitle() string {
	return "修改应用"
}

