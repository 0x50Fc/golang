package media

import (
	"bytes"
	"fmt"

	"github.com/hailongz/golang/db"
	"github.com/hailongz/golang/dynamic"
	"github.com/hailongz/golang/json"
	"github.com/hailongz/golang/micro"
)

func (S *Service) Set(app micro.IContext, task *SetTask) (*Media, error) {

	conn, name, err := app.GetDB("wd")

	if err != nil {
		return nil, err
	}

	if task.Region == nil {
		name = fmt.Sprintf("%s%s", name, dynamic.StringValue(task.Name, ""))
	} else {
		name = fmt.Sprintf("%s%d_%s", name, dynamic.IntValue(task.Region, 0), dynamic.StringValue(task.Name, ""))
	}

	app.Println("[NAME]", name)

	v := Media{}

	sql := bytes.NewBuffer(nil)

	args := []interface{}{}

	sql.WriteString(" WHERE id=?")

	args = append(args, task.Id)

	if task.Uid != nil {
		sql.WriteString(" AND uid=?")
		args = append(args, task.Uid)
	}

	p, err := db.Get(conn, &v, name, sql.String(), args...)

	if err != nil {
		return nil, err
	}

	if p == nil {
		return nil, micro.NewError(ERROR_NOT_FOUND, "未找到媒体对象")
	}

	keys := map[string]bool{}

	if task.Uid != nil {
		v.Uid = dynamic.IntValue(task.Uid, v.Uid)
		keys["uid"] = true
	}

	if task.Type != nil {
		v.Type = dynamic.StringValue(task.Type, v.Type)
		keys["type"] = true
	}

	if task.Title != nil {
		v.Title = dynamic.StringValue(task.Title, v.Title)
		keys["title"] = true
	}

	if task.Keyword != nil {
		v.Keyword = dynamic.StringValue(task.Keyword, v.Keyword)
		keys["keyword"] = true
	}

	if task.Path != nil {
		v.Path = dynamic.StringValue(task.Path, v.Path)
		keys["path"] = true
	}

	if task.Options != nil {

		options := map[string]interface{}{}

		dynamic.Each(v.Options, func(key interface{}, value interface{}) bool {
			options[dynamic.StringValue(key, "")] = value
			return true
		})

		text := dynamic.StringValue(task.Options, "")

		var data interface{} = nil

		json.Unmarshal([]byte(text), &data)

		dynamic.Each(data, func(key interface{}, value interface{}) bool {
			options[dynamic.StringValue(key, "")] = value
			return true
		})

		v.Options = options
		keys["options"] = true
	}

	if len(keys) > 0 {

		_, err = db.UpdateWithKeys(conn, &v, name, keys)

		if err != nil {
			return nil, err
		}

	}

	app.SendMessage(task.GetName(), &v)

	return &v, nil
}
