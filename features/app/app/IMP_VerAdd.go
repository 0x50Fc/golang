package app

import (
	"time"

	"github.com/hailongz/golang/db"
	"github.com/hailongz/golang/micro"
)

func (S *Service) VerAdd(app micro.IContext, task *VerAddTask) (*Ver, error) {

	a := App{}
	v := Ver{}

	conn, prefix, err := app.GetDB("wd")

	if err != nil {
		return nil, err
	}

	prefix = Prefix(app, prefix, task.Appid)

	err = db.Transaction(conn, func(conn db.Database) error {

		p, err := db.Get(conn, &a, prefix, "WHERE id=?", task.Appid)

		if err != nil {
			return err
		}

		if p == nil {
			return micro.NewError(ERROR_NOT_FOUND, "未找到App")
		}

		v.Appid = task.Appid
		v.Ver = a.LastVer + 1
		v.Info = task.Info
		v.Options = task.Options
		v.Ctime = time.Now().Unix()

		a.LastVer = v.Ver

		_, err = db.UpdateWithKeys(conn, &a, prefix, map[string]bool{"lastver": true})

		if err != nil {
			return err
		}

		_, err = db.Insert(conn, &v, prefix)

		if err != nil {
			return err
		}

		return nil
	})

	return &v, nil
}
