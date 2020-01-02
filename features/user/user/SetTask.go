package user

type SetTask struct {
	Id	int64	`json:"id" name:"id" title:"用户ID"`
	Name	interface{}	`json:"name,omitempty" name:"name" title:"用户名"`
	Nick	interface{}	`json:"nick,omitempty" name:"nick" title:"昵称"`
	Password	interface{}	`json:"password,omitempty" name:"password" title:"密码"`
}

func (T *SetTask) GetName() string {
	return "set.json"
}

func (T *SetTask) GetTitle() string {
	return "修改用户"
}

