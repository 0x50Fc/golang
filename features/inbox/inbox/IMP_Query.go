package inbox

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

func (S *Service) Query(app micro.IContext, task *QueryTask) (*InboxQueryData, error) {

	maxSecond := dynamic.IntValue(dynamic.GetWithKeys(app.GetConfig(), []string{"cache", "maxSecond"}), 1800)

	p := int32(dynamic.IntValue(task.P, 1))
	n := int32(dynamic.IntValue(task.N, 20))

	if n < 1 {
		n = 20
	}

	if p < 1 {
		p = 1
	}

	cacheKey := cache.SignKey(task.Uid, task.Fuid, task.Type, task.GroupBy, n)

	data := InboxQueryData{}

	if p == 1 && task.P != nil {
		cache, err := app.GetCache("default")
		if err == nil {
			text, err := cache.GetItem(fmt.Sprintf("%d", task.Uid), cacheKey, time.Second*time.Duration(maxSecond))
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

		countTask.Uid = task.Uid
		countTask.Type = task.Type
		countTask.Mid = task.Mid
		countTask.Iid = task.Iid
		countTask.TopId = task.TopId
		countTask.GroupBy = task.GroupBy

		countData, err := S.Count(app, &countTask)

		if err != nil {
			return nil, err
		}

		data.Page = &TopPage{
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

	prefix = Prefix(app, prefix, task.Uid)

	sql := bytes.NewBuffer(nil)

	args := []interface{}{}

	sql.WriteString(" WHERE uid=?")

	args = append(args, task.Uid)

	if task.Fuid != nil {
		sql.WriteString(" AND fuid=?")
		args = append(args, task.Fuid)
	}

	if task.Type != nil {
		sql.WriteString(" AND (type & ?) != 0")
		args = append(args, task.Type)
	}

	if task.Mid != nil {
		sql.WriteString(" AND mid = ?")
		args = append(args, task.Mid)
	}

	if task.Iid != nil {
		sql.WriteString(" AND iid = ?")
		args = append(args, task.Iid)
	}

	if task.TopId != nil {
		sql.WriteString(" AND ctime < ?")
		args = append(args, task.TopId)
	}

	switch dynamic.StringValue(task.GroupBy, "") {
	case GroupBy_mid:
		sql.WriteString(" GROUP BY mid")
		break
	case GroupBy_fuid:
		sql.WriteString(" GROUP BY fuid")
		break
	}

	sql.WriteString(fmt.Sprintf(" ORDER BY ctime DESC, id DESC LIMIT %d,%d", (p-1)*n, n))

	v := Inbox{}

	data.Items = []*Inbox{}

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

		item := Inbox{}
		item = v

		data.Items = append(data.Items, &item)
	}

	if data.Page != nil {
		if task.TopId != nil {
			data.Page.TopId = dynamic.IntValue(task.TopId, 0)
		} else if len(data.Items) > 0 && p == 1 {
			data.Page.TopId = data.Items[0].Ctime
		}
	}

	if p == 1 && task.P != nil {
		cache, err := app.GetCache("default")
		if err == nil {
			b, _ := json.Marshal(&data)
			cache.SetItem(fmt.Sprintf("%d", task.Uid), cacheKey, string(b), time.Second*time.Duration(maxSecond))
		}
	}

	return &data, nil
}
