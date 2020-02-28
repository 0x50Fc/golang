package authority

type RoleAddTask struct {
	Name	string	`json:"name" name:"name" title:"角色名"`
	Title	interface{}	`json:"title,omitempty" name:"title" title:"说明"`
	Options	interface{}	`json:"options,omitempty" name:"options" title:"其他选项 JSON 叠加"`
}

func (T *RoleAddTask) GetName() string {
	return "role/add.json"
}

func (T *RoleAddTask) GetTitle() string {
	return "添加角色"
}

