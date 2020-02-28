package authority

import (
	"bytes"

	"github.com/hailongz/golang/db"
	"github.com/hailongz/golang/dynamic"
	"github.com/hailongz/golang/micro"
)

func (S *Service) RoleGet(app micro.IContext, task *RoleGetTask) (*Role, error) {

	conn, prefix, err := app.GetDB("rd")

	if err != nil {
		return nil, err
	}

	sql := bytes.NewBuffer(nil)

	args := []interface{}{}

	if task.Id != nil {
		sql.WriteString(" WHERE id=?")
		args = append(args, task.Id)
	} else if task.Name != nil {
		sql.WriteString(" WHERE name=?")
		args = append(args, task.Name)
	} else {
		return nil, micro.NewError(ERROR_NOT_FOUND, "未找到授权角色ID")
	}

	v := Role{}

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

		if task.Name == nil {
			return nil, micro.NewError(ERROR_NOT_FOUND, "未找到授权角色")
		}

		conn, prefix, err := app.GetDB("wd")

		if err != nil {
			return nil, err
		}

		v.Name = dynamic.StringValue(task.Name, "")

		_, err = db.Insert(conn, &v, prefix)

		if err != nil {
			return nil, err
		}

	}

	return &v, nil
}
