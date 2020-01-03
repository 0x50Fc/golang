package auth

import (
	"time"

	"github.com/hailongz/golang/db"
	"github.com/hailongz/golang/dynamic"
	"github.com/hailongz/golang/json"
	"github.com/hailongz/golang/micro"
)

func (S *Service) Set(app micro.IContext, task *SetTask) (interface{}, error) {

	maxSecond := dynamic.IntValue(dynamic.GetWithKeys(app.GetConfig(), []string{"cache", "maxSecond"}), 1800)
	stype := dynamic.StringValue(task.Type, AuthType_JSON)
	key := task.Key

	var output interface{} = nil

	if stype == AuthType_JSON && task.Value != nil {
		err := json.Unmarshal([]byte(dynamic.StringValue(task.Value, "")), &output)
		if err != nil {
			return nil, err
		}
	} else {
		output = task.Value
	}

	conn, prefix, err := app.GetDB("wd")

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

	keys := map[string]bool{}

	if task.Value != nil {
		keys["value"] = true
		if v.Value != "" && stype == AuthType_JSON {
			output = Subjoin(v.Value, output)
			b, _ := json.Marshal(output)
			v.Value = string(b)
		} else {
			v.Value = dynamic.StringValue(task.Value, "")
		}
	}

	if task.Expires != nil {
		keys["etime"] = true
		v.Etime = time.Now().Unix() + dynamic.IntValue(task.Expires, 0)
	}

	if len(keys) == 0 {
		return output, nil
	}

	_, err = db.UpdateWithKeys(conn, &v, prefix, keys)

	if err != nil {
		return nil, err
	}

	{
		tv := v.Etime - time.Now().Unix()

		// 缓存
		cli, prefix, err := app.GetRedis("default")

		if err == nil {
			if tv > 0 {

				if tv > maxSecond {
					tv = maxSecond
				}

				_, err = cli.Set(prefix+key, v.Value, time.Second*time.Duration(tv)).Result()

				if err != nil {
					return nil, err
				}
			}
		} else {
			app.Println("[Redis] [ERROR]", err)
		}
	}

	// MQ 消息
	app.SendMessage(task.GetName(), &v)

	return output, nil
}
