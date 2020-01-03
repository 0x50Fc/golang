package auth

import (
	"github.com/hailongz/golang/db"
	"github.com/hailongz/golang/micro"
)

func (S *Service) Rm(app micro.IContext, task *RmTask) (interface{}, error) {

	key := task.Key

	{
		// 缓存
		cli, prefix, err := app.GetRedis("default")

		if err == nil {
			_, _ = cli.Del(prefix + key).Result()
		} else {
			app.Println("[Redis] [ERROR]", err)
		}

	}

	conn, prefix, err := app.GetDB("rd")

	if err != nil {
		return nil, err
	}

	v := Auth{}

	_, err = db.DeleteWithSQL(conn, &v, prefix, " WHERE `key`=?", key)

	if err != nil {
		return nil, err
	}

	return nil, nil
}
