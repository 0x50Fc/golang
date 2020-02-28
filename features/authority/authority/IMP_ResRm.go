package authority

import (
	"github.com/hailongz/golang/db"
	"github.com/hailongz/golang/micro"
)

func (S *Service) ResRm(app micro.IContext, task *ResRmTask) (*Res, error) {

	conn, prefix, err := app.GetDB("wd")

	if err != nil {
		return nil, err
	}

	v := Res{}

	err = db.Transaction(conn, func(conn db.Database) error {

		p, err := db.Get(conn, &v, prefix, " WHERE id=?", task.Id)

		if err != nil {
			return err
		}

		if p == nil {
			return micro.NewError(ERROR_NOT_FOUND, "未找到授权资源")
		}

		_, err = db.Delete(conn, &v, prefix)

		if err != nil {
			return err
		}

		a := Authority{}

		_, err = db.DeleteWithSQL(conn, &a, prefix, " WHERE resid=?", v.Id)

		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return &v, nil
}
