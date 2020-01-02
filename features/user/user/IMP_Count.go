package user

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/hailongz/golang/db"
	"github.com/hailongz/golang/dynamic"
	"github.com/hailongz/golang/micro"
)

func (S *Service) Count(app micro.IContext, task *CountTask) (*CountData, error) {
	conn, prefix, err := app.GetDB("rd")

	if err != nil {
		return nil, err
	}
	q := CountData{}
	v := User{}
	var args []interface{}
	sql := bytes.NewBuffer(nil)
	sql.WriteString(" WHERE 1")
	if task.Ids != nil {
		sql.WriteString(" AND id IN (")
		ids := strings.Split(dynamic.StringValue(task.Ids, ""), ",")
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
	if task.Name != nil {
		sql.WriteString(" AND name=?")
		args = append(args, task.Name)
	}
	if task.Nick != nil {
		sql.WriteString(" AND nick=?")
		args = append(args, task.Nick)
	}
	if task.Q != nil {
		sql.WriteString(" AND nick LIKE ?")
		args = append(args, fmt.Sprintf("%%%s%%", task.Q))
	}
	if task.Prefix != nil {
		sql.WriteString(" AND name LIKE ?")
		args = append(args, fmt.Sprintf("%s%%", task.Prefix))
	}
	if task.Suffix != nil {
		sql.WriteString(" AND name LIKE ?")
		args = append(args, fmt.Sprintf("%%%s", task.Prefix))
	}
	count, err := db.Count(conn, &v, prefix, sql.String(), args...)
	if err != nil {
		return nil, err
	}
	if count == 0 {
		count = 0
	}
	q.Total = int32(count)
	return &q, nil
}
