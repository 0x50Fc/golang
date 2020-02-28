package authority

import (
	"bytes"

	"github.com/hailongz/golang/db"
	"github.com/hailongz/golang/dynamic"
	"github.com/hailongz/golang/micro"
)

func (S *Service) ResGet(app micro.IContext, task *ResGetTask) (*Res, error) {

	conn, prefix, err := app.GetDB("rd")

	if err != nil {
		return nil, err
	}

	sql := bytes.NewBuffer(nil)

	args := []interface{}{}

	if task.Id != nil {
		sql.WriteString(" WHERE id=?")
		args = append(args, task.Id)
	} else if task.Path != nil {
		sql.WriteString(" WHERE path=?")
		args = append(args, task.Path)
	} else {
		return nil, micro.NewError(ERROR_NOT_FOUND, "未找到授权资源ID")
	}

	v := Res{}

	err = func() error {

		rs, err := db.Query(conn, &v, prefix, sql.String(), args...)

		if err != nil {
			return err
		}

		defer rs.Close()

		if rs.Next() {
			scaner := db.NewScaner(&v)
			err = scaner.Scan(rs)
			if err != nil {
				return err
			}
		}

		return nil
	}()

	if err != nil {
		return nil, err
	}

	if v.Id == 0 {

		if task.Path == nil {
			return nil, micro.NewError(ERROR_NOT_FOUND, "未找到授权资源")
		}

		conn, prefix, err := app.GetDB("wd")

		if err != nil {
			return nil, err
		}

		v.Path = dynamic.StringValue(task.Path, "")

		_, err = db.Insert(conn, &v, prefix)

		if err != nil {
			return nil, err
		}

	}

	return &v, nil
}
