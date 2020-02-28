package feed

import (
	"fmt"

	"github.com/hailongz/golang/db"
	"github.com/hailongz/golang/micro"
)

func (S *Service) TrashRm(app micro.IContext, task *TrashRmTask) (*Feed, error) {

	conn, fix, err := app.GetDB("wd")

	if err != nil {
		return nil, err
	}

	prefix := Prefix(app, fix, task.Id)

	v := Feed{}

	rs, err := db.Query(conn, &v, prefix, " WHERE id=?", task.Id)

	if err != nil {
		return nil, err
	}

	defer rs.Close()

	if rs.Next() {

		scaner := db.NewScaner(&v)

		err = scaner.Scan(rs)

		if err != nil {
			return nil, err
		}

	} else {
		return nil, micro.NewError(ERROR_NOT_FOUND, "未找到动态")
	}

	v.State = FeedState_None

	_, err = db.UpdateWithKeys(conn, &v, prefix, map[string]bool{"state": true})

	if err != nil {
		return nil, err
	}

	{

		// 缓存
		cli, prefix, err := app.GetRedis("default")

		if err != nil {
			return nil, err
		}
		_, _ = cli.Del(fmt.Sprintf("%s%d", prefix, v.Id)).Result()

	}

	app.SendMessage(task.GetName(), &v)

	return &v, nil
}
