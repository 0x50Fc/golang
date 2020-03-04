package job

import (
	"github.com/hailongz/golang/db"
	"github.com/hailongz/golang/micro"
)

func (S *Service) JobCancel(app micro.IContext, task *JobCancelTask) (*Job, error) {

	conn, prefix, err := app.GetDB("wd")

	if err != nil {
		return nil, err
	}

	v := Job{}

	err = db.Transaction(conn, func(conn db.Database) error {

		p, err := db.Get(conn, &v, prefix, " WHERE id=?", task.Id)

		if err != nil {
			return err
		}

		if p == nil {
			return micro.NewError(ERROR_NOT_FOUND, "未找到工作")
		}

		if v.State == JobState_None || v.State == JobState_Running {

			v.State = JobState_Cancel

			_, err = db.UpdateWithKeys(conn, &v, prefix, map[string]bool{"state": true})

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
