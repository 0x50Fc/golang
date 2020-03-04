package wx

import (
	"time"

	"github.com/hailongz/golang/db"
	"github.com/hailongz/golang/micro"
)

func (S *Service) AppFormidUse(app micro.IContext, task *AppFormidUseTask) (*AppFormIdUseData, error) {

	conn, prefix, err := app.GetDB("wd")

	if err != nil {
		return nil, err
	}

	v := FormId{}

	now := time.Now().Unix()

	err = db.Transaction(conn, func(conn db.Database) error {

		p, err := db.Get(conn, &v, prefix, " WHERE appid=? AND openid=? AND etime>? ORDER BY id ASC LIMIT 1", task.Appid, task.Openid, now)

		if err != nil {
			return err
		}

		if p == nil {
			return micro.NewError(ERROR_FORM_ID, "未找到可用 formid")
		}

		_, err = db.DeleteWithSQL(conn, &v, prefix, " WHERE appid=? AND openid=? AND (etime<=? OR id=?)", task.Appid, task.Openid, now, v.Id)

		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return &AppFormIdUseData{Formid: v.Formid}, nil
}
