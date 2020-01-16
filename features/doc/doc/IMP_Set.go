package doc

import (
	"fmt"
	"time"

	"github.com/hailongz/golang/db"
	"github.com/hailongz/golang/dynamic"
	"github.com/hailongz/golang/json"
	"github.com/hailongz/golang/micro"
)

func (S *Service) Set(app micro.IContext, task *SetTask) (*Doc, error) {

	v := Doc{}

	conn, prefix, err := app.GetDB("wd")

	if err != nil {
		return nil, err
	}

	prefix = Prefix(app, prefix, task.Uid)

	p, err := db.Get(conn, &v, prefix, "WHERE uid=? AND id=?", task.Uid, task.Id)

	if err != nil {
		return nil, err
	}

	if p == nil {
		return nil, micro.NewError(ERROR_NOT_FOUND, "未找到文档")
	}

	keys := map[string]bool{}

	pid := dynamic.IntValue(task.Pid, 0)

	if task.Pid != nil && v.Pid != pid {

		var p *Doc = nil

		if pid != 0 {
			getTask := GetTask{}
			getTask.Id = pid
			getTask.Uid = task.Uid
			p, err = S.Get(app, &getTask)
			if err != nil {
				return nil, err
			}
		}

		v.Pid = pid

		if p != nil {
			v.Path = fmt.Sprintf("%s%d/", p.Path, pid)
		} else {
			v.Path = ""
		}

		keys["pid"] = true
		keys["path"] = true
	}

	if task.Title != nil {

		v.Title = dynamic.StringValue(task.Title, "")

		keys["title"] = true
	}

	if task.Keyword != nil {

		v.Keyword = dynamic.StringValue(task.Keyword, "")

		keys["keyword"] = true
	}

	if task.Ext != nil {

		v.Ext = dynamic.StringValue(task.Ext, "")

		keys["ext"] = true
	}

	if task.Options != nil {

		options := map[string]interface{}{}

		dynamic.Each(v.Options, func(key interface{}, value interface{}) bool {
			options[dynamic.StringValue(key, "")] = value
			return true
		})

		text := dynamic.StringValue(task.Options, "")

		var data interface{} = nil

		json.Unmarshal([]byte(text), &data)

		dynamic.Each(data, func(key interface{}, value interface{}) bool {
			options[dynamic.StringValue(key, "")] = value
			return true
		})

		v.Options = options
		keys["options"] = true
	}

	if dynamic.BooleanValue(task.Mtime, false) {
		v.Mtime = time.Now().Unix()
		keys["mtime"] = true
	}

	if dynamic.BooleanValue(task.Atime, false) {
		v.Atime = time.Now().Unix()
		keys["atime"] = true
	}

	if len(keys) > 0 {
		_, err = db.UpdateWithKeys(conn, &v, prefix, keys)
		if err != nil {
			return nil, err
		}
	}

	cache, _ := app.GetCache("default")

	if cache != nil {
		cache.Del(fmt.Sprintf("%d/%d", task.Uid, task.Id))
	}

	// MQ 消息
	app.SendMessage(task.GetName(), &v)

	return &v, nil
}
