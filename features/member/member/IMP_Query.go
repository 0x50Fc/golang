package member

import (
	"bytes"
	"encoding/json"
	"fmt"
	"time"

	"github.com/hailongz/golang/cache"
	"github.com/hailongz/golang/db"
	"github.com/hailongz/golang/dynamic"
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

	var key string

	if task.Bid != nil {
		key = fmt.Sprintf("B_%d", dynamic.IntValue(task.Bid, 0))
	} else if task.Uid != nil {
		key = fmt.Sprintf("U_%d", dynamic.IntValue(task.Uid, 0))
	} else {
		key = "Q"
	}

	cacheKey := cache.SignKey("q", task.Bid, task.Uid, n, task.P, task.Q)

	if p == 1 {

		cache, err := app.GetCache("default")

		if err == nil {

			text, err := cache.GetItem(key, cacheKey, maxSecond*time.Second)

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

		countTask.Bid = task.Bid
		countTask.Uid = task.Uid
		countTask.Q = task.Q

		countData, err := S.Count(app, &countTask)

		if err != nil {
			return nil, err
		}

		data.Page = &QueryDataPage{
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

	sql := bytes.NewBuffer(nil)

	args := []interface{}{}

	v := Member{}

	sql.WriteString(" WHERE 1")

	if task.Bid != nil {
		sql.WriteString(" AND bid=?")
		args = append(args, task.Bid)
	}

	if task.Uid != nil {
		sql.WriteString(" AND uid=?")
		args = append(args, task.Uid)
	}

	if task.Q != nil {
		q := fmt.Sprintf("%%%s%%", dynamic.StringValue(task.Q, ""))
		sql.WriteString(" AND (title LIKE ? OR keyword LIKE ?)")
		args = append(args, q, q)
	}

	sql.WriteString(fmt.Sprintf(" ORDER BY id ASC LIMIT %d,%d", (p-1)*n, n))

	data.Items = []*Member{}

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

		item := Member{}
		item = v

		data.Items = append(data.Items, &item)
	}

	if p == 1 {

		cache, err := app.GetCache("default")

		if err == nil {

			b, _ := json.Marshal(&data)

			cache.SetItem(key, cacheKey, string(b), time.Second*maxSecond)

		}
	}

	return &data, nil
}
