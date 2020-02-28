package top

import (
	"fmt"

	"github.com/hailongz/golang/db"
	"github.com/hailongz/golang/dynamic"
	"github.com/hailongz/golang/json"
	"github.com/hailongz/golang/micro"
)

func (S *Service) BatchAdd(app micro.IContext, task *BatchAddTask) ([]*Top, error) {

	conn, prefix, err := app.GetDB("wd")

	if err != nil {
		return nil, err
	}

	prefix = fmt.Sprintf("%s%s_", prefix, task.Name)

	items := []*Top{}

	v := Top{}

	err = db.Transaction(conn, func(conn db.Database) error {

		var err error = nil

		dynamic.Each(task.Items, func(_ interface{}, item interface{}) bool {

			tid := dynamic.IntValue(dynamic.Get(item, "tid"), 0)
			rate := int32(dynamic.IntValue(dynamic.Get(item, "rate"), 0))
			opt := dynamic.Get(item, "options")
			keyword := dynamic.Get(item, "keyword")

			if tid != 0 {

				err = func() error {

					rs, err := db.Query(conn, &v, prefix, " WHERE tid=?", tid)

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
						v.Tid = tid
					}

					return nil
				}()

				if err != nil {
					return false
				}

				v.Sid = NewSID(rate)

				keys := map[string]bool{"sid": true}

				if opt != nil {

					options := map[string]interface{}{}

					dynamic.Each(v.Options, func(key interface{}, value interface{}) bool {
						options[dynamic.StringValue(key, "")] = value
						return true
					})

					text := dynamic.StringValue(opt, "")

					var data interface{} = nil

					json.Unmarshal([]byte(text), &data)

					dynamic.Each(data, func(key interface{}, value interface{}) bool {
						options[dynamic.StringValue(key, "")] = value
						return true
					})

					v.Options = options
					keys["options"] = true
				}

				if keyword != nil {
					v.Keyword = dynamic.StringValue(keyword, v.Keyword)
					keys["keyword"] = true
				}

				if v.Id == 0 {
					_, err = db.Insert(conn, &v, prefix)

					if err != nil {
						return false
					}
				} else {

					_, err = db.UpdateWithKeys(conn, &v, prefix, keys)

					if err != nil {
						return false
					}
				}

				item := Top{}
				item = v
				items = append(items, &item)
			}

			return true
		})

		return err
	})

	if len(items) > 0 {

		{
			// 清除缓存

			redis, prefix, err := app.GetRedis("default")

			if err == nil {
				redis.Del(fmt.Sprintf("%s%s", prefix, task.Name)).Result()
				redis.Del(fmt.Sprintf("%s%s_rank", prefix, task.Name)).Result()
			}
		}

		// MQ 消息
		app.SendMessage(task.GetName(), items)
	}

	return items, nil
}
