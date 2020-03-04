package job

import (
	"encoding/json"
	"time"

	"github.com/hailongz/golang/db"
	"github.com/hailongz/golang/dynamic"
	"github.com/hailongz/golang/micro"
)

func (S *Service) AppCreate(app micro.IContext, task *AppCreateTask) (*App, error) {

	conn, prefix, err := app.GetDB("wd")

	if err != nil {
		return nil, err
	}

	v := App{}

	v.Alias = task.Alias
	v.Type = task.Type
	v.Content = task.Content
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
