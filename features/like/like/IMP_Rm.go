package like

import (
	"bytes"
	"fmt"

	"github.com/hailongz/golang/db"
	"github.com/hailongz/golang/dynamic"
	"github.com/hailongz/golang/micro"
)

func (S *Service) Rm(app micro.IContext, task *RmTask) (*Like, error) {

	conn, prefix, err := app.GetDB("wd")

	if err != nil {
		return nil, err
	}

	prefix = Prefix(app, prefix, task.Tid)

	//查询是否存在
	like := Like{}

	iid := dynamic.IntValue(task.Iid, 0)

	sql := bytes.NewBuffer(nil)

	args := []interface{}{}

	sql.WriteString(" WHERE tid = ? and uid = ? AND iid=?")

	args = append(args, task.Tid, task.Uid, iid)

	//先查询是否赞过
	rs, err := db.Query(conn, &like, prefix, sql.String(), args...)

	if err != nil {
		return nil, err
	}

	defer rs.Close()

	if rs.Next() {
		scaner := db.NewScaner(&like)

		err = scaner.Scan(rs)

		if err != nil {
			return nil, err
		}
	} else {
		return nil, micro.NewError(ERROR_NOT_FOUND, "未找到赞")
	}

	//删除赞
	_, err = db.Delete(conn, &like, prefix)

	if err != nil {
		return nil, err
	}

	//清除缓存
	{
		cache, err := app.GetCache("default")

		if err == nil {

			cache.Del(fmt.Sprintf("%d", task.Tid))

		}
	}

	//MQ
	app.SendMessage(task.GetName(), &like)

	return &like, nil
}
