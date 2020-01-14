package inbox

import (
	"bytes"
	"fmt"
	"github.com/hailongz/golang/cache"
	"github.com/hailongz/golang/dynamic"
	"github.com/hailongz/golang/json"
	"github.com/hailongz/golang/micro"
	"time"
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

	cacheKey := fmt.Sprintf("%s.%s.%s", task.Type, task.Mid, task.Iid)
	itemKey := cache.SignKey("q", task.P)

	if p == 1 {

		cache, err := app.GetCache("default")

		if err == nil {

			text, err := cache.GetItem(cacheKey, itemKey, maxSecond*time.Second)

			if err == nil && text != "" {
				data := &UserQueryData{}
				err = json.Unmarshal([]byte(text), data)
				if err == nil {
					return data, nil
				}
			}
		}
	}

	if task.P != nil {

		countTask := UserCountTask{}

		countTask.Type = task.Type
		countTask.Mid = task.Mid
		countTask.Iid = task.Iid

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

	v := Inbox{}

	tableCount := dynamic.IntValue(dynamic.GetWithKeys(app.GetConfig(), []string{"table", "count"}), 128)

	sql := bytes.NewBuffer(nil)

	args := []interface{}{}

	sql.WriteString("SELECT uid,mid,iid,fuid,MAX(ctime) as ctime FROM (")

	for i := int64(1); i <= tableCount; i++ {
		if i != 1 {
			sql.WriteString(" UNION ")
		}
		sql.WriteString(fmt.Sprintf("(SELECT uid,mid,iid,fuid,ctime FROM %s%d_%s WHERE type = ? ", prefix, i, v.GetName()))
		args = append(args, task.Type)

		if task.Mid != nil {
			sql.WriteString(" AND mid = ?")
			args = append(args, task.Mid)
		}

		if task.Iid != nil {
			sql.WriteString(" AND iid = ?")
			args = append(args, task.Iid)
		}

		sql.WriteString(" ORDER BY id DESC )")
	}

	sql.WriteString(") as t GROUP BY t.uid")

	sql.WriteString(fmt.Sprintf(" LIMIT %d,%d", (p-1)*n, n))

	app.Println("[SQL]", sql.String())

	data.Items = []*User{}

	rs, err := conn.Query(sql.String(), args...)

	if err != nil {
		return nil, err
	}

	defer rs.Close()

	for rs.Next() {

		item := User{}

		err = rs.Scan(&item.Uid, &item.Mid, &item.Iid, &item.Fuid, &item.Ctime)

		if err != nil {
			return nil, err
		}

		data.Items = append(data.Items, &item)
	}

	if p == 1 {

		cache, err := app.GetCache("default")

		if err == nil {

			b, _ := json.Marshal(&data)

			cache.SetItem(cacheKey, itemKey, string(b), time.Second*maxSecond)

		}
	}

	return &data, nil
}
