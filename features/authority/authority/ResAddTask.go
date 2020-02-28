package authority

type ResAddTask struct {
	Path	string	`json:"path" name:"path" title:"资源路径"`
	Title	interface{}	`json:"title,omitempty" name:"title" title:"说明"`
	Options	interface{}	`json:"options,omitempty" name:"options" title:"其他选项 JSON 叠加"`
}

func (T *ResAddTask) GetName() string {
	return "res/add.json"
}

func (T *ResAddTask) GetTitle() string {
	return "添加资源"
}

