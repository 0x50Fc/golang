package top

import (
	"fmt"

	"github.com/hailongz/golang/db"
	"github.com/hailongz/golang/micro"
)

func (S *Service) Clean(app micro.IContext, task *CleanTask) (interface{}, error) {

	conn, prefix, err := app.GetDB("wd")

	if err != nil {
		return nil, err
	}

	prefix = fmt.Sprintf("%s%s_", prefix, task.Name)

	v := Top{}

	if task.Limit == nil {

		_, err = conn.Exec(fmt.Sprintf("DELETE FROM %s", db.TableName(prefix, &v)))

		if err != nil {
			return nil, err
		}

	} else {

		rs, err := db.Query(conn, &v, prefix, " ORDER BY sid DESC LIMIT ?,1", task.Limit)

		if err != nil {
			return nil, err
		}

		if rs.Next() {

			scaner := db.NewScaner(&v)

			err = scaner.Scan(rs)

			rs.Close()

			if err != nil {
				return nil, err
			}

		} else {
			rs.Close()
		}

		_, err = conn.Exec(fmt.Sprintf("DELETE FROM %s WHERE sid <=?", db.TableName(prefix, &v)), v.Sid)

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

	return nil, nil
}
