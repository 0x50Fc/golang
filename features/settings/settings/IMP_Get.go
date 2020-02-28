package settings

import (
	"log"
	"time"

	"github.com/hailongz/golang/db"
	"github.com/hailongz/golang/dynamic"
	"github.com/hailongz/golang/json"
	"github.com/hailongz/golang/micro"
)

func (S *Service) Get(app micro.IContext, task *GetTask) (interface{}, error) {

	{
		// 缓存
		cli, prefix, err := app.GetRedis("default")

		if err == nil {

			text, err := cli.Get(prefix + task.Name).Result()

			if err == nil && text != "" {

				var output interface{} = nil
				err = json.Unmarshal([]byte(text), &output)
				if err == nil {
					return output, nil
				} else {
					log.Println("[Redis] [json.Unmarshal] [ERROR]", err)
				}

			}

		} else {
			log.Println("[Redis] [ERROR]", err)
		}

	}

	conn, prefix, err := app.GetDB("rd")

	if err != nil {
		return nil, err
	}

	v := Setting{}

	err = func() error {

		p, err := db.Get(conn, &v, prefix, " WHERE `name`=?", task.Name)

		if err != nil {
			return err
		}

		if p == nil {
			return micro.NewError(ERROR_NOT_FOUND, "未找到配置")
		}

		return nil
	}()

	if err != nil {
		return nil, err
	}

	{
		// 缓存
		cli, prefix, err := app.GetRedis("default")

		if err == nil {

			maxSecond := dynamic.IntValue(dynamic.GetWithKeys(app.GetConfig(), []string{"cache", "maxSecond"}), 1800)

			b, _ := json.Marshal(v.Options)

			if b != nil {
				cli.Set(prefix+task.Name, string(b), time.Duration(maxSecond)*time.Second).Result()
			}

		}
	}

	return v.Options, nil
}
