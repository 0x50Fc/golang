package comment

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

func (S *Service) UserCount(app micro.IContext, task *UserCountTask) (*CountData, error) {

	maxSecond := time.Duration(dynamic.IntValue(dynamic.GetWithKeys(app.GetConfig(), []string{"cache", "maxSecond"}), 60))

	cacheKey := cache.SignKey("un")

	{
		cache, err := app.GetCache("default")

		if err == nil {

			text, err := cache.GetItem(fmt.Sprintf("%d", task.Eid), cacheKey, maxSecond*time.Second)

			if err == nil && text != "" {
				c := &CountData{}
				err = json.Unmarshal([]byte(text), c)
				if err == nil {
					return c, nil
				}
			}
		}
	}

	conn, prefix, err := app.GetDB("rd")

	if err != nil {
		return nil, err
	}

	prefix = Prefix(app, prefix, task.Eid)

	sql := bytes.NewBuffer(nil)

	v := Comment{}

	args := []interface{}{}

	sql.WriteString(fmt.Sprintf("SELECT COUNT(DISTINCT  uid) as n FROM %s", db.TableName(prefix, &v)))
	sql.WriteString(" WHERE state=? AND eid=?")

	args = append(args, CommentState_None, task.Eid)

	result := &CountData{}

	err = func() error {

		rs, err := conn.Query(sql.String(), args...)

		if err != nil {
			return err
		}

		defer rs.Close()

		if rs.Next() {
			err = rs.Scan(&result.Total)
			if err != nil {
				return err
			}
		}

		return nil

	}()

	{
		cache, err := app.GetCache("default")

		if err == nil {

			b, _ := json.Marshal(result)

			cache.SetItem(fmt.Sprintf("%d", task.Eid), cacheKey, string(b), maxSecond*time.Second)

		}
	}

	return result, nil
}
