package wx

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
		countTask.Appid = task.Appid
		countTask.Openid = task.Openid
		countTask.Unionid = task.Unionid
		countTask.Q = task.Q
		countTask.StartTime = task.StartTime
		countTask.EndTime = task.EndTime
		countTask.State = task.State
		countTask.Bind = task.Bind
		countTask.Info = task.Info

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

	if task.Appid != nil {
		sql.WriteString(" AND appid=?")
		args = append(args, task.Appid)
	}

	if task.Openid != nil {
		sql.WriteString(" AND openid=?")
		args = append(args, task.Openid)
	}

	if task.Unionid != nil {
		sql.WriteString(" AND unionid=?")
		args = append(args, task.Unionid)
	}

	if task.Uid != nil {
		sql.WriteString(" AND uid=?")
		args = append(args, task.Uid)
	}

	if task.Q != nil {
		sql.WriteString(" AND nick LIKE ?")
		args = append(args, fmt.Sprintf("%%%s%%", task.Q))
	}

	if task.State != nil {

		sql.WriteString(" AND state IN (")

		ids := strings.Split(dynamic.StringValue(task.State, ""), ",")
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

	if task.Bind != nil {
		if dynamic.BooleanValue(task.Bind, true) {
			sql.WriteString(" AND uid!=0")
		} else {
			sql.WriteString(" AND uid=0")
		}
	}

	if task.Info != nil {
		if dynamic.BooleanValue(task.Info, true) {
			sql.WriteString(" AND nick!=''")
		} else {
			sql.WriteString(" AND nick=''")
		}
	}

	if task.StartTime != nil {
		sql.WriteString(" AND mtime>=?")
		args = append(args, task.StartTime)
	}

	if task.EndTime != nil {
		sql.WriteString(" AND mtime<=?")
		args = append(args, task.EndTime)
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
