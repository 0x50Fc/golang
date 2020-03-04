package job

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/hailongz/golang/db"
	"github.com/hailongz/golang/dynamic"
	"github.com/hailongz/golang/micro"
)

func (S *Service) JobQuery(app micro.IContext, task *JobQueryTask) (*JobQueryData, error) {
	p := int32(dynamic.IntValue(task.P, 1))
	n := int32(dynamic.IntValue(task.N, 20))

	if n < 1 {
		n = 20
	}

	if p < 1 {
		p = 1
	}

	data := JobQueryData{}

	if task.P != nil {

		countTask := JobCountTask{}

		countTask.Alias = task.Alias
		countTask.Type = task.Type
		countTask.Prefix = task.Prefix
		countTask.Appid = task.Appid
		countTask.Uid = task.Uid
		countTask.Sid = task.Sid

		countData, err := S.JobCount(app, &countTask)

		if err != nil {
			return nil, err
		}

		data.Page = &Page{
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

	if task.Uid != nil {
		sql.WriteString(" AND uid=?")
		args = append(args, task.Uid)
	}

	if task.Appid != nil {
		sql.WriteString(" AND appid=?")
		args = append(args, task.Appid)
	}

	if task.Sid != nil {
		sql.WriteString(" AND sid=?")
		args = append(args, task.Sid)
	}

	sql.WriteString(fmt.Sprintf(" ORDER BY id DESC LIMIT %d,%d", (p-1)*n, n))

	v := Job{}

	data.Items = []*Job{}

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

		item := Job{}
		item = v

		data.Items = append(data.Items, &item)
	}

	return &data, nil
}
