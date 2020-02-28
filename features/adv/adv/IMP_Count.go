package adv

import (
	"bytes"

	"github.com/hailongz/golang/db"
	"github.com/hailongz/golang/micro"
)

func (S *Service) Count(app micro.IContext, task *CountTask) (*CountData, error) {

	conn, prefix, err := app.GetDB("rd")

	if err != nil {
		return nil, err
	}

	sql := bytes.NewBuffer(nil)

	args := []interface{}{}

	sql.WriteString(" WHERE 1=1 ")

	if task.Id != nil {
		sql.WriteString(" AND id =?")
		args = append(args, task.Id)
	}

	if task.Channel != nil {
		sql.WriteString(" AND channel =?")
		args = append(args, task.Channel)
	}

	if task.Stime != nil {
		sql.WriteString(" AND starttime <=?")
		args = append(args, task.Stime)
	}

	if task.Etime != nil {
		sql.WriteString(" AND endtime >=?")
		args = append(args, task.Etime)
	}

	//获取广告数
	comment := Adv{}
	count, err := db.Count(conn, &comment, prefix, sql.String(), args...)

	if err != nil {
		return nil, err
	}

	result := &CountData{}
	result.Total = int32(count)

	return result, nil
}
