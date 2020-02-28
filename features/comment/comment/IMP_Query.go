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

func (S *Service) Query(app micro.IContext, task *QueryTask) (*QueryData, error) {

	p := int32(dynamic.IntValue(task.P, 1))
	n := int32(dynamic.IntValue(task.N, 20))

	if n < 1 {
		n = 20
	}

	if p < 1 {
		p = 1
	}

	data := QueryData{}

	maxSecond := time.Duration(dynamic.IntValue(dynamic.GetWithKeys(app.GetConfig(), []string{"cache", "maxSecond"}), 60))

	cacheKey := cache.SignKey("q", task.Id, task.Ids, task.Pid, task.Uid, task.Bloguid, task.Q, n, task.P, task.Path)

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

		countTask := CountTask{}

		countTask.Eid = task.Eid
		countTask.Pid = task.Pid
		countTask.Id = task.Id
		countTask.Ids = task.Ids
		countTask.Uid = task.Uid
		countTask.Q = task.Q
		countTask.Path = task.Path

		countData, err := S.Count(app, &countTask)

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

	sql.WriteString(" WHERE state=? AND eid=?")

	args = append(args, CommentState_None, task.Eid)

	if task.Id != nil {
		sql.WriteString(" AND id=?")
		args = append(args, task.Id)
	}

	if task.Pid != nil {
		sql.WriteString(" AND pid=?")
		args = append(args, task.Pid)
	}

	if task.Uid != nil {
		sql.WriteString(" AND uid=?")
		args = append(args, task.Uid)
	}
	if task.Bloguid != nil {
		sql.WriteString(" AND uid!=?")
		args = append(args, task.Bloguid)
	}

	if task.Q != nil {
		sql.WriteString(" AND body LIKE ?")
		args = append(args, fmt.Sprintf("%%%s%%", task.Q))
	}

	if task.Path != nil {
		sql.WriteString(" AND path LIKE ?")
		args = append(args, fmt.Sprintf("%s%%", task.Path))
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

	sql.WriteString(fmt.Sprintf(" ORDER BY id DESC LIMIT %d,%d", (p-1)*n, n))

	app.Println("commentquerypath", sql)
	app.Println(args)

	v := Comment{}

	data.Items = []*Comment{}

	rs, err := db.Query(conn, &v, prefix, sql.String(), args...)

	if err != nil {
		return nil, err
	}

	defer rs.Close()

	scaner := db.NewScaner(&v)

	for rs.Next() {

		err = scaner.Scan(rs)

		if err != nil {
			return nil, err
		}

		item := Comment{}
		item = v

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
