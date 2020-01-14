package inbox

import (
	"bytes"
	"fmt"

	"github.com/hailongz/golang/db"
	"github.com/hailongz/golang/micro"
)

func (S *Service) Clean(app micro.IContext, task *CleanTask) (interface{}, error) {

	conn, prefix, err := app.GetDB("wd")

	if err != nil {
		return nil, err
	}

	prefix = Prefix(app, prefix, task.Uid)

	v := Inbox{}

	sql := bytes.NewBuffer(nil)

	args := []interface{}{}

	sql.WriteString(" WHERE uid=?")

	args = append(args, task.Uid)

	if task.Fuid != nil {
		sql.WriteString(" AND fuid=?")
		args = append(args, task.Fuid)
	}

	if task.Mid != nil {
		sql.WriteString(" AND mid=?")
		args = append(args, task.Mid)
	}

	if task.Iid != nil {
		sql.WriteString(" AND iid=?")
		args = append(args, task.Iid)
	}

	_, err = db.DeleteWithSQL(conn, &v, prefix, sql.String(), args...)

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

	// MQ 消息
	app.SendMessage(task.GetName(), &v)

	return nil, nil
}
