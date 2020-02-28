package urelation

import (
	"fmt"
	"time"

	"github.com/hailongz/golang/db"
	"github.com/hailongz/golang/micro"
)

func (S *Service) Follow(app micro.IContext, task *FollowTask) (*Follow, error) {

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
		return nil, micro.NewError(ERROR_SAME_UID_FUID, "自己不能关注自己")
	}

	p := Prefix(app, prefix, task.Uid)

	fp := Prefix(app, prefix, task.Fuid)

	follow := Follow{}
	fans := Fans{}

	err = db.Transaction(conn, func(conn db.Database) error {

		//先查询是否存在

		err := func() error {

			rs, err := db.Query(conn, &follow, p, " WHERE uid = ? AND fuid = ? ", task.Uid, task.Fuid)

			if err != nil {
				return err
			}

			defer rs.Close()

			if rs.Next() {
				return micro.NewError(ERROR_URELATION_ISEXIT, "已经关注过该用户了")
			}

			return nil
		}()

		if err != nil {
			return err
		}

		//插入关注表
		follow.Uid = task.Uid
		follow.Fuid = task.Fuid
		follow.Type = RelationType_Weak
		follow.Title = task.Title
		follow.Ctime = time.Now().Unix()

		err = func() error {

			v := Follow{}
			rs, err := db.Query(conn, &v, p, " WHERE uid = ? AND fuid = ? ", task.Fuid, task.Uid)

			if err != nil {
				return err
			}

			defer rs.Close()

			if rs.Next() {
				follow.Type = RelationType_Strong
			}

			return nil
		}()

		if err != nil {
			return err
		}

		_, err = db.Insert(conn, &follow, p)

		if err != nil {
			return err
		}

		fans.Uid = task.Fuid
		fans.Fuid = task.Uid
		fans.Type = follow.Type
		fans.Ctime = time.Now().Unix()

		_, err = db.Insert(conn, &fans, fp)

		if err != nil {
			return err
		}

		if follow.Type == RelationType_Strong {
			_, err = conn.Exec(fmt.Sprintf("UPDATE %s SET type=? WHERE uid=? AND fuid=?", db.TableName(fp, &follow)), RelationType_Strong, task.Fuid, task.Uid)
			if err != nil {
				return err
			}
			_, err = conn.Exec(fmt.Sprintf("UPDATE %s SET type=? WHERE uid=? AND fuid=?", db.TableName(p, &fans)), RelationType_Strong, task.Uid, task.Fuid)
			if err != nil {
				return err
			}
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

	return &follow, nil
}
