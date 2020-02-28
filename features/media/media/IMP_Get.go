package media

import (
	"bytes"
	"fmt"

	"github.com/hailongz/golang/db"
	"github.com/hailongz/golang/dynamic"
	"github.com/hailongz/golang/micro"
)

func (S *Service) Get(app micro.IContext, task *GetTask) (*Media, error) {

	conn, name, err := app.GetDB("rd")

	if err != nil {
		return nil, err
	}

	if task.Region == nil {
		name = fmt.Sprintf("%s%s", name, dynamic.StringValue(task.Name, ""))
	} else {
		name = fmt.Sprintf("%s%d_%s", name, dynamic.IntValue(task.Region, 0), dynamic.StringValue(task.Name, ""))
	}

	app.Println("[NAME]", name)

	v := Media{}

	sql := bytes.NewBuffer(nil)

	args := []interface{}{}

	sql.WriteString(" WHERE id=?")

	args = append(args, task.Id)

	if task.Uid != nil {
		sql.WriteString(" AND uid=?")
		args = append(args, task.Uid)
	}

	p, err := db.Get(conn, &v, name, sql.String(), args...)

	if err != nil {
		return nil, err
	}

	if p == nil {
		return nil, micro.NewError(ERROR_NOT_FOUND, "未找到媒体对象")
	}

	return &v, nil
}
