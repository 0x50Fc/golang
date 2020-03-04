package job

import (
	"bytes"
	"strings"

	"github.com/hailongz/golang/db"
	"github.com/hailongz/golang/dynamic"
	"github.com/hailongz/golang/micro"
)

func (S *Service) AppCount(app micro.IContext, task *AppCountTask) (*CountData, error) {

	conn, prefix, err := app.GetDB("rd")

	if err != nil {
		return nil, err
	}

	sql := bytes.NewBuffer(nil)

	args := []interface{}{}

	sql.WriteString(" WHERE 1")

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

	if task.Prefix != nil {
		sql.WriteString(" AND alias LIKE ?")
		args = append(args, dynamic.StringValue(task.Prefix, "")+"%")
	}

	if task.Alias != nil {
		sql.WriteString(" AND alias=?")
		args = append(args, dynamic.StringValue(task.Alias, ""))
	}

	v := App{}

	count, err := db.Count(conn, &v, prefix, sql.String(), args...)

	if err != nil {
		return nil, err
	}

	return &CountData{Total: int32(count)}, nil
}
