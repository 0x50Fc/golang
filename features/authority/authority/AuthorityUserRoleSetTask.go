package authority

type AuthorityUserRoleSetTask struct {
	Uid	int64	`json:"uid" name:"uid" title:"用户ID"`
	Role	string	`json:"role" name:"role" title:"角色名, 多个逗号分割"`
}

func (T *AuthorityUserRoleSetTask) GetName() string {
	return "authority/user/role/set.json"
}

func (T *AuthorityUserRoleSetTask) GetTitle() string {
	return "用户设置角色"
}

