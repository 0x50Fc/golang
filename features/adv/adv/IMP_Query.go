package adv

import (
	"bytes"
	"fmt"

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

		countTask.Id = task.Id
		countTask.Channel = task.Channel
		countTask.Stime = task.Stime
		countTask.Etime = task.Etime

		countData, err := S.Count(app, &countTask)

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

	conn, prefix, err := app.GetDB("wd")
	if err != nil {
		return nil, err
	}

	sql := bytes.NewBuffer(nil)

	args := []interface{}{}

	sql.WriteString(" WHERE 1=1 ")

	if task.Id != nil {
		sql.WriteString(" AND id =?")
		args = append(args, task.Id)
	}

	if task.Channel != nil {
		sql.WriteString(" AND channel =?")
		args = append(args, task.Channel)
	}

	if task.Stime != nil {
		sql.WriteString(" AND starttime <=?")
		args = append(args, task.Stime)
	}

	if task.Etime != nil {
		sql.WriteString(" AND endtime >=?")
		args = append(args, task.Etime)
	}

	sql.WriteString(fmt.Sprintf(" ORDER BY id DESC LIMIT %d,%d", (p-1)*n, n))

	app.Println("advquerypath", sql)
	app.Println(args)

	v := Adv{}

	data.Items = []*Adv{}

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

		item := Adv{}
		item = v

		data.Items = append(data.Items, &item)
	}

	return &data, nil
}
