package notice

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
		countTask.Ids = task.Ids
		countTask.Type = task.Type
		countTask.Q = task.Q
		countTask.TopId = task.TopId

		countData, err := S.Count(app, &countTask)

		if err != nil {
			return nil, err
		}

		data.Page = &TopPage{
			Total: countData.Total,
			P:     p,
			N:     n,
			Count: countData.Total / n,
		}

		if countData.Total%n != 0 {
			data.Page.Count++
		}

		if task.TopId != nil {
			data.Page.TopId = dynamic.IntValue(task.TopId, 0)
		}

	}

	conn, prefix, err := app.GetDB("rd")

	if err != nil {
		return nil, err
	}

	prefix = Prefix(app, prefix, task.Uid)

	sql := bytes.NewBuffer(nil)

	args := []interface{}{}

	sql.WriteString(" WHERE uid=?")

	args = append(args, task.Uid)

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

	if task.Type != nil {
		vs := strings.Split(dynamic.StringValue(task.Type, ""), ",")
		sql.WriteString(" AND type IN (")
		for i, v := range vs {
			if i != 0 {
				sql.WriteString(",")
			}
			sql.WriteString("?")
			args = append(args, v)
		}
		sql.WriteString(")")
	}

	if task.Fid != nil {

		sql.WriteString(" AND fid IN (")

		ids := strings.Split(dynamic.StringValue(task.Fid, ""), ",")
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

	if task.Iid != nil {

		sql.WriteString(" AND iid IN (")

		ids := strings.Split(dynamic.StringValue(task.Iid, ""), ",")
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

	if task.Q != nil {
		sql.WriteString(" AND title LIKE ?")
		args = append(args, fmt.Sprintf("%%%s%%", task.Q))
	}

	if task.TopId != nil {
		sql.WriteString(" AND id<=?")
		args = append(args, task.TopId)
	}

	sql.WriteString(fmt.Sprintf(" ORDER BY id DESC LIMIT %d,%d", (p-1)*n, n))

	v := Notice{}

	data.Items = []*Notice{}

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

		item := Notice{}
		item = v

		data.Items = append(data.Items, &item)
	}

	if p == 1 && len(data.Items) > 0 && data.Page != nil {
		data.Page.TopId = data.Items[0].Id
	}

	return &data, nil
}
