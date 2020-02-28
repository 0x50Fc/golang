package authority

import (
	"bytes"
	"fmt"

	"github.com/hailongz/golang/db"
	"github.com/hailongz/golang/dynamic"
	"github.com/hailongz/golang/micro"
)

func (S *Service) RoleQuery(app micro.IContext, task *RoleQueryTask) (*RoleQueryData, error) {

	p := int32(dynamic.IntValue(task.P, 1))
	n := int32(dynamic.IntValue(task.N, 20))

	if n < 1 {
		n = 20
	}

	if p < 1 {
		p = 1
	}

	data := RoleQueryData{}

	if task.P != nil {

		countTask := RoleCountTask{}

		countTask.Prefix = task.Prefix

		countData, err := S.RoleCount(app, &countTask)

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

	if task.Prefix != nil {
		sql.WriteString(" AND name LIKE ?")
		args = append(args, fmt.Sprintf("%s%%", task.Prefix))
	}

	sql.WriteString(fmt.Sprintf(" ORDER BY name ASC,id DESC LIMIT %d,%d", (p-1)*n, n))

	v := Role{}

	data.Items = []*Role{}

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

		item := Role{}
		item = v

		data.Items = append(data.Items, &item)
	}

	return &data, nil
}
