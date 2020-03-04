package job

import (
	"encoding/json"

	"github.com/hailongz/golang/db"
	"github.com/hailongz/golang/dynamic"
	"github.com/hailongz/golang/micro"
)

func (S *Service) JobSet(app micro.IContext, task *JobSetTask) (*Job, error) {

	conn, prefix, err := app.GetDB("wd")

	if err != nil {
		return nil, err
	}

	v := Job{}

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
		return nil, micro.NewError(ERROR_NOT_FOUND, "未找到分组")
	}

	keys := map[string]bool{}

	if task.Type != nil {
		v.Type = int32(dynamic.IntValue(task.Type, int64(v.Type)))
		keys["type"] = true
	}

	if task.Appid != nil {
		v.Appid = dynamic.IntValue(task.Appid, v.Appid)
		keys["appid"] = true
	}

	if task.Uid != nil {
		v.Uid = dynamic.IntValue(task.Uid, v.Uid)
		keys["uid"] = true
	}

	if task.MaxCount != nil {
		v.MaxCount = int32(dynamic.IntValue(task.MaxCount, int64(v.MaxCount)))
		keys["maxcount"] = true
	}

	if task.Count != nil {
		v.Count = int32(dynamic.IntValue(task.Count, int64(v.Count)))
		keys["count"] = true
	}

	if task.ErrCount != nil {
		v.ErrCount = int32(dynamic.IntValue(task.ErrCount, int64(v.ErrCount)))
		keys["errcount"] = true
	}

	if task.AddCount != nil {
		v.Count = v.Count + int32(dynamic.IntValue(task.AddCount, 0))
		keys["count"] = true
	}

	if task.AddErrCount != nil {
		v.ErrCount = v.ErrCount + int32(dynamic.IntValue(task.AddErrCount, 0))
		keys["errcount"] = true
	}

	if task.Alias != nil {
		v.Alias = dynamic.StringValue(task.Alias, v.Alias)
		keys["alias"] = true
	}

	if task.Stime != nil {
		v.Stime = dynamic.IntValue(task.Stime, v.Stime)
		keys["stime"] = true
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
