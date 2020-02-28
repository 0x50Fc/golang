package authority

type AuthorityRoleResRmTask struct {
	RoleId	int64	`json:"roleId" name:"roleid" title:"角色"`
	ResId	int64	`json:"resId" name:"resid" title:"资源ID"`
}

func (T *AuthorityRoleResRmTask) GetName() string {
	return "authority/role/res/rm.json"
}

func (T *AuthorityRoleResRmTask) GetTitle() string {
	return "角色删除资源"
}

