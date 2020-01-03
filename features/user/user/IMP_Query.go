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

	p := int32(dynamic.IntValue(task.P, 1))
	n := int32(dynamic.IntValue(task.N, 20))

	if n < 1 {
		n = 20
	}

	if p < 1 {
		p = 1
	}

	data := QueryData{}

	if task.P != nil {

		countTask := CountTask{}

		countTask.Ids = task.Ids
		countTask.Name = task.Name
		countTask.Q = task.Q
		countTask.Prefix = task.Prefix
		countTask.Suffix = task.Suffix

		countData, err := S.Count(app, &countTask)

		if err != nil {
			return nil, err
		}

		data.Page = &QueryDataPage{
			Total: countData.Total,
			P:     p,
			N:     n,
			Count: countData.Total / n,
		}

		if countData.Total%n != 0 {
			data.Page.Count++
		}

	}

	conn, prefix, err := app.GetDB("rd")

	if err != nil {
		return nil, err
	}

	sql := bytes.NewBuffer(nil)

	args := []interface{}{}

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

	sql.WriteString(fmt.Sprintf(" ORDER BY id DESC LIMIT %d,%d", (p-1)*n, n))

	v := User{}

	data.Items = []*User{}

	rs, err := db.Query(conn, &v, prefix, sql.String(), args...)

	if err != nil {
		return nil, err
	}

	defer rs.Close()

	scaner := db.NewScaner(&v)

	for rs.Next() {

		err = scaner.Scan(rs)

		if err != nil {
			return nil, err
		}

		item := User{}
		item = v

		data.Items = append(data.Items, &item)
	}

	return &data, nil

}
