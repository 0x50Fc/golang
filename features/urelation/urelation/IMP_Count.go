package urelation

import (
	"fmt"
	"time"

	"github.com/hailongz/golang/db"
	"github.com/hailongz/golang/dynamic"
	"github.com/hailongz/golang/json"
	"github.com/hailongz/golang/micro"
)

func (S *Service) Count(app micro.IContext, task *CountTask) (*CountData, error) {

	//先读取缓存缓存
	{
		redis, prefix, err := app.GetRedis("default")

		if err == nil {

			key := fmt.Sprintf("%s%d_n", prefix, task.Uid)

			text, err := redis.Get(key).Result()

			if err == nil && text != "" {
				output := &CountData{}
				err = json.Unmarshal([]byte(text), output)
				if err == nil {
					return output, nil
				}
			}

		}
	}

	//缓存没有查询数据库
	conn, prefix, err := app.GetDB("rd")

	if err != nil {
		return nil, err
	}

	prefix = Prefix(app, prefix, task.Uid)

	//获取用户关注数
	follow := Follow{}

	rsFollow, err := db.Count(conn, &follow, prefix, " WHERE uid = ? ", task.Uid)

	//获取用户粉丝数
	fans := Fans{}

	rsFans, err := db.Count(conn, &fans, prefix, " WHERE uid = ? ", task.Uid)

	result := &CountData{}
	result.FollowCount = int32(rsFollow)
	result.FansCount = int32(rsFans)

	//设置缓存
	{
		redis, prefix, err := app.GetRedis("default")

		if err == nil {

			key := fmt.Sprintf("%s%d_n", prefix, task.Uid)

			text, _ := json.Marshal(result)

			maxSecond := dynamic.IntValue(dynamic.GetWithKeys(app.GetConfig(), []string{"cache", "maxSecond"}), 60)

			redis.Set(key, string(text), time.Second*time.Duration(maxSecond)).Result()

		}
	}

	return result, nil
}
