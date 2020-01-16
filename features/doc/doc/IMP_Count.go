package doc

import (
	"bytes"
	"fmt"

	"github.com/hailongz/golang/db"
	"github.com/hailongz/golang/micro"
)

func (S *Service) Count(app micro.IContext, task *CountTask) (*DocCountData, error) {

	conn, prefix, err := app.GetDB("rd")

	if err != nil {
		return nil, err
	}

	prefix = Prefix(app, prefix, task.Uid)

	sql := bytes.NewBuffer(nil)

	args := []interface{}{}

	v := Doc{}

	sql.WriteString(" WHERE uid=?")

	args = append(args, task.Uid)

	if task.Type != nil {
		sql.WriteString(" AND (type & ?) != 0")
		args = append(args, task.Type)
	}

	if task.Ext != nil {
		sql.WriteString(" AND (type = 2 OR ext = ?)")
		args = append(args, task.Ext)
	}

	if task.Pid != nil {
		sql.WriteString(" AND pid = ?")
		args = append(args, task.Pid)
	}

	if task.Prefix != nil {
		sql.WriteString(" AND path LIKE ?")
		args = append(args, fmt.Sprintf("%s%%", task.Prefix))
	}

	if task.Q != nil {
		q := fmt.Sprintf("%%%s%%", task.Q)
		sql.WriteString(" AND (title LIKE ? OR keyword LIKE ?)")
		args = append(args, q, q)
	}

	app.Println("[SQL]", sql.String())

	count, err := db.Count(conn, &v, prefix, sql.String(), args...)

	if err != nil {
		return nil, err
	}

	return &DocCountData{Total: int32(count)}, nil
}
