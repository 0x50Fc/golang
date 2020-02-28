package urelation

import (
	"fmt"
	"time"

	"github.com/hailongz/golang/db"
	"github.com/hailongz/golang/dynamic"
	"github.com/hailongz/golang/json"
	"github.com/hailongz/golang/micro"
)

func (S *Service) FollowGet(app micro.IContext, task *FollowGetTask) (*Follow, error) {

	v := Follow{}

	//先读取缓存缓存
	{
		redis, prefix, err := app.GetRedis("default")

		if err == nil {

			key := fmt.Sprintf("%s%d_%d", prefix, task.Uid, task.Fuid)

			text, err := redis.Get(key).Result()

			if err == nil && text != "" {
				if text == "false" {
					return nil, micro.NewError(ERROR_NOT_FOUND, "未找到好友")
				}
				err = json.Unmarshal([]byte(text), &v)
				if err == nil {
					return &v, nil
				}
			}

		}
	}

	conn, prefix, err := app.GetDB("rd")

	if err != nil {
		return nil, err
	}

	prefix = Prefix(app, prefix, task.Uid)

	rs, err := db.Query(conn, &v, prefix, " WHERE uid=? AND fuid=?", task.Uid, task.Fuid)

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

		//设置缓存
		{
			redis, prefix, err := app.GetRedis("default")

			if err == nil {

				key := fmt.Sprintf("%s%d_%d", prefix, task.Uid, task.Fuid)

				maxSecond := dynamic.IntValue(dynamic.GetWithKeys(app.GetConfig(), []string{"cache", "maxSecond"}), 60)

				redis.Set(key, "false", time.Second*time.Duration(maxSecond)).Result()

			}
		}

		return nil, micro.NewError(ERROR_NOT_FOUND, "未找到好友")
	}

	//设置缓存
	{
		redis, prefix, err := app.GetRedis("default")

		if err == nil {

			key := fmt.Sprintf("%s%d_%d", prefix, task.Uid, task.Fuid)

			text, _ := json.Marshal(&v)

			maxSecond := dynamic.IntValue(dynamic.GetWithKeys(app.GetConfig(), []string{"cache", "maxSecond"}), 60)

			redis.Set(key, string(text), time.Second*time.Duration(maxSecond)).Result()

		}
	}

	return &v, nil
}
