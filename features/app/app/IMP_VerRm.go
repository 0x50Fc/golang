package app

import (
	"fmt"

	"github.com/hailongz/golang/db"
	"github.com/hailongz/golang/micro"
)

func (S *Service) VerRm(app micro.IContext, task *VerRmTask) (*Ver, error) {
	v := Ver{}

	conn, prefix, err := app.GetDB("wd")

	if err != nil {
		return nil, err
	}

	prefix = Prefix(app, prefix, task.Appid)

	p, err := db.Get(conn, &v, prefix, "WHERE appid=? AND ver=?", task.Appid, task.Ver)

	if p == nil {
		return nil, micro.NewError(ERROR_NOT_FOUND, "未找到App版本")
	}

	_, err = db.Delete(conn, &v, prefix)

	if err != nil {
		return nil, err
	}

	cache, _ := app.GetCache("default")

	if cache != nil {
		cache.DelItem(fmt.Sprintf("%d", task.Appid), fmt.Sprintf("%d", task.Ver))
	}

	return &v, nil
}
