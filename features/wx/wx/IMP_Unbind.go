package wx

import (
	"bytes"

	"github.com/hailongz/golang/db"
	"github.com/hailongz/golang/micro"
)

func (S *Service) Unbind(app micro.IContext, task *UnbindTask) (interface{}, error) {
	conn, prefix, err := app.GetDB("wd")

	if err != nil {
		return nil, err
	}

	sql := bytes.NewBuffer(nil)

	v := User{}

	args := []interface{}{}

	sql.WriteString("UPDATE ")
	sql.WriteString(db.TableName(prefix, &v))
	sql.WriteString(" SET uid=0 WHERE uid=?")

	args = append(args, task.Uid)

	if task.Type != nil {
		sql.WriteString(" AND type=?")
		args = append(args, task.Type)
	}

	if task.Appid != nil {
		sql.WriteString(" AND appid=?")
		args = append(args, task.Appid)
	}

	if task.Openid != nil {
		sql.WriteString(" AND openid=?")
		args = append(args, task.Openid)
	}

	if task.Unionid != nil {
		sql.WriteString(" AND unionid=?")
		args = append(args, task.Unionid)
	}

	_, err = conn.Exec(sql.String(), args...)

	if err != nil {
		return nil, err
	}

	return nil, nil
}
