package top

import (
	"fmt"

	"github.com/hailongz/golang/db"
	"github.com/hailongz/golang/micro"
)

func (S *Service) Get(app micro.IContext, task *GetTask) (*Top, error) {

	conn, prefix, err := app.GetDB("rd")

	if err != nil {
		return nil, err
	}

	prefix = fmt.Sprintf("%s%s_", prefix, task.Name)

	v := Top{}

	err = func() error {

		rs, err := db.Query(conn, &v, prefix, " WHERE tid=?", task.Tid)

		if err != nil {
			return err
		}

		defer rs.Close()

		if rs.Next() {

			scaner := db.NewScaner(&v)

			err = scaner.Scan(rs)

			if err != nil {
				return err
			}
		} else {
			return micro.NewError(ERROR_NOT_FOUND, "未找到推荐项")
		}

		return nil
	}()

	if err != nil {
		return nil, err
	}

	return &v, nil
}
