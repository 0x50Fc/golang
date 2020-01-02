package user

import (
	"github.com/hailongz/golang/db"
	"github.com/hailongz/golang/micro"
)

func (S *Service) Login(app micro.IContext, task *LoginTask) (*User, error) {

	conn, prefix, err := app.GetDB("rd")

	if err != nil {
		return nil, err
	}

	v := User{}

	p, err := db.Get(conn, &v, prefix, " WHERE name=? AND password=?", task.Name, EncPassword(task.Password))

	if err != nil {
		return nil, err
	}

	if p == nil {
		return nil, micro.NewError(ERROR_NOT_FOUND, "未找到用户")
	}

	return &v, nil
}
