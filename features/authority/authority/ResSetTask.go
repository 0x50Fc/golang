package authority

type ResSetTask struct {
	Id	int64	`json:"id" name:"id" title:"资源ID"`
	Path	interface{}	`json:"path,omitempty" name:"path" title:"资源路径"`
	Title	interface{}	`json:"title,omitempty" name:"title" title:"说明"`
	Options	interface{}	`json:"options,omitempty" name:"options" title:"其他选项 JSON 叠加"`
}

func (T *ResSetTask) GetName() string {
	return "res/set.json"
}

func (T *ResSetTask) GetTitle() string {
	return "修改资源"
}

