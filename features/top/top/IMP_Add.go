package top

import (
	"fmt"
	"time"

	"github.com/hailongz/golang/db"
	"github.com/hailongz/golang/dynamic"
	"github.com/hailongz/golang/json"
	"github.com/hailongz/golang/micro"
)

func (S *Service) Add(app micro.IContext, task *AddTask) (*Top, error) {

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
			v.Tid = task.Tid
		}

		return nil
	}()

	if err != nil {
		return nil, err
	}

	if task.Time == nil {
		v.Sid = NewSID(task.Rate)
	} else {
		v.Sid = NewSIDWithTimestamp(task.Rate, dynamic.IntValue(task.Time, time.Now().UnixNano()/int64(time.Millisecond)))
	}

	keys := map[string]bool{"sid": true}

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

	if v.Id == 0 {
		_, err = db.Insert(conn, &v, prefix)

		if err != nil {
			return nil, err
		}
	} else {

		_, err = db.UpdateWithKeys(conn, &v, prefix, keys)

		if err != nil {
			return nil, err
		}
	}

	{
		// 清除缓存

		redis, prefix, err := app.GetRedis("default")

		if err == nil {
			redis.Del(fmt.Sprintf("%s%s", prefix, task.Name)).Result()
			redis.Del(fmt.Sprintf("%s%s_rank", prefix, task.Name)).Result()
		}
	}

	// MQ 消息
	app.SendMessage(task.GetName(), &v)

	return &v, nil
}
