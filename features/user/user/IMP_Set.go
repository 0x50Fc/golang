package user

import (
	"github.com/hailongz/golang/db"
	"github.com/hailongz/golang/dynamic"
	"github.com/hailongz/golang/micro"
)

func (S *Service) Set(app micro.IContext, task *SetTask) (*User, error) {

	conn, prefix, err := app.GetDB("wd")

	if err != nil {
		return nil, err
	}

	v := User{}

	err = db.Transaction(conn, func(conn db.Database) error {

		p, err := db.Get(conn, &v, prefix, " WHERE id=?", task.Id)

		if err != nil {
			return err
		}

		if p == nil {
			return micro.NewError(ERROR_NOT_FOUND, "未找到用户")
		}

		keys := map[string]bool{}

		if task.Name != nil {

			name := dynamic.StringValue(task.Name, "")

			if name == "" {
				return micro.NewError(ERROR_NOT_FOUND, "未找到用户名")
			}

			p, err = db.Get(conn, &v, prefix, " WHERE name=? AND id!=?", name, v.Id)

			if err != nil {
				return err
			}

			if p != nil {
				return micro.NewError(ERROR_USER_ISEXIT, "用户名已存在")
			}

			v.Name = name

			keys["name"] = true
		}

		if task.Nick != nil {

			nick := dynamic.StringValue(task.Nick, "")

			if nick != "" {

				p, err = db.Get(conn, &v, prefix, " WHERE nick=? AND id!=?", nick, v.Id)

				if err != nil {
					return err
				}

				if p != nil {
					return micro.NewError(ERROR_USER_ISEXIT, "昵称已存在")
				}
			}

			v.Nick = nick

			keys["nick"] = true
		}

		if task.Password != nil {
			v.Password = EncPassword(dynamic.StringValue(task.Password, ""))
			keys["password"] = true
		}

		if len(keys) > 0 {

			_, err = db.UpdateWithKeys(conn, &v, prefix, keys)

			if err != nil {
				return err
			}
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return &v, nil
}
