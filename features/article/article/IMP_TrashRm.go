package article

import (
	"fmt"
	"time"

	"github.com/hailongz/golang/db"
	"github.com/hailongz/golang/dynamic"
	"github.com/hailongz/golang/json"
	"github.com/hailongz/golang/micro"
)

func (S *Service) TrashRm(app micro.IContext, task *TrashRmTask) (*Article, error) {

	conn, fix, err := app.GetDB("wd")

	if err != nil {
		return nil, err
	}

	prefix := Prefix(app, fix, task.Id)

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

	v.State = ArticleState_None

	_, err = db.UpdateWithKeys(conn, &v, prefix, map[string]bool{"state": true})

	if err != nil {
		return nil, err
	}

	{
		maxSecond := time.Duration(dynamic.IntValue(dynamic.GetWithKeys(app.GetConfig(), []string{"cache", "maxSecond"}), 60))

		cache, _ := app.GetCache("default")

		if cache != nil {
			b, _ := json.Marshal(&v)
			cache.Set(fmt.Sprintf("%d", v.Id), string(b), maxSecond*time.Second)
		}
	}

	app.SendMessage(task.GetName(), &v)

	return &v, nil
}
