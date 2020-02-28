package top

import (
	"bytes"
	"fmt"
	"strings"
	"time"

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

	if p == 1 && task.Tids == nil && task.TopId == nil && task.Q == nil {

		{

			// 缓存
			redis, prefix, err := app.GetRedis("default")

			if err == nil {

				text, err := redis.Get(fmt.Sprintf("%s%s", prefix, task.Name)).Result()

				if err == nil && text != "" {

					err = json.Unmarshal([]byte(text), &data)

					if err == nil {

						if data.Items != nil && len(data.Items) >= int(n) {

							data.Items = data.Items[0:int(n)]

							if task.P == nil {

								data.Page = nil

							} else if data.Page == nil {

								countTask := CountTask{}

								countTask.Name = task.Name
								countTask.TopId = task.TopId
								countTask.Tids = task.Tids

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

							} else if data.Page.N != n {
								data.Page.N = n
								data.Page.Count = data.Page.Total / n
								if data.Page.Total%n != 0 {
									data.Page.Count++
								}
							}

							return &data, nil
						}

					}

				}
			}

		}
	}

	if task.P != nil {

		countTask := CountTask{}

		countTask.Name = task.Name
		countTask.TopId = task.TopId
		countTask.Tids = task.Tids
		countTask.Q = task.Q

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

	prefix = fmt.Sprintf("%s%s_", prefix, task.Name)

	sql := bytes.NewBuffer(nil)

	args := []interface{}{}

	sql.WriteString(" WHERE 1")

	if task.TopId != nil {
		sql.WriteString(" AND sid <= ?")
		args = append(args, task.TopId)
	}

	if task.Tids != nil {
		sql.WriteString(" AND tid IN(")
		for i, s := range strings.Split(dynamic.StringValue(task.Tids, ""), ",") {
			if i != 0 {
				sql.WriteString(",")
			}
			sql.WriteString("?")
			args = append(args, s)
		}
		sql.WriteString(")")
	}

	if task.Q != nil {
		sql.WriteString(" AND keyword LIKE ?")
		args = append(args, fmt.Sprintf("%%%s%%", task.Q))
	}

	sql.WriteString(fmt.Sprintf(" ORDER BY sid DESC,id DESC LIMIT %d,%d", (p-1)*n, n))

	v := Top{}

	data.Items = []*Top{}

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

		item := Top{}
		item = v

		data.Items = append(data.Items, &item)
	}

	if data.Page != nil {
		if task.TopId != nil {
			data.Page.TopId = dynamic.IntValue(task.TopId, 0)
		} else if len(data.Items) > 0 && p == 1 {
			data.Page.TopId = data.Items[0].Sid
		}
	}

	if p == 1 && task.Tids == nil && task.TopId == nil && task.Q == nil {

		redis, prefix, err := app.GetRedis("default")

		if err == nil {

			b, _ := json.Marshal(&data)

			maxSecond := dynamic.IntValue(dynamic.GetWithKeys(app.GetConfig(), []string{"cache", "maxSecond"}), 60)

			redis.Set(fmt.Sprintf("%s%s", prefix, task.Name), string(b), time.Second*time.Duration(maxSecond)).Result()

		}

	}

	return &data, nil
}
