package authority

import (
	"bytes"
	"fmt"

	"github.com/hailongz/golang/db"
	"github.com/hailongz/golang/micro"
)

func (S *Service) RoleCount(app micro.IContext, task *RoleCountTask) (*RoleCountData, error) {

	conn, prefix, err := app.GetDB("rd")

	if err != nil {
		return nil, err
	}

	sql := bytes.NewBuffer(nil)

	args := []interface{}{}

	sql.WriteString(" WHERE 1")

	if task.Prefix != nil {
		sql.WriteString(" AND name LIKE ?")
		args = append(args, fmt.Sprintf("%s%%", task.Prefix))
	}

	v := Role{}

	count, err := db.Count(conn, &v, prefix, sql.String(), args...)

	if err != nil {
		return nil, err
	}

	return &RoleCountData{Total: int32(count)}, nil
}
