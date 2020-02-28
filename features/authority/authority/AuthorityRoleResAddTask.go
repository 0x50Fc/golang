package authority

type AuthorityRoleResAddTask struct {
	RoleId	int64	`json:"roleId" name:"roleid" title:"角色"`
	ResId	int64	`json:"resId" name:"resid" title:"资源ID"`
	Options	interface{}	`json:"options,omitempty" name:"options" title:"其他选项 JSON 叠加"`
}

func (T *AuthorityRoleResAddTask) GetName() string {
	return "authority/role/res/add.json"
}

func (T *AuthorityRoleResAddTask) GetTitle() string {
	return "角色添加资源"
}

