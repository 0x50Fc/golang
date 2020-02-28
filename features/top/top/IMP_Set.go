package top

import (
	"fmt"

	"github.com/hailongz/golang/db"
	"github.com/hailongz/golang/dynamic"
	"github.com/hailongz/golang/json"
	"github.com/hailongz/golang/micro"
)

func (S *Service) Set(app micro.IContext, task *SetTask) (*Top, error) {

	conn, prefix, err := app.GetDB("wd")

	if err != nil {
		return nil, err
	}

	prefix = fmt.Sprintf("%s%s_", prefix, task.Name)

	v := Top{}

	err = func() error {

		rs, err := db.Query(conn, &v, prefix, " WHERE tid=?", task.Tid)

		if err != nil {
			return err
		}

		defer rs.Close()

		if rs.Next() {

			scaner := db.NewScaner(&v)

			err = scaner.Scan(rs)

			if err != nil {
				return err
			}
		} else {
			return micro.NewError(ERROR_NOT_FOUND, "未找到推荐项")
		}

		return nil
	}()

	if err != nil {
		return nil, err
	}

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

	if task.Keyword != nil {
		v.Keyword = dynamic.StringValue(task.Keyword, v.Keyword)
		keys["keyword"] = true
	}

	if len(keys) > 0 {

		_, err = db.UpdateWithKeys(conn, &v, prefix, keys)

		if err != nil {
			return nil, err
		}

		{
			// 清除缓存

			redis, prefix, err := app.GetRedis("default")

			if err == nil {
				redis.Del(fmt.Sprintf("%s%s", prefix, task.Name)).Result()
				redis.Del(fmt.Sprintf("%s%s_rank", prefix, task.Name)).Result()
			}
		}

		app.SendMessage(task.GetName(), &v)
	}

	return &v, nil
}
