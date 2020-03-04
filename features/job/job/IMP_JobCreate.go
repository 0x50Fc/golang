package job

import (
	"encoding/json"
	"time"

	"github.com/hailongz/golang/db"
	"github.com/hailongz/golang/dynamic"
	"github.com/hailongz/golang/micro"
)

func (S *Service) JobCreate(app micro.IContext, task *JobCreateTask) (*Job, error) {

	conn, prefix, err := app.GetDB("wd")

	if err != nil {
		return nil, err
	}

	v := Job{}

	v.Type = int32(dynamic.IntValue(task.Type, 0))
	v.Appid = dynamic.IntValue(task.Appid, 0)
	v.Uid = dynamic.IntValue(task.Uid, 0)

	if task.Alias != nil {
		v.Alias = dynamic.StringValue(task.Alias, "")
	} else if v.Appid != 0 {
		appTask := AppGetTask{}
		appTask.Id = v.Appid
		app, err := S.AppGet(app, &appTask)
		if err != nil {
			return nil, err
		}
		v.Alias = app.Alias
	}

	v.Ctime = time.Now().Unix()
	v.Stime = v.Ctime

	if task.Stime != nil {
		v.Stime = dynamic.IntValue(task.Stime, v.Stime)
	}

	if v.Stime == 0 {
		v.Stime = v.Ctime
	}

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
