package user

import (
	"time"

	"github.com/hailongz/golang/db"
	"github.com/hailongz/golang/dynamic"
	"github.com/hailongz/golang/micro"
)

func (S *Service) Create(app micro.IContext, task *CreateTask) (*User, error) {

	if task.Name == "" {
		return nil, micro.NewError(ERROR_NOT_FOUND, "未找到用户名")
	}

	conn, prefix, err := app.GetDB("wd")

	if err != nil {
		return nil, err
	}

	v := User{}
	err = db.Transaction(conn, func(conn db.Database) error {

		p, err := db.Get(conn, &v, prefix, " WHERE name=?", task.Name)

		if err != nil {
			return err
		}

		if p != nil {
			return micro.NewError(ERROR_USER_ISEXIT, "用户名已存在")
		}

		if task.Nick != nil {

			nick := dynamic.StringValue(task.Nick, "")

			if nick != "" {

				p, err = db.Get(conn, &v, prefix, " WHERE nick=?", task.Nick)

				if err != nil {
					return err
				}

				if p != nil {
					return micro.NewError(ERROR_USER_ISEXIT, "昵称已存在")
				}
			}

			v.Nick = nick
		}

		v.Name = task.Name
		if task.Password == nil {
			v.Password = EncPassword(NewPassword())
		} else {
			v.Password = EncPassword(dynamic.StringValue(task.Password, ""))
		}
		v.Ctime = time.Now().Unix()

		_, err = db.Insert(conn, &v, prefix)

		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	// MQ 消息
	app.SendMessage(task.GetName(), &v)

	return &v, nil
}
