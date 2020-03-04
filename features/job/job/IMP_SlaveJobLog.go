package job

import (
	"github.com/hailongz/golang/db"
	"github.com/hailongz/golang/micro"
)

func (S *Service) SlaveJobLog(app micro.IContext, task *SlaveJobLogTask) (interface{}, error) {

	conn, prefix, err := app.GetDB("rd")

	if err != nil {
		return nil, err
	}

	slave := Slave{}

	p, err := db.Get(conn, &slave, prefix, " WHERE token=?", task.Token)

	if err != nil {
		return nil, err
	}

	if p == nil {
		return nil, micro.NewError(ERROR_NOT_FOUND, "未找到主机")
	}

	logTask := LogAddTask{}
	logTask.JobId = task.JobId
	logTask.Appid = task.Appid
	logTask.Sid = slave.Id
	logTask.Type = task.Type
	logTask.Body = task.Body
	_, err = S.LogAdd(app, &logTask)

	if err != nil {
		return nil, err
	}

	return nil, nil
}
