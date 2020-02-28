package urelation

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/hailongz/golang/db"
	"github.com/hailongz/golang/dynamic"
	"github.com/hailongz/golang/micro"
)

func (S *Service) FansQuery(app micro.IContext, task *FansQueryTask) (*FansQueryData, error) {

	if task.Uid == 0 {
		return nil, micro.NewError(ERROR_NOT_FOUND, "用户不存在")
	}

	p := int32(dynamic.IntValue(task.P, 1))
	n := int32(dynamic.IntValue(task.N, 20))

	if n < 1 {
		n = 20
	}

	if p < 1 {
		p = 1
	}

	data := FansQueryData{}

	if task.P != nil {

		countTask := CountTask{}

		countTask.Uid = task.Uid

		countData, err := S.Count(app, &countTask)

		if err != nil {
			return nil, err
		}

		data.Page = &QueryDataPage{
			Total: countData.FollowCount,
			P:     int32(dynamic.IntValue(task.P, 0)),
			N:     n,
			Count: countData.FollowCount / n,
		}

		if countData.FollowCount%n != 0 {
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

	if task.In != nil {

		sql.WriteString(" AND id IN (")

		ids := strings.Split(dynamic.StringValue(task.In, ""), ",")
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

	sql.WriteString(fmt.Sprintf(" ORDER BY id ASC LIMIT %d,%d", (p-1)*n, n))

	v := Fans{}

	data.Items = []*Fans{}

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

		item := Fans{}
		item = v

		data.Items = append(data.Items, &item)
	}

	return &data, nil
}
