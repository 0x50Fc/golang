package authority

type AuthorityUserRoleRmTask struct {
	Uid	int64	`json:"uid" name:"uid" title:"用户ID"`
	RoleId	int64	`json:"roleId" name:"roleid" title:"角色ID"`
}

func (T *AuthorityUserRoleRmTask) GetName() string {
	return "authority/user/role/rm.json"
}

func (T *AuthorityUserRoleRmTask) GetTitle() string {
	return "用户删除角色"
}

