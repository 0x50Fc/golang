package media

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

		countTask.Uid = task.Uid
		countTask.Type = task.Type
		countTask.Prefix = task.Prefix
		countTask.Q = task.Q

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

	sql.WriteString(fmt.Sprintf(" ORDER BY id DESC LIMIT %d,%d", (p-1)*n, n))

	v := Media{}

	data.Items = []*Media{}

	rs, err := db.Query(conn, &v, name, sql.String(), args...)

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

		item := Media{}
		item = v

		data.Items = append(data.Items, &item)
	}

	return &data, nil
}
