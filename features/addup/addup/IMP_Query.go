package addup

import (
	"bytes"
	"fmt"
	"time"

	"github.com/hailongz/golang/cache"
	"github.com/hailongz/golang/dynamic"
	"github.com/hailongz/golang/json"
	"github.com/hailongz/golang/micro"
)

func (S *Service) Query(app micro.IContext, task *QueryTask) (*QueryData, error) {

	conn, name, err := app.GetDB("wd")

	if err != nil {
		return nil, err
	}

	if task.Region == nil {
		name = fmt.Sprintf("%s%s", name, task.Name)
	} else {
		name = fmt.Sprintf("%s%d_%s", name, dynamic.IntValue(task.Region, 0), task.Name)
	}

	app.Println("[NAME]", name)

	maxSecond := time.Duration(dynamic.IntValue(dynamic.GetWithKeys(app.GetConfig(), []string{"cache", "maxSecond"}), 60))

	cacheKey := cache.SignKey(task.CacheKey)

	if task.CacheKey != nil {

		cache, err := app.GetCache("default")

		if err == nil {

			text, err := cache.GetItem(name, cacheKey, maxSecond*time.Second)

			if err == nil && text != "" {
				data := &QueryData{}
				err = json.Unmarshal([]byte(text), data)
				if err == nil {
					return data, nil
				}
			}
		}
	}

	sql := bytes.NewBuffer(nil)
	args := []interface{}{}

	sql.WriteString("SELECT ")

	sql.WriteString(dynamic.StringValue(task.Fields, "*"))

	sql.WriteString(" FROM ")

	sql.WriteString(name)

	if task.Where != nil {
		sql.WriteString(" WHERE ")
		sql.WriteString(dynamic.StringValue(task.Where, ""))
	}

	if task.OrderBy != nil {
		sql.WriteString(" ORDER BY ")
		sql.WriteString(dynamic.StringValue(task.OrderBy, ""))
	}

	if task.GroupBy != nil {
		sql.WriteString(" GROUP BY ")
		sql.WriteString(dynamic.StringValue(task.GroupBy, ""))
	}

	if task.Having != nil {
		sql.WriteString(" HAVING ")
		sql.WriteString(dynamic.StringValue(task.Having, ""))
	}

	if task.Limit != nil {
		sql.WriteString(" LIMIT ")
		sql.WriteString(dynamic.StringValue(task.Limit, ""))
	}

	if task.Args != nil {

		var data interface{} = nil

		err = json.Unmarshal([]byte(dynamic.StringValue(task.Args, "")), &data)

		if err != nil {
			return nil, err
		}

		dynamic.Each(data, func(_, value interface{}) bool {
			args = append(args, value)
			return true
		})

	}

	rs, err := conn.Query(sql.String(), args...)

	if err != nil {
		return nil, err
	}

	items := []interface{}{}

	var columns []string = nil
	var values []interface{} = nil
	var ptrs []interface{} = nil

	n := 0

	defer rs.Close()

	for rs.Next() {

		if columns == nil {
			columns, err = rs.Columns()
			if err != nil {
				return nil, err
			}
			n = len(columns)
			values = make([]interface{}, n)
			ptrs = make([]interface{}, n)
			for i := 0; i < n; i++ {
				ptrs[i] = &values[i]
			}
		}

		for i := 0; i < n; i++ {
			values[i] = nil
		}

		err = rs.Scan(ptrs...)

		if err != nil {
			return nil, err
		}

		item := map[string]interface{}{}

		for i, v := range values {
			if v == nil {
				continue
			}
			{
				b, ok := v.([]byte)
				if ok {
					v = string(b)
				}
			}
			item[columns[i]] = v
		}

		items = append(items, item)
	}

	data := QueryData{Items: items}

	if task.CacheKey != nil {

		cache, err := app.GetCache("default")

		if err == nil {

			b, _ := json.Marshal(&data)

			cache.SetItem(name, cacheKey, string(b), time.Second*maxSecond)

		}
	}

	return &data, nil
}
