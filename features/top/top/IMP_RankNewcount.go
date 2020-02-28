package top

import (
	"bytes"
	"fmt"

	"github.com/hailongz/golang/db"
	"github.com/hailongz/golang/micro"
)

func (S *Service) RankNewcount(app micro.IContext, task *RankNewcountTask) (*CountData, error) {

	conn, prefix, err := app.GetDB("rd")

	if err != nil {
		return nil, err
	}

	prefix = fmt.Sprintf("%s%s_", prefix, task.Name)

	sql := bytes.NewBuffer(nil)

	args := []interface{}{}

	sql.WriteString(" WHERE 1")

	sql.WriteString(" AND `rank` < ?")
	args = append(args, task.TopId)

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
