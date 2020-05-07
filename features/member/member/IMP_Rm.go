package member

import (
	"fmt"

	"github.com/hailongz/golang/db"
	"github.com/hailongz/golang/micro"
)

func (S *Service) Rm(app micro.IContext, task *RmTask) (*Member, error) {

	conn, prefix, err := app.GetDB("wd")

	if err != nil {
		return nil, err
	}

	v := Member{}

	p, err := db.Get(conn, &v, prefix, "WHERE bid=? AND uid=?", task.Bid, task.Uid)

	if p == nil {
		return nil, micro.NewError(ERROR_NOT_FOUND, "未找到成员")
	}

	_, err = db.Delete(conn, &v, prefix)

	if err != nil {
		return nil, err
	}

	cache, _ := app.GetCache("default")

	if cache != nil {
		cache.Del(fmt.Sprintf("B_%d", v.Bid))
		cache.Del(fmt.Sprintf("U_%d", v.Uid))
		cache.Del("Q")
	}

	app.SendMessage(task.GetName(), &v)

	return &v, nil
}
