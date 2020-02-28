package addup

import (
	"bytes"
	"fmt"

	"github.com/hailongz/golang/dynamic"
	"github.com/hailongz/golang/json"
	"github.com/hailongz/golang/micro"
)

func (S *Service) Rm(app micro.IContext, task *RmTask) (interface{}, error) {

	conn, name, err := app.GetDB("wd")

	if err != nil {
		return nil, err
	}

	if task.Region == nil {
		name = fmt.Sprintf("%s%s", name, task.Name)
	} else {
		name = fmt.Sprintf("%s%d_%s", name, dynamic.IntValue(task.Region, 0), task.Name)
	}

	app.Println("[NAME]", name)

	sql := bytes.NewBuffer(nil)
	args := []interface{}{}

	sql.WriteString("DELETE FROM ")
	sql.WriteString(name)
	if task.Where != nil {
		sql.WriteString(" WHERE ")
		sql.WriteString(dynamic.StringValue(task.Where, ""))
	}
	if task.OrderBy != nil {
		sql.WriteString(" ORDER BY ")
		sql.WriteString(dynamic.StringValue(task.OrderBy, ""))
	}

	if task.Limit != nil {
		sql.WriteString(" LIMIT ")
		sql.WriteString(dynamic.StringValue(task.Limit, ""))
	}

	if task.Args != nil {

		var data interface{} = nil

		err = json.Unmarshal([]byte(dynamic.StringValue(task.Args, "")), &data)

		if err != nil {
			return nil, err
		}

		dynamic.Each(data, func(_, value interface{}) bool {
			args = append(args, value)
			return true
		})

	}

	app.Println("[SQL]", sql.String())

	_, err = conn.Exec(sql.String(), args...)

	if err != nil {
		return nil, err
	}

	{

		cache, err := app.GetCache("default")
		if err == nil {
			cache.Del(name)
		}
	}

	return nil, nil
}
