package authority

type RoleSetTask struct {
	Id	int64	`json:"id" name:"id" title:"资源ID"`
	Name	interface{}	`json:"name,omitempty" name:"name" title:"角色名"`
	Title	interface{}	`json:"title,omitempty" name:"title" title:"说明"`
	Options	interface{}	`json:"options,omitempty" name:"options" title:"其他选项 JSON 叠加"`
}

func (T *RoleSetTask) GetName() string {
	return "role/set.json"
}

func (T *RoleSetTask) GetTitle() string {
	return "修改角色"
}

