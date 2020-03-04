package job

import (
	"time"

	"github.com/hailongz/golang/db"
	"github.com/hailongz/golang/micro"
)

func (S *Service) LogAdd(app micro.IContext, task *LogAddTask) (*Log, error) {

	conn, prefix, err := app.GetDB("wd")

	if err != nil {
		return nil, err
	}

	v := Log{}

	v.JobId = task.JobId
	v.Type = task.Type
	v.Appid = task.Appid
	v.Sid = task.Sid
	v.Body = task.Body
	v.Ctime = time.Now().Unix()

	prefix = Prefix(app, prefix, task.JobId)

	_, err = db.Insert(conn, &v, prefix)

	if err != nil {
		return nil, err
	}

	// MQ 消息
	app.SendMessage(task.GetName(), &v)

	return &v, nil
}
