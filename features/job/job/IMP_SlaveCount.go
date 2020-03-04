package job

import (
	"bytes"

	"github.com/hailongz/golang/db"
	"github.com/hailongz/golang/micro"
)

func (S *Service) SlaveCount(app micro.IContext, task *SlaveCountTask) (*CountData, error) {

	conn, prefix, err := app.GetDB("rd")

	if err != nil {
		return nil, err
	}

	sql := bytes.NewBuffer(nil)

	args := []interface{}{}

	sql.WriteString(" WHERE 1")

	v := Slave{}

	count, err := db.Count(conn, &v, prefix, sql.String(), args...)

	if err != nil {
		return nil, err
	}

	return &CountData{Total: int32(count)}, nil
}
