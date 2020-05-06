package lookin

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/hailongz/golang/cache"
	"github.com/hailongz/golang/db"
	"github.com/hailongz/golang/dynamic"
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

	key := fmt.Sprintf("%d_%d", task.Tid, iid)

	cacheKey := cache.SignKey("q", iid, task.Uid, n, task.P, task.Fuid, task.Flevel, task.GroupBy)

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

		countTask.Tid = task.Tid
		countTask.Iid = task.Iid
		countTask.Uid = task.Uid
		countTask.Fuid = task.Fuid
		countTask.Flevel = task.Flevel
		countTask.GroupBy = task.GroupBy

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

	v := Lookin{}

	sql.WriteString(" WHERE tid=? AND iid=?")

	args = append(args, task.Tid, iid)

	if task.Uid != nil {
		sql.WriteString(" AND uid=?")
		args = append(args, task.Uid)
	}

	if task.Fuid != nil {
		sql.WriteString(" AND fuid=?")
		args = append(args, task.Fuid)
	}

	if task.Flevel != nil {
		vs := strings.Split(dynamic.StringValue(task.Flevel, ""), ",")
		sql.WriteString(" AND flevel IN(")
		for i, v := range vs {
			if i != 0 {
				sql.WriteString(",")
			}
			sql.WriteString("?")
			args = append(args, v)
		}
		sql.WriteString(")")
	}

	if task.GroupBy == GroupBy_uid {
		sql.WriteString(" GROUP BY uid")
	} else if task.GroupBy == GroupBy_fuid {
		sql.WriteString(" GROUP BY fuid")
	}

	sql.WriteString(fmt.Sprintf(" ORDER BY flevel DESC, id DESC LIMIT %d,%d", (p-1)*n, n))

	data.Items = []*Lookin{}

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

		item := Lookin{}
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
