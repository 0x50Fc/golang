package article

import (
	"bytes"
	"fmt"

	"github.com/hailongz/golang/db"
	"github.com/hailongz/golang/dynamic"
	"github.com/hailongz/golang/micro"
)

func (S *Service) OutboxCount(app micro.IContext, task *OutboxCountTask) (*OutboxCountData, error) {

	conn, prefix, err := app.GetDB("rd")

	if err != nil {
		return nil, err
	}

	prefix = Prefix(app, prefix, task.Uid)

	sql := bytes.NewBuffer(nil)

	args := []interface{}{}

	sql.WriteString(" WHERE uid=?")

	args = append(args, task.Uid)

	if task.IsPublished != nil {

		if dynamic.BooleanValue(task.IsPublished, false) {
			sql.WriteString(" AND mid=0")
		} else {
			sql.WriteString(" AND mid!=0")
		}

	}

	if task.Q != nil {
		sql.WriteString(" AND title LIKE ?")
		args = append(args, fmt.Sprintf("%%%s%%", task.Q))
	}

	v := Outbox{}

	count, err := db.Count(conn, &v, prefix, sql.String(), args...)

	if err != nil {
		return nil, err
	}

	return &OutboxCountData{Total: int32(count)}, nil
}
