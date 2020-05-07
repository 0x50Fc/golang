package member

import (
	"encoding/json"
	"fmt"

	"github.com/hailongz/golang/db"
	"github.com/hailongz/golang/dynamic"
	"github.com/hailongz/golang/micro"
)

func (S *Service) Set(app micro.IContext, task *SetTask) (*Member, error) {
	conn, prefix, err := app.GetDB("wd")

	if err != nil {
		return nil, err
	}

	v := Member{}

	p, err := db.Get(conn, &v, prefix, "WHERE bid=? AND uid=?", task.Bid, task.Uid)

	if p == nil {
		return nil, micro.NewError(ERROR_NOT_FOUND, "未找到成员")
	}

	var options interface{} = nil
	keys := map[string]bool{}
	if task.Title != nil {
		v.Title = dynamic.StringValue(task.Title, "")
		keys["title"] = true
	}
	if task.Keyword != nil {
		v.Keyword = dynamic.StringValue(task.Keyword, "")
		keys["keyword"] = true
	}
	if task.Options != nil {
		var opt interface{} = nil
		json.Unmarshal([]byte(dynamic.StringValue(task.Options, "")), &opt)
		options = db.Merge(v.Options, opt)
		v.Options = opt
		if opt != nil {
			keys["options"] = true
		}
	}

	if len(keys) > 0 {
		_, err = db.UpdateWithKeys(conn, &v, prefix, keys)
		if keys["options"] {
			v.Options = options
		}
	}

	if err != nil {
		return nil, err
	}

	cache, _ := app.GetCache("default")

	if cache != nil {
		cache.Del(fmt.Sprintf("B_%d", v.Bid))
		cache.Del(fmt.Sprintf("U_%d", v.Uid))
		cache.Del("Q")
	}

	app.SendMessage(task.GetName(), &v)

	return &v, nil
}
