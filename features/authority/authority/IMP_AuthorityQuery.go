package authority

import (
	"bytes"
	"fmt"

	"github.com/hailongz/golang/db"
	"github.com/hailongz/golang/dynamic"
	"github.com/hailongz/golang/micro"
)

func (S *Service) AuthorityQuery(app micro.IContext, task *AuthorityQueryTask) (*AuthorityQueryData, error) {

	p := int32(dynamic.IntValue(task.P, 1))
	n := int32(dynamic.IntValue(task.N, 20))

	if n < 1 {
		n = 20
	}

	if p < 1 {
		p = 1
	}

	data := AuthorityQueryData{}

	if task.P != nil {

		countTask := AuthorityCountTask{}

		countTask.Uid = task.Uid
		countTask.RoleId = task.RoleId
		countTask.ResId = task.ResId

		countData, err := S.AuthorityCount(app, &countTask)

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

	if task.Uid != nil {
		sql.WriteString(" AND uid=?")
		args = append(args, task.Uid)
	}

	if task.RoleId != nil {
		sql.WriteString(" AND roleid=?")
		args = append(args, task.RoleId)
	}

	if task.ResId != nil {
		sql.WriteString(" AND resid=?")
		args = append(args, task.ResId)
	}

	sql.WriteString(fmt.Sprintf(" ORDER BY id DESC LIMIT %d,%d", (p-1)*n, n))

	v := Authority{}

	data.Items = []*Authority{}

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

		item := Authority{}
		item = v

		data.Items = append(data.Items, &item)
	}

	return &data, nil
}
