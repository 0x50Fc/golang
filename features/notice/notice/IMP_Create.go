package notice

import (
	"time"

	"github.com/hailongz/golang/db"
	"github.com/hailongz/golang/dynamic"
	"github.com/hailongz/golang/json"
	"github.com/hailongz/golang/micro"
)

func (S *Service) Create(app micro.IContext, task *CreateTask) (*Notice, error) {

	conn, prefix, err := app.GetDB("wd")

	if err != nil {
		return nil, err
	}

	prefix = Prefix(app, prefix, task.Uid)

	v := Notice{}

	v.Uid = task.Uid
	v.Type = int32(dynamic.IntValue(task.Type, 0))
	v.Fid = dynamic.IntValue(task.Fid, 0)
	v.Iid = dynamic.IntValue(task.Iid, 0)
	v.Body = task.Body
	v.Ctime = time.Now().Unix()

	if task.Options != nil {
		text := dynamic.StringValue(task.Options, "")
		json.Unmarshal([]byte(text), &v.Options)
	}

	_, err = db.Insert(conn, &v, prefix)

	if err != nil {
		return nil, err
	}

	// MQ 消息
	app.SendMessage(task.GetName(), &v)

	return &v, nil
}
