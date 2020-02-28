package comment

import (
	"fmt"
	"time"

	"bytes"
	"github.com/hailongz/golang/db"
	"github.com/hailongz/golang/dynamic"
	"github.com/hailongz/golang/json"
	"github.com/hailongz/golang/micro"
)

func (S *Service) Create(app micro.IContext, task *CreateTask) (*Comment, error) {

	conn, prefix, err := app.GetDB("wd")

	if err != nil {
		return nil, err
	}

	prefix = Prefix(app, prefix, task.Eid)

	v := Comment{}

	//插入
	v.Pid = task.Pid
	v.Eid = task.Eid
	v.Uid = task.Uid
	v.Body = dynamic.StringValue(task.Body, "string")
	v.Ctime = time.Now().Unix()
	if task.Options != nil {
		text := dynamic.StringValue(task.Options, "")
		json.Unmarshal([]byte(text), &v.Options)
	}

	if task.Pid != 0 {
		c := Comment{}
		//获取pid的path
		sql := bytes.NewBuffer(nil)

		args := []interface{}{}

		sql.WriteString(" WHERE state=? AND eid=? AND id = ?")

		args = append(args, CommentState_None, task.Eid, task.Pid)

		rs, err := db.Query(conn, &c, prefix, sql.String(), args...)

		if err != nil {
			return nil, err
		}

		scaner := db.NewScaner(&c)

		for rs.Next() {

			err = scaner.Scan(rs)

			if err != nil {
				return nil, err
			}

		}

		v.Path = fmt.Sprintf("%s%d|", c.Path, task.Pid)
	}
	_, err = db.Insert(conn, &v, prefix)

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

	// MQ 消息
	app.SendMessage(task.GetName(), &v)

	return &v, nil
}
