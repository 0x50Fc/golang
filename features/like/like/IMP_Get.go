package like

import (
	"bytes"
	"fmt"
	"time"

	"github.com/hailongz/golang/cache"
	"github.com/hailongz/golang/db"
	"github.com/hailongz/golang/dynamic"
	"github.com/hailongz/golang/json"
	"github.com/hailongz/golang/micro"
)

func (S *Service) Get(app micro.IContext, task *GetTask) (*Like, error) {

	like := Like{}

	iid := dynamic.IntValue(task.Iid, 0)

	maxSecond := time.Duration(dynamic.IntValue(dynamic.GetWithKeys(app.GetConfig(), []string{"cache", "maxSecond"}), 60))

	cacheKey := cache.SignKey("i", iid, task.Uid)

	//从缓存中读取
	{
		cache, err := app.GetCache("default")
		if err == nil {
			text, err := cache.GetItem(fmt.Sprintf("%d", task.Tid), cacheKey, maxSecond*time.Second)
			if err == nil {
				if text == "false" {
					return nil, micro.NewError(ERROR_NOT_FOUND, "未找到赞")
				}
				err = json.Unmarshal([]byte(text), &like)
				if err == nil {
					return &like, nil
				}
			}
		}
	}

	conn, prefix, err := app.GetDB("rd")

	if err != nil {
		return nil, err
	}

	prefix = Prefix(app, prefix, task.Tid)

	sql := bytes.NewBuffer(nil)

	args := []interface{}{}

	sql.WriteString(" WHERE tid = ? and uid = ? and iid=?")

	args = append(args, task.Tid, task.Uid, iid)

	isReady := false

	err = func() error {

		rs, err := db.Query(conn, &like, prefix, sql.String(), args...)

		if err != nil {
			return err
		}

		defer rs.Close()

		if rs.Next() {

			scaner := db.NewScaner(&like)

			err = scaner.Scan(rs)

			if err != nil {
				return err
			}
			isReady = true
		} else {
			isReady = true
			return micro.NewError(ERROR_NOT_FOUND, "未找到赞")
		}

		return nil
	}()

	if err != nil {
		if isReady {

			cache, err := app.GetCache("default")

			if err == nil {
				cache.SetItem(fmt.Sprintf("%d", task.Tid), cacheKey, "false", maxSecond*time.Second)
			}

		}
		return nil, err
	}

	{

		cache, err := app.GetCache("default")

		if err == nil {
			maxSecond := dynamic.IntValue(dynamic.GetWithKeys(app.GetConfig(), []string{"cache", "maxSecond"}), 60)
			b, _ := json.Marshal(&like)
			cache.SetItem(fmt.Sprintf("%d", task.Tid), cacheKey, string(b), time.Duration(maxSecond)*time.Second)
		}

	}

	return &like, nil
}
