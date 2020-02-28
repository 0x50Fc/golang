package authority

type RoleQueryTask struct {
	Prefix	interface{}	`json:"prefix,omitempty" name:"prefix" title:"前缀"`
	P	interface{}	`json:"p,omitempty" name:"p" title:"分页位置, 从1开始, 0 不处理分页"`
	N	interface{}	`json:"n,omitempty" name:"n" title:"分页大小，默认 20"`
}

func (T *RoleQueryTask) GetName() string {
	return "role/query.json"
}

func (T *RoleQueryTask) GetTitle() string {
	return "查询角色"
}

