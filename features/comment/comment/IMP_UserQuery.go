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

func (S *Service) UserQuery(app micro.IContext, task *UserQueryTask) (*UserQueryData, error) {

	p := int32(dynamic.IntValue(task.P, 1))
	n := int32(dynamic.IntValue(task.N, 20))

	if n < 1 {
		n = 20
	}

	if p < 1 {
		p = 1
	}

	data := UserQueryData{}

	maxSecond := time.Duration(dynamic.IntValue(dynamic.GetWithKeys(app.GetConfig(), []string{"cache", "maxSecond"}), 60))

	cacheKey := cache.SignKey("uq", task.P, task.Maxtime, task.Mintime)

	if p == 1 {

		cache, err := app.GetCache("default")

		if err == nil {

			text, err := cache.GetItem(fmt.Sprintf("%d", task.Eid), cacheKey, maxSecond*time.Second)

			if err == nil && text != "" {
				err = json.Unmarshal([]byte(text), &data)
				if err == nil {
					return &data, nil
				}
			}
		}
	}

	if task.P != nil {

		countTask := UserCountTask{}

		countTask.Eid = task.Eid

		countData, err := S.UserCount(app, &countTask)

		if err != nil {
			return nil, err
		}

		data.Page = &Page{
			Total: countData.Total,
			P:     p,
			N:     n,
			Count: countData.Total / n,
		}

		if countData.Total%n != 0 {
			data.Page.Count++
		}
	}

	conn, prefix, err := app.GetDB("rd")

	if err != nil {
		return nil, err
	}

	prefix = Prefix(app, prefix, task.Eid)

	sql := bytes.NewBuffer(nil)

	args := []interface{}{}

	v := Comment{}

	sql.WriteString(fmt.Sprintf("SELECT uid, COUNT(*) as `count`, MAX(ctime) as ctime FROM %s", db.TableName(prefix, &v)))

	sql.WriteString(" WHERE state=? AND eid=? ")

	args = append(args, CommentState_None, task.Eid)

	if task.Maxtime != nil {
		sql.WriteString(" AND ctime <=?")
		args = append(args, task.Maxtime)
	}

	if task.Mintime != nil {
		sql.WriteString(" AND ctime >=?")
		args = append(args, task.Mintime)
	}

	sql.WriteString("  GROUP BY uid")

	sql.WriteString(fmt.Sprintf(" ORDER BY id DESC LIMIT %d,%d", (p-1)*n, n))

	data.Items = []*User{}

	rs, err := conn.Query(sql.String(), args...)

	if err != nil {
		return nil, err
	}

	defer rs.Close()

	for rs.Next() {

		item := User{}

		err = rs.Scan(&item.Uid, &item.Count, &item.Ctime)

		if err != nil {
			return nil, err
		}

		data.Items = append(data.Items, &item)
	}

	if p == 1 {

		cache, err := app.GetCache("default")

		if err == nil {

			b, _ := json.Marshal(&data)

			cache.SetItem(fmt.Sprintf("%d", task.Eid), cacheKey, string(b), time.Second*maxSecond)

		}
	}

	return &data, nil
}
