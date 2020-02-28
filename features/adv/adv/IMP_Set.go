package adv

import (
	"bytes"

	"github.com/hailongz/golang/db"
	"github.com/hailongz/golang/dynamic"
	"github.com/hailongz/golang/micro"
)

func (S *Service) Set(app micro.IContext, task *SetTask) (*Adv, error) {

	conn, prefix, err := app.GetDB("wd")

	if err != nil {
		return nil, err
	}

	//查询是否存在
	v := Adv{}

	sql := bytes.NewBuffer(nil)

	args := []interface{}{}

	sql.WriteString(" WHERE id = ? AND channel = ? AND position = ?")

	args = append(args, task.Id)
	args = append(args, task.Channel)
	args = append(args, task.Position)

	rs, err := db.Query(conn, &v, prefix, sql.String(), args...)

	if err != nil {
		return nil, err
	}

	defer rs.Close()

	if rs.Next() {
		scaner := db.NewScaner(&v)

		err = scaner.Scan(rs)

		if err != nil {
			return nil, err
		}
	} else {
		return nil, micro.NewError(ERROR_NOT_FOUND, "未找到广告")
	}

	//修改广告
	keys := map[string]bool{}
	if task.Title != "" {
		v.Title = dynamic.StringValue(task.Title, "")
		keys["title"] = true
	}

	if task.Channel != "" {
		v.Channel = dynamic.StringValue(task.Channel, "")
		keys["channel"] = true
	}

	if task.Position != 0 {
		v.Position = int32(dynamic.IntValue(task.Position, 0))
	}
	keys["position"] = true

	if task.Description != "" {
		v.Description = dynamic.StringValue(task.Description, "")
		keys["description"] = true
	}

	if task.Pic != "" {
		v.Pic = dynamic.StringValue(task.Pic, "")
		keys["pic"] = true
	}

	if task.Link != "" {
		v.Link = dynamic.StringValue(task.Link, "")
		keys["link"] = true
	}

	if task.Starttime != 0 {
		v.Starttime = dynamic.IntValue(task.Starttime, 0)
		keys["starttime"] = true
	}

	if task.Endtime != 0 {
		v.Endtime = dynamic.IntValue(task.Endtime, 0)
		keys["endtime"] = true
	}

	if task.Linktype != 0 {
		v.Linktype = int32(dynamic.IntValue(task.Linktype, 0))
		keys["linktype"] = true
	}

	if task.Sort != 0 {
		v.Sort = task.Sort
		keys["sort"] = true
	}

	_, err = db.UpdateWithKeys(conn, &v, prefix, keys)

	if err != nil {
		return nil, err
	}

	return &v, nil
}
