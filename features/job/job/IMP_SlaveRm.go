package job

import (
	"github.com/hailongz/golang/db"
	"github.com/hailongz/golang/micro"
)

func (S *Service) SlaveRm(app micro.IContext, task *SlaveRmTask) (*Slave, error) {

	conn, prefix, err := app.GetDB("wd")

	if err != nil {
		return nil, err
	}

	v := Slave{}

	rs, err := db.Query(conn, &v, prefix, " WHERE id=?", task.Id)

	if err != nil {
		return nil, err
	}

	defer rs.Close()

	if rs.Next() {

		scaner := db.NewScaner(&v)

		err = scaner.Scan(rs)

		if err != nil {
			return nil, err
		}

	} else {
		return nil, micro.NewError(ERROR_NOT_FOUND, "未找到主机")
	}

	_, err = db.Delete(conn, &v, prefix)

	if err != nil {
		return nil, err
	}

	app.SendMessage(task.GetName(), &v)

	return &v, nil
}
