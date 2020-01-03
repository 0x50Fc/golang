package vcode

import (
	"github.com/hailongz/golang/db"
	"github.com/hailongz/golang/micro"
)

func (S *Service) Del(app micro.IContext, task *DelTask) (interface{}, error) {

	conn, prefix, err := app.GetDB("wd")

	if err != nil {
		return nil, err
	}

	v := VCode{}

	_, err = db.DeleteWithSQL(conn, &v, prefix, " WHERE `key`=?", task.Key)

	if err != nil {
		return nil, err
	}
	return nil, nil
}
