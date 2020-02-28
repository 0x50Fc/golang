package authority

type AuthorityUserResRmTask struct {
	Uid	int64	`json:"uid" name:"uid" title:"用户ID"`
	ResId	int64	`json:"resId" name:"resid" title:"资源ID"`
}

func (T *AuthorityUserResRmTask) GetName() string {
	return "authority/user/res/rm.json"
}

func (T *AuthorityUserResRmTask) GetTitle() string {
	return "用户删除资源"
}

