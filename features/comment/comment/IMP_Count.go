package comment

import (
	"bytes"
	"fmt"
	"strings"
	"time"

	"github.com/hailongz/golang/cache"
	"github.com/hailongz/golang/db"
	"github.com/hailongz/golang/dynamic"
	"github.com/hailongz/golang/json"
	"github.com/hailongz/golang/micro"
)

func (S *Service) Count(app micro.IContext, task *CountTask) (*CountData, error) {

	maxSecond := time.Duration(dynamic.IntValue(dynamic.GetWithKeys(app.GetConfig(), []string{"cache", "maxSecond"}), 60))

	cacheKey := cache.SignKey("n", task.Id, task.Ids, task.Pid, task.Uid, task.Q, task.Path)

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

	args := []interface{}{}

	sql.WriteString(" WHERE state=? AND eid=?")

	args = append(args, CommentState_None, task.Eid)

	if task.Id != nil {
		sql.WriteString(" AND id=?")
		args = append(args, task.Id)
	}

	if task.Ids != nil {

		sql.WriteString(" AND id IN (")

		ids := strings.Split(dynamic.StringValue(task.Ids, ""), ",")
		i := 0

		for _, id := range ids {
			if id != "" {
				if i != 0 {
					sql.WriteString(",")
				}
				sql.WriteString("?")
				args = append(args, id)
				i++
			}
		}

		sql.WriteString(")")

	}

	if task.Pid != nil {
		sql.WriteString(" AND pid=?")
		args = append(args, task.Pid)
	}

	if task.Uid != nil {
		sql.WriteString(" AND uid=?")
		args = append(args, task.Uid)
	}

	if task.Q != nil {
		sql.WriteString(" AND body LIKE ?")
		args = append(args, fmt.Sprintf("%%%s%%", task.Q))
	}

	if task.Path != nil {
		sql.WriteString(" AND path LIKE ?")
		args = append(args, fmt.Sprintf("%s%%", task.Path))
	}

	//获取目标id的评论数
	comment := Comment{}
	count, err := db.Count(conn, &comment, prefix, sql.String(), args...)

	if err != nil {
		return nil, err
	}

	result := &CountData{}
	result.Total = int32(count)

	{
		cache, err := app.GetCache("default")

		if err == nil {

			b, _ := json.Marshal(result)

			cache.SetItem(fmt.Sprintf("%d", task.Eid), cacheKey, string(b), time.Second*maxSecond)

		}
	}

	return result, nil
}
