package authority

type RoleGetTask struct {
	Id	interface{}	`json:"id,omitempty" name:"id" title:"角色ID"`
	Name	interface{}	`json:"name,omitempty" name:"name" title:"角色名称"`
}

func (T *RoleGetTask) GetName() string {
	return "role/get.json"
}

func (T *RoleGetTask) GetTitle() string {
	return "获取角色"
}

