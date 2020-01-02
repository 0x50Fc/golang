package user

type CreateTask struct {
	Name	string	`json:"name" name:"name" title:"用户名"`
	Nick	interface{}	`json:"nick,omitempty" name:"nick" title:"昵称"`
	Password	interface{}	`json:"password,omitempty" name:"password" title:"密码"`
}

func (T *CreateTask) GetName() string {
	return "create.json"
}

func (T *CreateTask) GetTitle() string {
	return "创建用户"
}

