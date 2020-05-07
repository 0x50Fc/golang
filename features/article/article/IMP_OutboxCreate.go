package article

import (
	"encoding/json"
	"time"

	"github.com/hailongz/golang/db"
	"github.com/hailongz/golang/dynamic"
	"github.com/hailongz/golang/micro"
)

func (S *Service) OutboxCreate(app micro.IContext, task *OutboxCreateTask) (*Outbox, error) {

	conn, prefix, err := app.GetDB("wd")

	if err != nil {
		return nil, err
	}

	prefix = Prefix(app, prefix, task.Uid)

	v := Outbox{}

	v.Uid = task.Uid
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
