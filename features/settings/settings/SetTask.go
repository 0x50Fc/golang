package settings

type SetTask struct {
	Name	string	`json:"name" name:"name" title:"名称"`
	Options	interface{}	`json:"options,omitempty" name:"options" title:"其他选项 JSON 叠加"`
}

func (T *SetTask) GetName() string {
	return "set.json"
}

func (T *SetTask) GetTitle() string {
	return "设置"
}

