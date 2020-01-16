package doc

import (
	"fmt"
	"time"

	"github.com/hailongz/golang/db"
	"github.com/hailongz/golang/dynamic"
	"github.com/hailongz/golang/json"
	"github.com/hailongz/golang/micro"
)

func (S *Service) Create(app micro.IContext, task *CreateTask) (*Doc, error) {

	pid := dynamic.IntValue(task.Pid, 0)

	var err error = nil
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

	conn, prefix, err := app.GetDB("wd")

	if err != nil {
		return nil, err
	}

	prefix = Prefix(app, prefix, task.Uid)

	v := Doc{}
	v.Uid = task.Uid
	v.Pid = pid
	v.Title = task.Title
	v.Type = int32(dynamic.IntValue(task.Type, DocType_File))
	v.Keyword = dynamic.StringValue(task.Keyword, "")
	v.Ext = dynamic.StringValue(task.Ext, "")
	v.Ctime = time.Now().Unix()
	v.Atime = v.Ctime
	v.Mtime = v.Ctime

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

	}

	if p != nil {
		v.Path = fmt.Sprintf("%s%d/", p.Path, pid)
	}

	_, err = db.Insert(conn, &v, prefix)

	if err != nil {
		return nil, err
	}

	// MQ 消息
	app.SendMessage(task.GetName(), &v)

	return &v, nil
}
