package lookin

import (
	"bytes"
	SQL "database/sql"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/hailongz/golang/cache"
	"github.com/hailongz/golang/dynamic"
	"github.com/hailongz/golang/micro"
)

func (S *Service) Count(app micro.IContext, task *CountTask) (*CountData, error) {

	iid := dynamic.IntValue(task.Iid, 0)

	data := CountData{}

	maxSecond := time.Duration(dynamic.IntValue(dynamic.GetWithKeys(app.GetConfig(), []string{"cache", "maxSecond"}), 60))

	key := fmt.Sprintf("%d_%d", task.Tid, iid)

	cacheKey := cache.SignKey("n", iid, task.Uid, task.Fuid, task.Flevel, task.GroupBy)

	//从缓存中读取
	{
		cache, err := app.GetCache("default")
		if err == nil {
			text, err := cache.GetItem(key, cacheKey, time.Second*maxSecond)
			if err == nil && text != "" {
				err = json.Unmarshal([]byte(text), &data)
				if err == nil {
					return &data, nil
				}
			}
		}
	}

	//缓存没有查询数据库
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

	v := Lookin{}

	var rs *SQL.Rows

	if task.GroupBy == GroupBy_uid {
		rs, err = conn.Query(fmt.Sprintf("SELECT COUNT(DISTINCT uid) FROM %s%s %s", prefix, v.GetName(), sql.String()), args...)
	} else if task.GroupBy == GroupBy_fuid {
		rs, err = conn.Query(fmt.Sprintf("SELECT COUNT(DISTINCT fuid) FROM %s%s %s", prefix, v.GetName(), sql.String()), args...)
	} else {
		rs, err = conn.Query(fmt.Sprintf("SELECT COUNT(*) FROM %s%s %s", prefix, v.GetName(), sql.String()), args...)
	}

	if err != nil {
		return nil, err
	}

	if rs.Next() {
		_ = rs.Scan(&data.Total)
	}

	rs.Close()

	{
		cache, err := app.GetCache("default")

		if err == nil {

			b, _ := json.Marshal(&data)

			cache.SetItem(key, cacheKey, string(b), time.Second*maxSecond)

		}
	}

	return &data, nil
}
