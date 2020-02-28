package authority

type AuthorityCountTask struct {
	Uid    interface{} `json:"uid,omitempty" name:"uid" title:"用户ID"`
	RoleId interface{} `json:"roleId,omitempty" name:"roleid" title:"角色ID"`
	ResId  interface{} `json:"resId,omitempty" name:"resid" title:"资源ID"`
}

func (T *AuthorityCountTask) GetName() string {
	return "authority/count.json"
}

func (T *AuthorityCountTask) GetTitle() string {
	return "授权数量"
}
