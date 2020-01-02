package user

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/hailongz/golang/db"
	"github.com/hailongz/golang/dynamic"
	"github.com/hailongz/golang/micro"
)

func (S *Service) Query(app micro.IContext, task *QueryTask) (*QueryData, error) {
	conn, prefix, err := app.GetDB("rd")

	if err != nil {
		return nil, err
	}
	q := QueryData{}
	v := User{}
	var args []interface{}
	p := int32(dynamic.IntValue(task.P, 1))
	n := int32(dynamic.IntValue(task.N, 20))
	if n < 1 {
		n = 20
	}
	if p < 1 {
		p = 1
	}
	sql := bytes.NewBuffer(nil)
	sql.WriteString(" WHERE 1=1")
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
		q.Page = &QueryDataPage{
			Total: 0,
			P:     p,
			N:     n,
			Count: 0,
		}
		return &q, nil
	}

	if task.P != nil && task.N != nil {
		sql.WriteString(fmt.Sprintf(" ORDER BY id ASC LIMIT %d,%d", (p-1)*n, n))
	}

	rs, err := db.Query(conn, &v, prefix, sql.String(), args...)
	if err != nil {
		return nil, err
	}
	defer rs.Close()
	scanner := db.NewScaner(&v)
	for rs.Next() {
		err = scanner.Scan(rs)
		if err != nil {
			return nil, err
		}
		item := User{}
		item = v
		q.Items = append(q.Items, &item)
		q.Page = &QueryDataPage{
			Total: int32(count),
			P:     p,
			N:     n,
			Count: (int32(count) + n - 1) / n,
		}

	}
	return &q, nil
}
