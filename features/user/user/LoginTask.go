package user

type LoginTask struct {
	Name	string	`json:"name" name:"name" title:"用户名"`
	Password	string	`json:"password" name:"password" title:"密码"`
}

func (T *LoginTask) GetName() string {
	return "login.json"
}

func (T *LoginTask) GetTitle() string {
	return "登录"
}

