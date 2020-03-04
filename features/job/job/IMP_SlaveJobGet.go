package job

import (
	"bytes"
	"time"

	"github.com/hailongz/golang/db"
	"github.com/hailongz/golang/micro"
)

func (S *Service) SlaveJobGet(app micro.IContext, task *SlaveJobGetTask) (*Job, error) {

	conn, prefix, err := app.GetDB("wd")

	if err != nil {
		return nil, err
	}

	slave := Slave{}
	v := Job{}

	err = db.Transaction(conn, func(conn db.Database) error {

		p, err := db.Get(conn, &slave, prefix, " WHERE token=?", task.Token)

		if err != nil {
			return err
		}

		if p == nil {
			return micro.NewError(ERROR_NOT_FOUND, "未找到主机")
		}

		slave.Etime = time.Now().Unix() + task.Expires
		slave.State = SlaveState_Running

		_, err = db.UpdateWithKeys(conn, &slave, prefix, map[string]bool{"etime": true, "state": true})

		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	err = db.Transaction(conn, func(conn db.Database) error {

		sql := bytes.NewBuffer(nil)
		args := []interface{}{}

		sql.WriteString(" WHERE sid=0 AND state=? AND stime<=?")

		args = append(args, JobState_None, time.Now().Unix())

		if slave.Prefix != "" {
			sql.WriteString(" AND alias LIKE ?")
			args = append(args, slave.Prefix)
		}

		sql.WriteString(" ORDER BY stime ASC , id ASC LIMIT 1")

		p, err := db.Get(conn, &v, prefix, sql.String(), args...)

		if err != nil {
			return err
		}

		if p == nil {
			return micro.NewError(ERROR_NOT_FOUND, "未找到工作")
		}

		v.State = JobState_Running
		v.Sid = slave.Id

		_, err = db.UpdateWithKeys(conn, &v, prefix, map[string]bool{"state": true, "sid": true})

		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return &v, nil
}
