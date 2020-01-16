package doc

import (
	"bytes"
	"fmt"

	"github.com/hailongz/golang/db"
	"github.com/hailongz/golang/dynamic"
	"github.com/hailongz/golang/micro"
)

func (S *Service) Query(app micro.IContext, task *QueryTask) (*DocQueryData, error) {

	p := int32(dynamic.IntValue(task.P, 1))
	n := int32(dynamic.IntValue(task.N, 20))

	if n < 1 {
		n = 20
	}

	if p < 1 {
		p = 1
	}

	data := DocQueryData{}

	if task.P != nil {

		countTask := CountTask{}

		countTask.Uid = task.Uid
		countTask.Type = task.Type
		countTask.Ext = task.Ext
		countTask.Pid = task.Pid
		countTask.Q = task.Q
		countTask.Prefix = task.Prefix

		countData, err := S.Count(app, &countTask)

		if err != nil {
			return nil, err
		}

		data.Page = &Page{
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

	if task.Type != nil {
		sql.WriteString(" AND (type & ?) != 0")
		args = append(args, task.Type)
	}

	if task.Ext != nil {
		sql.WriteString(" AND (type = 2 OR ext = ?)")
		args = append(args, task.Ext)
	}

	if task.Pid != nil {
		sql.WriteString(" AND pid = ?")
		args = append(args, task.Pid)
	}

	if task.Prefix != nil {
		sql.WriteString(" AND path LIKE ?")
		args = append(args, fmt.Sprintf("%s%%", task.Prefix))
	}

	if task.Q != nil {
		q := fmt.Sprintf("%%%s%%", task.Q)
		sql.WriteString(" AND (title LIKE ? OR keyword LIKE ?)")
		args = append(args, q, q)
	}

	sql.WriteString(fmt.Sprintf(" ORDER BY type DESC, id DESC LIMIT %d,%d", (p-1)*n, n))

	v := Doc{}

	data.Items = []*Doc{}

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

		item := Doc{}
		item = v

		data.Items = append(data.Items, &item)
	}

	return &data, nil
}
