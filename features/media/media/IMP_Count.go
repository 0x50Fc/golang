package media

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/hailongz/golang/db"
	"github.com/hailongz/golang/dynamic"
	"github.com/hailongz/golang/micro"
)

func (S *Service) Count(app micro.IContext, task *CountTask) (*CountData, error) {

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

	sql.WriteString(" WHERE 1")

	if task.Uid != nil {
		sql.WriteString(" AND uid=?")
		args = append(args, task.Uid)
	}

	if task.Type != nil {

		sql.WriteString(" AND type IN (")

		ids := strings.Split(dynamic.StringValue(task.Type, ""), ",")
		i := 0

		for _, id := range ids {
			if id != "" {
				if i != 0 {
					sql.WriteString(",")
				}
				sql.WriteString("?")
				args = append(args, id)
				i++
			}
		}

		sql.WriteString(")")

	}

	if task.Uid != nil {
		sql.WriteString(" AND uid=?")
		args = append(args, task.Uid)
	}

	if task.Q != nil {
		sql.WriteString(" AND (title LIKE ? OR keyword LIKE ?)")
		q := fmt.Sprintf("%%%s%%", task.Q)
		args = append(args, q, q)
	}

	if task.Prefix != nil {
		sql.WriteString(" AND path LIKE ?")
		args = append(args, fmt.Sprintf("%s%%", task.Prefix))
	}

	count, err := db.Count(conn, &v, name, sql.String(), args...)

	if err != nil {
		return nil, err
	}

	return &CountData{Total: int32(count)}, nil
}
