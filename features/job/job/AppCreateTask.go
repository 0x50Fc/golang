package job

type AppCreateTask struct {
	Alias	string	`json:"alias" name:"alias" title:"别名"`
	Type	string	`json:"type" name:"type" title:"类型"`
	Content	string	`json:"content" name:"content" title:"内容"`
	Options	interface{}	`json:"options,omitempty" name:"options" title:"其他数据 JSON 叠加数据"`
}

func (T *AppCreateTask) GetName() string {
	return "app/create.json"
}

func (T *AppCreateTask) GetTitle() string {
	return "创建应用"
}

