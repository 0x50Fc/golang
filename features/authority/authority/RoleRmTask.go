package authority

type RoleRmTask struct {
	Id	int64	`json:"id" name:"id" title:"角色ID"`
}

func (T *RoleRmTask) GetName() string {
	return "role/rm.json"
}

func (T *RoleRmTask) GetTitle() string {
	return "删除角色"
}

