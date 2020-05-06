package lookin

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/hailongz/golang/db"
	"github.com/hailongz/golang/dynamic"
	"github.com/hailongz/golang/micro"
)

func (S *Service) Rm(app micro.IContext, task *RmTask) (interface{}, error) {

	conn, prefix, err := app.GetDB("wd")

	if err != nil {
		return nil, err
	}

	prefix = Prefix(app, prefix, task.Tid)

	iid := dynamic.IntValue(task.Iid, 0)

	v := Lookin{}

	sql := bytes.NewBuffer(nil)

	args := []interface{}{}

	sql.WriteString(" WHERE tid=? AND iid=?")

	args = append(args, task.Tid, iid)

	if task.Uid != nil {
		sql.WriteString(" AND uid=?")
		args = append(args, task.Uid)
	}

	if task.Fuid != nil {
		sql.WriteString(" AND fuid=?")
		args = append(args, task.Fuid)
	}

	if task.Flevel != nil {
		vs := strings.Split(dynamic.StringValue(task.Flevel, ""), ",")
		sql.WriteString(" AND flevel IN(")
		for i, v := range vs {
			if i != 0 {
				sql.WriteString(",")
			}
			sql.WriteString("?")
			args = append(args, v)
		}
		sql.WriteString(")")
	}

	_, err = db.DeleteWithSQL(conn, &v, prefix, sql.String(), args...)

	//清除缓存
	{
		cache, err := app.GetCache("default")

		if err == nil {

			cache.Del(fmt.Sprintf("%d_%d", task.Tid, iid))

		}
	}

	app.SendMessage(task.GetName(), task)

	return nil, nil
}
