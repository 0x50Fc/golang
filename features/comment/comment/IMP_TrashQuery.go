package comment

import (
	"bytes"
	"fmt"

	"github.com/hailongz/golang/db"
	"github.com/hailongz/golang/dynamic"
	"github.com/hailongz/golang/micro"
)

func (S *Service) TrashQuery(app micro.IContext, task *TrashQueryTask) (*QueryData, error) {

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

		countTask := TrashCountTask{}

		countTask.Eid = task.Eid

		countData, err := S.TrashCount(app, &countTask)

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

	prefix = Prefix(app, prefix, task.Eid)

	sql := bytes.NewBuffer(nil)

	args := []interface{}{}

	sql.WriteString(" WHERE state=? AND eid=?")

	args = append(args, CommentState_Recycle, task.Eid)

	if task.Id != nil {
		sql.WriteString(" AND id=?")
		args = append(args, task.Id)
	}

	if task.Pid != nil {
		sql.WriteString(" AND pid=?")
		args = append(args, task.Pid)
	}

	if task.Uid != nil {
		sql.WriteString(" AND uid=?")
		args = append(args, task.Uid)
	}

	if task.Q != nil {
		sql.WriteString(" AND body LIKE ?")
		args = append(args, fmt.Sprintf("%%%s%%", task.Q))
	}

	sql.WriteString(fmt.Sprintf(" ORDER BY id DESC LIMIT %d,%d", (p-1)*n, n))

	v := Comment{}

	data.Items = []*Comment{}

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

		item := Comment{}
		item = v

		data.Items = append(data.Items, &item)
	}

	return &data, nil
}
