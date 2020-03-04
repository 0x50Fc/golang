package wx

import (
	"bytes"

	"github.com/hailongz/golang/db"
	"github.com/hailongz/golang/micro"
)

func (S *Service) Set(app micro.IContext, task *SetTask) (*User, error) {

	conn, prefix, err := app.GetDB("wd")

	if err != nil {
		return nil, err
	}

	sql := bytes.NewBuffer(nil)

	v := User{}

	args := []interface{}{}

	sql.WriteString("UPDATE ")
	sql.WriteString(db.TableName(prefix, &v))
	sql.WriteString(" SET state=? WHERE ")

	args = append(args, task.State)

	if task.Type != nil && task.Appid != nil && task.Openid != nil {
		sql.WriteString("type=? AND appid=? AND openid=?")
		args = append(args, task.Type, task.Appid, task.Openid)
	} else if task.Unionid != nil {
		sql.WriteString("unionid=?")
		args = append(args, task.Unionid)
	} else {
		return nil, micro.NewError(ERROR_NOT_FOUND, "未找到 openid 或 unionid")
	}

	_, err = conn.Exec(sql.String(), args...)

	if err != nil {
		return nil, err
	}

	app.SendMessage(task.GetName(), task)

	return nil, nil
}
