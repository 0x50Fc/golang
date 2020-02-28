package comment

import (
	"fmt"

	"github.com/hailongz/golang/db"
	"github.com/hailongz/golang/micro"
)

func (S *Service) TrashAdd(app micro.IContext, task *TrashAddTask) (*Comment, error) {

	conn, prefix, err := app.GetDB("wd")

	if err != nil {
		return nil, err
	}

	prefix = Prefix(app, prefix, task.Eid)

	v := Comment{}

	err = func() error {

		rs, err := db.Query(conn, &v, prefix, " WHERE id=? AND eid=?", task.Id, task.Eid)

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
			return micro.NewError(ERROR_NOT_FOUND, "未找到分组")
		}

		return nil
	}()

	if err != nil {
		return nil, err
	}

	v.State = CommentState_Recycle

	_, err = db.UpdateWithKeys(conn, &v, prefix, map[string]bool{"state": true})

	if err != nil {
		return nil, err
	}

	//清除缓存
	{
		cache, err := app.GetCache("default")

		if err == nil {

			cache.Del(fmt.Sprintf("%d", task.Eid))

		}
	}

	app.SendMessage(task.GetName(), &v)

	return &v, nil
}
