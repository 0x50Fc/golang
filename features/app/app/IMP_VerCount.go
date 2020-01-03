package app

import (
	"github.com/hailongz/golang/db"
	"github.com/hailongz/golang/micro"
)

func (S *Service) VerCount(app micro.IContext, task *VerCountTask) (*CountData, error) {

	v := Ver{}

	conn, prefix, err := app.GetDB("rd")

	if err != nil {
		return nil, err
	}

	prefix = Prefix(app, prefix, task.Appid)

	n, err := db.Count(conn, &v, prefix, "WHERE appid=?", task.Appid)

	if err != nil {
		return nil, err
	}

	return &CountData{Total: int32(n)}, nil
}
