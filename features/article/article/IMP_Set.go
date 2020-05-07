package article

import (
	"fmt"
	"time"

	"github.com/hailongz/golang/db"
	"github.com/hailongz/golang/dynamic"
	"github.com/hailongz/golang/json"
	"github.com/hailongz/golang/micro"
)

func (S *Service) Set(app micro.IContext, task *SetTask) (*Article, error) {

	conn, prefix, err := app.GetDB("wd")

	if err != nil {
		return nil, err
	}

	prefix = Prefix(app, prefix, task.Id)

	v := Article{}

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

	if task.Body != nil {
		v.Body = dynamic.StringValue(task.Body, "")
		keys["body"] = true
	}

	var options interface{} = nil

	if task.Options != nil {

		text := dynamic.StringValue(task.Options, "")
		var data interface{} = nil
		json.Unmarshal([]byte(text), &data)

		options = db.Merge(v.Options, data)

		v.Options = data
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

		if keys["options"] {
			v.Options = options
		}

		{
			maxSecond := time.Duration(dynamic.IntValue(dynamic.GetWithKeys(app.GetConfig(), []string{"cache", "maxSecond"}), 60))

			cache, _ := app.GetCache("default")

			if cache != nil {
				b, _ := json.Marshal(&v)
				cache.Set(fmt.Sprintf("%d", v.Id), string(b), maxSecond*time.Second)
			}
		}

		app.SendMessage(task.GetName(), &v)
	}

	return &v, nil
}
