package settings

import (
	"log"

	"github.com/hailongz/golang/db"
	"github.com/hailongz/golang/dynamic"
	"github.com/hailongz/golang/json"
	"github.com/hailongz/golang/micro"
)

func (S *Service) Set(app micro.IContext, task *SetTask) (interface{}, error) {

	conn, prefix, err := app.GetDB("wd")

	if err != nil {
		return nil, err
	}

	v := Setting{}

	rs, err := db.Query(conn, &v, prefix, " WHERE name=?", task.Name)

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
	}

	v.Name = task.Name

	keys := map[string]bool{}

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

	if v.Id == 0 {

		_, err = db.Insert(conn, &v, prefix)

		if err != nil {
			return nil, err
		}

		{
			// 缓存
			cli, prefix, err := app.GetRedis("default")

			if err == nil {

				cli.Del(prefix + task.Name).Result()

			} else {
				log.Println("[Redis] [ERROR]", err)
			}

		}

		app.SendMessage(task.GetName(), &v)

	} else if len(keys) > 0 {

		_, err = db.UpdateWithKeys(conn, &v, prefix, keys)

		if err != nil {
			return nil, err
		}

		{
			// 缓存
			cli, prefix, err := app.GetRedis("default")

			if err == nil {

				cli.Del(prefix + task.Name).Result()

			} else {
				log.Println("[Redis] [ERROR]", err)
			}

		}

		app.SendMessage(task.GetName(), &v)
	}

	return &v, nil
}
