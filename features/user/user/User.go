package user

import (
	"github.com/hailongz/golang/db"
)

type User struct {
	db.Object
	Name	string	`json:"name" name:"name" title:"用户名" length:"128" unique:"ASC"`
	Nick	string	`json:"nick" name:"nick" title:"昵称" length:"128" index:"ASC"`
	Password	string	`json:"-" name:"password" title:"密码" length:"32"`
	Ctime	int64	`json:"ctime" name:"ctime" title:"创建时间"`
}

func (O *User) GetName() string {
	return "user"
}

func (O *User) GetTitle() string {
	return "用户"
}

