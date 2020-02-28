package authority

type AuthorityUserResAddTask struct {
	Uid	int64	`json:"uid" name:"uid" title:"用户ID"`
	ResId	int64	`json:"resId" name:"resid" title:"资源ID"`
	Options	interface{}	`json:"options,omitempty" name:"options" title:"其他选项 JSON 叠加"`
}

func (T *AuthorityUserResAddTask) GetName() string {
	return "authority/user/res/add.json"
}

func (T *AuthorityUserResAddTask) GetTitle() string {
	return "用户添加资源"
}

