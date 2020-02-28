package notice

import (
	"github.com/hailongz/golang/db"
	"github.com/hailongz/golang/dynamic"
	"github.com/hailongz/golang/json"
	"github.com/hailongz/golang/micro"
)

func (S *Service) Set(app micro.IContext, task *SetTask) (*Notice, error) {

	conn, prefix, err := app.GetDB("wd")

	if err != nil {
		return nil, err
	}

	prefix = Prefix(app, prefix, task.Uid)

	v := Notice{}

	rs, err := db.Query(conn, &v, prefix, " WHERE id=?", task.Id)

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
		return nil, micro.NewError(ERROR_NOT_FOUND, "未找到通知")
	}

	keys := map[string]bool{}

	if task.Type != nil {
		v.Type = int32(dynamic.IntValue(task.Type, int64(v.Type)))
		keys["type"] = true
	}

	if task.Fid != nil {
		v.Fid = dynamic.IntValue(task.Fid, v.Fid)
		keys["fid"] = true
	}

	if task.Iid != nil {
		v.Iid = dynamic.IntValue(task.Iid, v.Iid)
		keys["iid"] = true
	}

	if task.Body != nil {
		v.Body = dynamic.StringValue(task.Body, v.Body)
		keys["body"] = true
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

		_, err = db.UpdateWithKeys(conn, &v, prefix, keys)

		if err != nil {
			return nil, err
		}

		app.SendMessage(task.GetName(), &v)
	}

	return &v, nil
}
