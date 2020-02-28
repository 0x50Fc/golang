package authority

type AuthorityQueryTask struct {
	Uid	interface{}	`json:"uid,omitempty" name:"uid" title:"用户ID"`
	RoleId	interface{}	`json:"roleId,omitempty" name:"roleid" title:"角色ID"`
	ResId	interface{}	`json:"resId,omitempty" name:"resid" title:"资源ID"`
	P	interface{}	`json:"p,omitempty" name:"p" title:"分页位置, 从1开始, 0 不处理分页"`
	N	interface{}	`json:"n,omitempty" name:"n" title:"分页大小，默认 20"`
}

func (T *AuthorityQueryTask) GetName() string {
	return "authority/query.json"
}

func (T *AuthorityQueryTask) GetTitle() string {
	return "查询授权"
}

