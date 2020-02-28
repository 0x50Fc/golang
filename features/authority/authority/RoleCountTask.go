package authority

type RoleCountTask struct {
	Prefix	interface{}	`json:"prefix,omitempty" name:"prefix" title:"前缀"`
}

func (T *RoleCountTask) GetName() string {
	return "role/count.json"
}

func (T *RoleCountTask) GetTitle() string {
	return "角色数量"
}

