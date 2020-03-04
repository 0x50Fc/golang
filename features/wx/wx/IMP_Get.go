package wx

import (
	"time"

	"github.com/hailongz/golang/db"
	"github.com/hailongz/golang/dynamic"
	"github.com/hailongz/golang/micro"
)

func (S *Service) Get(app micro.IContext, task *GetTask) (*User, error) {

	conn, prefix, err := app.GetDB("rd")

	if err != nil {
		return nil, err
	}

	v := User{}

	p, err := db.Get(conn, &v, prefix, " WHERE type=? AND appid=? AND openid=?", task.Type, task.Appid, task.Openid)

	if err != nil {
		return nil, err
	}

	if p == nil {

		v.Type = task.Type
		v.Appid = task.Appid
		v.Openid = task.Openid
		v.Ctime = time.Now().Unix()
		v.Mtime = v.Ctime

		err = MP_UpdateUser(app, &v)

		if err != nil {
			return nil, err
		}

		{
			conn, prefix, err := app.GetDB("wd")

			if err != nil {
				return nil, err
			}

			_, err = db.Insert(conn, &v, prefix)

			if err != nil {
				return nil, err
			}
		}

		return &v, nil
	}

	if dynamic.BooleanValue(task.Update, false) {

		err = MP_UpdateUser(app, &v)

		if err != nil {
			return nil, err
		}

		conn, prefix, err = app.GetDB("wd")

		if err != nil {
			return nil, err
		}

		_, err = db.Update(conn, &v, prefix)

		if err != nil {
			return nil, err
		}

	}

	return &v, nil
}
