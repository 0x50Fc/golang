package authority

import (
	"bytes"

	"github.com/hailongz/golang/db"
	"github.com/hailongz/golang/micro"
)

func (S *Service) AuthorityCount(app micro.IContext, task *AuthorityCountTask) (*AuthorityCountData, error) {

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

	v := Authority{}

	count, err := db.Count(conn, &v, prefix, sql.String(), args...)

	if err != nil {
		return nil, err
	}

	return &AuthorityCountData{Total: int32(count)}, nil
}
