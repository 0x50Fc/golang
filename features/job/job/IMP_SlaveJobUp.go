package job

import (
	"encoding/json"
	"time"

	"github.com/hailongz/golang/db"
	"github.com/hailongz/golang/dynamic"
	"github.com/hailongz/golang/micro"
)

func (S *Service) SlaveJobUp(app micro.IContext, task *SlaveJobUpTask) (*Job, error) {

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

	err = db.Transaction(conn, func(conn db.Database) error {

		p, err := db.Get(conn, &v, prefix, " WHERE sid=? AND id=?", slave.Id, task.JobId)

		if err != nil {
			return err
		}

		if p == nil {
			return micro.NewError(ERROR_NOT_FOUND, "未找到工作")
		}

		keys := map[string]bool{}

		if task.MaxCount != nil {
			v.MaxCount = int32(dynamic.IntValue(task.MaxCount, int64(v.MaxCount)))
			keys["maxcount"] = true
		}

		if task.Count != nil {
			v.Count = int32(dynamic.IntValue(task.Count, int64(v.Count)))
			keys["count"] = true
		}

		if task.ErrCount != nil {
			v.ErrCount = int32(dynamic.IntValue(task.ErrCount, int64(v.ErrCount)))
			keys["errcount"] = true
		}

		if task.AddCount != nil {
			v.Count = v.Count + int32(dynamic.IntValue(task.AddCount, 0))
			keys["count"] = true
		}

		if task.AddErrCount != nil {
			v.ErrCount = v.ErrCount + int32(dynamic.IntValue(task.AddErrCount, 0))
			keys["errcount"] = true
		}

		if task.Options != nil {

			options := map[string]interface{}{}

			dynamic.Each(v.Options, func(key interface{}, value interface{}) bool {
				options[dynamic.StringValue(key, "")] = value
				return true
			})

			text := dynamic.StringValue(task.Options, "")

			var data interface{} = nil

			json.Unmarshal([]byte(text), &data)

			dynamic.Each(data, func(key interface{}, value interface{}) bool {
				options[dynamic.StringValue(key, "")] = value
				return true
			})

			v.Options = options
			keys["options"] = true
		}

		if task.Done && v.State == JobState_Running {
			v.State = JobState_Finish
			keys["state"] = true
		}

		if len(keys) > 0 {

			_, err = db.UpdateWithKeys(conn, &v, prefix, keys)

			if err != nil {
				return err
			}

		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return &v, nil
}
