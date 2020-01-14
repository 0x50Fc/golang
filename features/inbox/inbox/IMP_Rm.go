package inbox

import (
	"fmt"

	"github.com/hailongz/golang/db"
	"github.com/hailongz/golang/micro"
)

func (S *Service) Rm(app micro.IContext, task *RmTask) (*Inbox, error) {

	conn, prefix, err := app.GetDB("wd")

	if err != nil {
		return nil, err
	}

	prefix = Prefix(app, prefix, task.Uid)

	v := Inbox{}

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
		return nil, micro.NewError(ERROR_NOT_FOUND, "未找到收件")
	}

	_, err = db.Delete(conn, &v, prefix)

	if err != nil {
		return nil, err
	}

	{
		//清除缓存
		cache, err := app.GetCache("default")
		if err == nil {
			cache.Del(fmt.Sprintf("%d", task.Uid))
		}
	}

	app.SendMessage(task.GetName(), &v)

	return &v, nil
}
