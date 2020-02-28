package media

import (
	"fmt"
	"time"

	"github.com/hailongz/golang/db"
	"github.com/hailongz/golang/dynamic"
	"github.com/hailongz/golang/json"
	"github.com/hailongz/golang/micro"
)

func (S *Service) Add(app micro.IContext, task *AddTask) (*Media, error) {

	conn, name, err := app.GetDB("wd")

	if err != nil {
		return nil, err
	}

	if task.Region == nil {
		name = fmt.Sprintf("%s%s", name, dynamic.StringValue(task.Name, ""))
	} else {
		name = fmt.Sprintf("%s%d_%s", name, dynamic.IntValue(task.Region, 0), dynamic.StringValue(task.Name, ""))
	}

	app.Println("[NAME]", name)

	v := Media{}

	v.Type = dynamic.StringValue(task.Type, "")
	v.Title = dynamic.StringValue(task.Title, "")
	v.Keyword = dynamic.StringValue(task.Keyword, "")
	v.Path = dynamic.StringValue(task.Path, "")
	v.Uid = dynamic.IntValue(task.Uid, 0)
	v.Ctime = time.Now().Unix()

	if task.Options != nil {

		options := map[string]interface{}{}

		text := dynamic.StringValue(task.Options, "")

		var data interface{} = nil

		json.Unmarshal([]byte(text), &data)

		dynamic.Each(data, func(key interface{}, value interface{}) bool {
			options[dynamic.StringValue(key, "")] = value
			return true
		})

		v.Options = options
	}

	_, err = db.Insert(conn, &v, name)

	if err != nil {
		return nil, err
	}

	app.SendMessage(task.GetName(), &v)

	return &v, nil
}
