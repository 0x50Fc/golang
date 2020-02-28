package urelation

import (
	"fmt"

	"github.com/hailongz/golang/db"
	"github.com/hailongz/golang/micro"
)

func (S *Service) Rm(app micro.IContext, task *RmTask) (*Follow, error) {

	conn, prefix, err := app.GetDB("wd")

	if err != nil {
		return nil, err
	}

	if task.Uid == 0 {
		return nil, micro.NewError(ERROR_NOT_FOUND, "用户不存在")
	}

	if task.Fuid == 0 {
		return nil, micro.NewError(ERROR_NOT_FUID, "好友不存在")
	}

	if task.Uid == task.Fuid {
		return nil, micro.NewError(ERROR_SAME_UID_FUID, "自己不能取消关注自己")
	}

	p := Prefix(app, prefix, task.Uid)
	fp := Prefix(app, prefix, task.Fuid)

	follow := Follow{}

	err = db.Transaction(conn, func(conn db.Database) error {

		err := func() error {

			rs, err := db.Query(conn, &follow, p, " WHERE uid = ? AND fuid = ? ", task.Uid, task.Fuid)

			if err != nil {
				return err
			}

			defer rs.Close()

			if rs.Next() {
				scaner := db.NewScaner(&follow)
				err = scaner.Scan(rs)
				if err != nil {
					return err
				}
			} else {
				return micro.NewError(ERROR_URELATION_NOT_ISEXIT, "您没有关注该用户，不能取消关注")
			}

			return nil
		}()

		if err != nil {
			return err
		}

		//1.删除关注表
		_, err = db.Delete(conn, &follow, p)

		if err != nil {
			return err
		}

		//2.删除粉丝表

		fans := Fans{}

		_, err = conn.Exec(fmt.Sprintf("DELETE FROM %s WHERE uid=? AND fuid=?", db.TableName(fp, &fans)), task.Fuid, task.Uid)

		if err != nil {
			return err
		}

		//3.查询如果是互相关注了修改关注的type为0

		_, err = conn.Exec(fmt.Sprintf("UPDATE %s SET type=? WHERE uid=? AND fuid=?", db.TableName(fp, &follow)), RelationType_Weak, task.Fuid, task.Uid)

		if err != nil {
			return err
		}

		_, err = conn.Exec(fmt.Sprintf("UPDATE %s SET type=? WHERE uid=? AND fuid=?", db.TableName(p, &fans)), RelationType_Weak, task.Uid, task.Fuid)

		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	//清除缓存
	{
		redis, prefix, err := app.GetRedis("default")

		if err == nil {

			key := fmt.Sprintf("%s%d_%d", prefix, task.Uid, task.Fuid)
			_, _ = redis.Del(key).Result()

			key = fmt.Sprintf("%s%d_%d", prefix, task.Fuid, task.Uid)
			_, _ = redis.Del(key).Result()

			key = fmt.Sprintf("%s%d_n", prefix, task.Uid)
			_, _ = redis.Del(key).Result()

			key = fmt.Sprintf("%s%d_n", prefix, task.Fuid)
			_, _ = redis.Del(key).Result()
		}
	}

	//MQ 消息队列
	app.SendMessage(task.GetName(), &follow)

	return &follow, err
}
