package feed

import (
	"fmt"

	"github.com/hailongz/golang/db"
	"github.com/hailongz/golang/dynamic"
	"github.com/hailongz/golang/json"
	"github.com/hailongz/golang/micro"
)

func (S *Service) Set(app micro.IContext, task *SetTask) (*Feed, error) {

	conn, prefix, err := app.GetDB("wd")

	if err != nil {
		return nil, err
	}

	prefix = Prefix(app, prefix, task.Id)

	v := Feed{}

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
		return nil, micro.NewError(ERROR_NOT_FOUND, "未找到动态")
	}

	keys := map[string]bool{}

	if task.Status != nil {
		v.Status = int32(dynamic.IntValue(task.Status, int64(v.Status)))
		keys["status"] = true
	}

	if task.Body != nil {
		v.Body = dynamic.StringValue(task.Body, "")
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

	if task.Ctime != nil {
		v.Ctime = dynamic.IntValue(task.Ctime, v.Ctime)
		keys["ctime"] = true
	}

	if len(keys) > 0 {

		_, err = db.UpdateWithKeys(conn, &v, prefix, keys)

		if err != nil {
			return nil, err
		}

		{

			// 缓存
			cli, prefix, err := app.GetRedis("default")

			if err != nil {
				return nil, err
			}
			_, _ = cli.Del(fmt.Sprintf("%s%d", prefix, v.Id)).Result()

		}

		app.SendMessage(task.GetName(), &v)
	}

	return &v, nil
}
