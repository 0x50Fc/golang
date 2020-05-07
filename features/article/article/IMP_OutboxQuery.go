package article

import (
	"bytes"
	"fmt"

	"github.com/hailongz/golang/db"
	"github.com/hailongz/golang/dynamic"
	"github.com/hailongz/golang/micro"
)

func (S *Service) OutboxQuery(app micro.IContext, task *OutboxQueryTask) (*OutboxQueryData, error) {

	p := int32(dynamic.IntValue(task.P, 1))
	n := int32(dynamic.IntValue(task.N, 20))

	if n < 1 {
		n = 20
	}

	if p < 1 {
		p = 1
	}

	data := OutboxQueryData{}

	if task.P != nil {

		countTask := OutboxCountTask{}

		countTask.Uid = task.Uid
		countTask.IsPublished = task.IsPublished
		countTask.Q = task.Q

		countData, err := S.OutboxCount(app, &countTask)

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
		sql.WriteString(" AND body LIKE ?")
		args = append(args, fmt.Sprintf("%%%s%%", task.Q))
	}

	sql.WriteString(fmt.Sprintf(" ORDER BY id DESC LIMIT %d,%d", (p-1)*n, n))

	v := Outbox{}

	data.Items = []*Outbox{}

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

		item := Outbox{}
		item = v

		data.Items = append(data.Items, &item)
	}

	return &data, nil
}
