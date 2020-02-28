package authority

type AuthorityRoleResSetTask struct {
	Role	string	`json:"role" name:"role" title:"角色名"`
	Res	string	`json:"res" name:"res" title:"资源路径, 多个逗号分割"`
}

func (T *AuthorityRoleResSetTask) GetName() string {
	return "authority/role/res/set.json"
}

func (T *AuthorityRoleResSetTask) GetTitle() string {
	return "角色设置资源"
}

