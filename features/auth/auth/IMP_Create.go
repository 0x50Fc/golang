package auth

import (
	"time"

	"github.com/hailongz/golang/db"
	"github.com/hailongz/golang/dynamic"
	"github.com/hailongz/golang/json"
	"github.com/hailongz/golang/micro"
)

/**
 * JSON 数据叠加
 */
func Subjoin(v string, output interface{}) interface{} {
	var data interface{} = nil

	_ = json.Unmarshal([]byte(v), &data)

	if data == nil {
		data = map[string]interface{}{}
	}

	dynamic.Each(output, func(key interface{}, value interface{}) bool {
		dynamic.Set(data, dynamic.StringValue(key, ""), value)
		return true
	})

	return data
}

func (S *Service) Create(app micro.IContext, task *CreateTask) (interface{}, error) {

	maxSecond := dynamic.IntValue(dynamic.GetWithKeys(app.GetConfig(), []string{"cache", "maxSecond"}), 1800)
	stype := dynamic.StringValue(task.Type, AuthType_JSON)
	expires := int64(task.Expires)
	key := task.Key

	if expires == 0 {
		expires = 60
	}

	var output interface{} = nil

	if stype == AuthType_JSON {
		err := json.Unmarshal([]byte(task.Value), &output)
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

		rs, err := db.Query(conn, &v, prefix, " WHERE `key`=?", key)

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

		}

		return nil
	}()

	if err != nil {
		return nil, err
	}

	v.Key = task.Key

	if v.Value != "" && stype == AuthType_JSON {
		output = Subjoin(v.Value, output)
		b, _ := json.Marshal(output)
		v.Value = string(b)
	} else {
		v.Value = task.Value
	}

	v.Etime = time.Now().Unix() + expires

	if v.Id == 0 {
		_, err = db.Insert(conn, &v, prefix)
		if err != nil {
			return nil, err
		}
	} else {
		_, err = db.UpdateWithKeys(conn, &v, prefix, map[string]bool{"value": true, "etime": true})
		if err != nil {
			return nil, err
		}
	}

	{
		// 缓存
		cli, prefix, err := app.GetRedis("default")

		if err == nil {
			tv := expires

			if tv > maxSecond {
				tv = maxSecond
			}

			_, _ = cli.Set(prefix+key, v.Value, time.Second*time.Duration(tv)).Result()

		} else {
			app.Println("[Redis] [ERROR]", err)
		}

	}

	// MQ 消息
	app.SendMessage(task.GetName(), &v)

	return output, nil
}
