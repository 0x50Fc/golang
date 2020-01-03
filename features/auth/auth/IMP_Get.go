package auth

import (
	"github.com/hailongz/golang/db"
	"github.com/hailongz/golang/dynamic"
	"github.com/hailongz/golang/json"
	"github.com/hailongz/golang/micro"
)

func (S *Service) Get(app micro.IContext, task *GetTask) (interface{}, error) {

	key := task.Key
	stype := dynamic.StringValue(task.Type, AuthType_JSON)

	{
		// 缓存
		cli, prefix, err := app.GetRedis("default")

		if err == nil {

			text, err := cli.Get(prefix + key).Result()

			if err == nil && text != "" {

				if stype == AuthType_JSON {
					var output interface{} = nil
					err = json.Unmarshal([]byte(text), &output)
					if err != nil {
						return nil, err
					}
					return output, nil
				}

				return text, nil
			}

		} else {
			app.Println("[Redis] [ERROR]", err)
		}

	}

	conn, prefix, err := app.GetDB("rd")

	if err != nil {
		return nil, err
	}

	v := Auth{}

	err = func() error {

		rs, err := db.Query(conn, &v, prefix, " WHERE `key`=?", task.Key)

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

			return nil
		}

		return micro.NewError(ERROR_NOT_FOUND, "未找到验证对象")

	}()

	if err != nil {
		return nil, err
	}

	if stype == AuthType_JSON {
		var output interface{} = nil
		err = json.Unmarshal([]byte(v.Value), &output)
		if err != nil {
			return nil, err
		}
		return output, nil
	}

	return v.Value, nil
}
