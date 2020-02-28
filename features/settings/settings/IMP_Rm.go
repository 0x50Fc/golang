package settings

import (
	"log"

	"github.com/hailongz/golang/db"
	"github.com/hailongz/golang/micro"
)

func (S *Service) Rm(app micro.IContext, task *RmTask) (interface{}, error) {

	conn, prefix, err := app.GetDB("wd")

	if err != nil {
		return nil, err
	}

	v := Setting{}

	rs, err := db.Query(conn, &v, prefix, " WHERE name=?", task.Name)

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
		return nil, micro.NewError(ERROR_NOT_FOUND, "未找到配置")
	}

	_, err = db.Delete(conn, &v, prefix)

	if err != nil {
		return nil, err
	}

	{
		// 缓存
		cli, prefix, err := app.GetRedis("default")

		if err == nil {

			cli.Del(prefix + task.Name).Result()

		} else {
			log.Println("[Redis] [ERROR]", err)
		}

	}

	app.SendMessage(task.GetName(), &v)

	return v.Options, nil
}
