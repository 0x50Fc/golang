package wx

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/hailongz/golang/db"
	"github.com/hailongz/golang/dynamic"
	"github.com/hailongz/golang/micro"
)

func (S *Service) Count(app micro.IContext, task *CountTask) (*CountData, error) {

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

	v := User{}

	count, err := db.Count(conn, &v, prefix, sql.String(), args...)

	if err != nil {
		return nil, err
	}

	return &CountData{Total: int32(count)}, nil
}
