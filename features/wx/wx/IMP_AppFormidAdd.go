package wx

import (
	"time"

	"github.com/hailongz/golang/db"
	"github.com/hailongz/golang/dynamic"
	"github.com/hailongz/golang/json"
	"github.com/hailongz/golang/micro"
)

func (S *Service) AppFormidAdd(app micro.IContext, task *AppFormidAddTask) (interface{}, error) {

	var items interface{} = nil

	err := json.Unmarshal([]byte(task.Items), &items)

	if err != nil {
		return nil, err
	}

	conn, prefix, err := app.GetDB("wd")

	if err != nil {
		return nil, err
	}

	err = db.Transaction(conn, func(conn db.Database) error {

		var err error = nil

		now := time.Now().Unix()

		dynamic.Each(items, func(_ interface{}, item interface{}) bool {

			formid := dynamic.StringValue(dynamic.Get(item, "formid"), "")

			if formid == "" {
				return true
			}

			etime := dynamic.IntValue(dynamic.Get(item, "etime"), 0)

			if etime <= now {
				return true
			}

			v := FormId{}

			v.Appid = task.Appid
			v.Openid = task.Openid
			v.Formid = formid
			v.Etime = etime

			_, err = db.Insert(conn, &v, prefix)

			if err != nil {
				return false
			}

			return true
		})

		v := FormId{}

		_, err = db.DeleteWithSQL(conn, &v, prefix, " WHERE etime<=? AND appid=? AND openid=?", now, task.Appid, task.Openid)

		return err
	})

	if err != nil {
		return nil, err
	}

	return nil, nil
}
