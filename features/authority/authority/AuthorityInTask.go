package authority

type AuthorityInTask struct {
	Uid	int64	`json:"uid" name:"uid" title:"用户ID"`
	Path	string	`json:"path" name:"path" title:"资源路径"`
}

func (T *AuthorityInTask) GetName() string {
	return "authority/in.json"
}

func (T *AuthorityInTask) GetTitle() string {
	return "验证用户权限"
}

