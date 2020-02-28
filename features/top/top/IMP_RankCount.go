package top

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/hailongz/golang/db"
	"github.com/hailongz/golang/dynamic"
	"github.com/hailongz/golang/micro"
)

func (S *Service) RankCount(app micro.IContext, task *RankCountTask) (*CountData, error) {

	conn, prefix, err := app.GetDB("rd")

	if err != nil {
		return nil, err
	}

	prefix = fmt.Sprintf("%s%s_", prefix, task.Name)

	sql := bytes.NewBuffer(nil)

	args := []interface{}{}

	sql.WriteString(" WHERE 1")

	if task.TopId != nil {
		sql.WriteString(" AND `rank` >= ?")
		args = append(args, task.TopId)
	}

	if task.Tids != nil {
		sql.WriteString(" AND tid IN(")
		for i, s := range strings.Split(dynamic.StringValue(task.Tids, ""), ",") {
			if i != 0 {
				sql.WriteString(",")
			}
			sql.WriteString("?")
			args = append(args, s)
		}
		sql.WriteString(")")
	}

	if task.Q != nil {
		sql.WriteString(" AND keyword LIKE ?")
		args = append(args, fmt.Sprintf("%%%s%%", task.Q))
	}

	v := Top{}

	count, err := db.Count(conn, &v, prefix, sql.String(), args...)

	if err != nil {
		return nil, err
	}

	return &CountData{Total: int32(count)}, nil
}
