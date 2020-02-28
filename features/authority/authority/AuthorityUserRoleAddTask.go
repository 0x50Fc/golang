package authority

type AuthorityUserRoleAddTask struct {
	Uid	int64	`json:"uid" name:"uid" title:"用户ID"`
	RoleId	int64	`json:"roleId" name:"roleid" title:"角色ID"`
	Options	interface{}	`json:"options,omitempty" name:"options" title:"其他选项 JSON 叠加"`
}

func (T *AuthorityUserRoleAddTask) GetName() string {
	return "authority/user/role/add.json"
}

func (T *AuthorityUserRoleAddTask) GetTitle() string {
	return "用户添加角色"
}

