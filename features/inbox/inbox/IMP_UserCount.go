package inbox

import (
	"bytes"
	"fmt"
	"github.com/hailongz/golang/dynamic"
	"github.com/hailongz/golang/json"
	"github.com/hailongz/golang/micro"
	"time"
)

func (S *Service) UserCount(app micro.IContext, task *UserCountTask) (*UserCountData, error) {

	maxSecond := time.Duration(dynamic.IntValue(dynamic.GetWithKeys(app.GetConfig(), []string{"cache", "maxSecond"}), 60))

	cacheKey := fmt.Sprintf("%s.%s.%s", task.Type, task.Mid, task.Iid)

	{
		cache, err := app.GetCache("default")

		if err == nil {

			text, err := cache.GetItem(cacheKey, "n", maxSecond*time.Second)

			if err == nil && text != "" {
				c := &UserCountData{}
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

	v := Inbox{}

	tableCount := dynamic.IntValue(dynamic.GetWithKeys(app.GetConfig(), []string{"table", "count"}), 128)

	sql := bytes.NewBuffer(nil)

	args := []interface{}{}

	sql.WriteString("SELECT COUNT(DISTINCT uid) as n FROM (")

	for i := int64(1); i <= tableCount; i++ {
		if i != 1 {
			sql.WriteString(" UNION ")
		}
		sql.WriteString(fmt.Sprintf("(SELECT uid,mid,iid,MAX(ctime) as ctime FROM %s%d_%s WHERE Type = ?  ", prefix, i, v.GetName()))
		args = append(args, task.Type)
		if task.Mid != nil {
			sql.WriteString(" AND mid = ?")
			args = append(args, task.Mid)
		}

		if task.Iid != nil {
			sql.WriteString(" AND iid = ?")
			args = append(args, task.Iid)
		}

		sql.WriteString(")")

	}

	sql.WriteString(") as t")

	app.Println("[SQL]", sql.String())

	data := UserCountData{}

	err = func() error {

		rs, err := conn.Query(sql.String(), args...)

		if err != nil {
			return err
		}

		defer rs.Close()

		if rs.Next() {

			err = rs.Scan(&data.Total)

			if err != nil {
				return err
			}
		}

		return nil
	}()

	if err != nil {
		return nil, err
	}

	{
		cache, err := app.GetCache("default")

		if err == nil {

			b, _ := json.Marshal(&data)

			cache.SetItem(cacheKey, "n", string(b), time.Second*maxSecond)

		}
	}

	return &data, nil
}
