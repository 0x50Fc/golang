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

func (S *Service) Query(app micro.IContext, task *QueryTask) (*QueryData, error) {

	iid := dynamic.IntValue(task.Iid, 0)

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

	cacheKey := cache.SignKey("q", iid, task.Uid, n, task.P)

	if p == 1 {

		cache, err := app.GetCache("default")

		if err == nil {

			text, err := cache.GetItem(fmt.Sprintf("%d", task.Tid), cacheKey, maxSecond*time.Second)

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

		countTask.Tid = task.Tid
		countTask.Iid = task.Iid
		countTask.Uid = task.Uid

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

	prefix = Prefix(app, prefix, task.Tid)

	sql := bytes.NewBuffer(nil)

	args := []interface{}{}

	sql.WriteString(" WHERE tid=? AND iid=?")

	args = append(args, task.Tid, iid)

	if task.Uid != nil {
		sql.WriteString(" AND uid=?")
		args = append(args, task.Uid)
	}

	sql.WriteString(fmt.Sprintf(" ORDER BY id DESC LIMIT %d,%d", (p-1)*n, n))

	v := Like{}

	data.Items = []*Like{}

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

		item := Like{}
		item = v

		data.Items = append(data.Items, &item)
	}

	if p == 1 {

		cache, err := app.GetCache("default")

		if err == nil {

			b, _ := json.Marshal(&data)

			cache.SetItem(fmt.Sprintf("%d", task.Tid), cacheKey, string(b), time.Second*maxSecond)

		}
	}

	return &data, nil
}
