package notice

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/hailongz/golang/db"
	"github.com/hailongz/golang/dynamic"
	"github.com/hailongz/golang/micro"
)

func (S *Service) NewcountByTypes(app micro.IContext, task *NewcountByTypesTask) (*NewCountByTypesData, error) {

	conn, prefix, err := app.GetDB("rd")

	if err != nil {
		return nil, err
	}

	v := Notice{}

	prefix = Prefix(app, prefix, task.Uid)

	sql := bytes.NewBuffer(nil)

	args := []interface{}{}

	sql.WriteString(fmt.Sprintf("SELECT type,COUNT(*) as `count` FROM %s", db.TableName(prefix, &v)))
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

	if task.Q != nil {
		sql.WriteString(" AND body LIKE ?")
		args = append(args, fmt.Sprintf("%%%s%%", task.Q))
	}

	sql.WriteString(" AND id>?")
	args = append(args, task.TopId)

	sql.WriteString(" GROUP BY type")

	rs, err := conn.Query(sql.String(), args...)

	if err != nil {
		return nil, err
	}

	defer rs.Close()

	item := NewCountByTypesItem{}

	items := []*NewCountByTypesItem{}

	for rs.Next() {

		err := rs.Scan(&item.Type, &item.Count)

		if err != nil {
			return nil, err
		}

		i := NewCountByTypesItem{}
		i = item
		items = append(items, &i)
	}

	return &NewCountByTypesData{Items: items}, nil

}
