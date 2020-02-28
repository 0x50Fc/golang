package feed

import (
	"fmt"
	"time"

	"github.com/hailongz/golang/db"
	"github.com/hailongz/golang/dynamic"
	"github.com/hailongz/golang/json"
	"github.com/hailongz/golang/micro"
)

func (S *Service) Get(app micro.IContext, task *GetTask) (*Feed, error) {

	{

		// 缓存
		cli, prefix, err := app.GetRedis("default")

		if err == nil {

			text, err := cli.Get(fmt.Sprintf("%s%d", prefix, task.Id)).Result()

			if err == nil && text != "" {

				v := Feed{}

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

	prefix = Prefix(app, prefix, task.Id)

	v := Feed{}

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

	{

		maxSecond := dynamic.IntValue(dynamic.GetWithKeys(app.GetConfig(), []string{"cache", "maxSecond"}), 60)

		// 缓存
		cli, prefix, err := app.GetRedis("default")

		if err == nil {

			b, _ := json.Marshal(&v)

			_, _ = cli.Set(fmt.Sprintf("%s%d", prefix, v.Id), string(b), time.Duration(maxSecond)*time.Second).Result()

		}

	}

	return &v, nil
}
