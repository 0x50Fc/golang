package article

import (
	"fmt"
	"time"

	"github.com/hailongz/golang/db"
	"github.com/hailongz/golang/dynamic"
	"github.com/hailongz/golang/json"
	"github.com/hailongz/golang/micro"
)

func (S *Service) Get(app micro.IContext, task *GetTask) (*Article, error) {

	maxSecond := time.Duration(dynamic.IntValue(dynamic.GetWithKeys(app.GetConfig(), []string{"cache", "maxSecond"}), 60))

	cache, _ := app.GetCache("default")

	if cache != nil {

		text, err := cache.Get(fmt.Sprintf("%d", task.Id), maxSecond*time.Second)

		if err == nil && text != "" {

			v := Article{}

			err = json.Unmarshal([]byte(text), &v)

			if err == nil {
				return &v, nil
			}

		}
	}

	conn, prefix, err := app.GetDB("rd")

	if err != nil {
		return nil, err
	}

	prefix = Prefix(app, prefix, task.Id)

	v := Article{}

	rs, err := db.Query(conn, &v, prefix, " WHERE id=?", task.Id)

	if err != nil {
		return nil, err
	}

	defer rs.Close()

	if rs.Next() {

		scaner := db.NewScaner(&v)

		err = scaner.Scan(rs)

		if err != nil {
			return nil, err
		}

	} else {
		return nil, micro.NewError(ERROR_NOT_FOUND, "未找到动态")
	}

	if cache != nil {
		b, _ := json.Marshal(&v)
		cache.Set(fmt.Sprintf("%d", v.Id), string(b), maxSecond*time.Second)
	}

	return &v, nil
}
